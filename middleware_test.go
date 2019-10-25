package mplus

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestUse(t *testing.T) {
	var ms []MiddlewareHandlerFunc
	var orders []int

	for i := 1; i <= 10; i++ {
		(func(i int) {
			ms = append(ms, func(handler http.HandlerFunc) http.HandlerFunc {
				return func(writer http.ResponseWriter, request *http.Request) {
					orders = append(orders, i)
					handler.ServeHTTP(writer, request)
					orders = append(orders, i)
				}
			})
		})(i)
	}

	Use(ms...).MidHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 	do nothing
	})).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "http:127.0.0.1", nil))

	assert.ElementsMatch(t, orders, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestUseHandlerMiddleware(t *testing.T) {
	var ms []MiddlewareHandler
	var orders []int

	for i := 1; i <= 10; i++ {
		(func(i int) {
			ms = append(ms, func(handler http.Handler) http.Handler {
				return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
					orders = append(orders, i)
					handler.ServeHTTP(writer, request)
					orders = append(orders, i)
				})
			})
		})(i)
	}

	UseHandlerMiddleware(ms...).MidHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 	do nothing
	})).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "http:127.0.0.1", nil))

	assert.ElementsMatch(t, orders, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestThunkHandler(t *testing.T) {

	var (
		cacheLock sync.Mutex
		cacheNums []int
	)

	var handlerFunc = func(num int, abort bool) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			if abort {
				Abort(r)
				return
			}

			cacheLock.Lock()
			cacheNums = append(cacheNums, num)
			cacheLock.Unlock()
		}
	}

	handlers := []http.Handler{
		handlerFunc(1, false),
		handlerFunc(2, false),
		handlerFunc(3, true),
		handlerFunc(4, false),
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	response := httptest.NewRecorder()

	ThunkHandler(handlers...).ServeHTTP(response, request.WithContext(NewContext(request.Context())))

	assert.Contains(t, cacheNums, 1)
	assert.Contains(t, cacheNums, 2)
	assert.NotContains(t, cacheNums, 3)
	assert.NotContains(t, cacheNums, 4)
}

func TestThunk(t *testing.T) {

	var (
		cacheLock sync.Mutex
		cacheNums []int
	)

	var handlerFunc = func(num int, abort bool) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			if abort {
				Abort(r)
				return
			}

			cacheLock.Lock()
			cacheNums = append(cacheNums, num)
			cacheLock.Unlock()
		}
	}

	handlers := []http.HandlerFunc{
		handlerFunc(1, false),
		handlerFunc(2, false),
		handlerFunc(3, true),
		handlerFunc(4, false),
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	response := httptest.NewRecorder()

	Thunk(handlers...).ServeHTTP(response, request.WithContext(NewContext(request.Context())))

	assert.Contains(t, cacheNums, 1)
	assert.Contains(t, cacheNums, 2)
	assert.NotContains(t, cacheNums, 3)
	assert.NotContains(t, cacheNums, 4)
}

func TestMiddlewareOrder(t *testing.T) {

	var orders []int

	before1 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 7)
	}

	before2 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 8)
	}

	after1 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 10)
	}

	after2 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 11)
	}

	mid1 := func(next http.HandlerFunc) http.HandlerFunc {

		orders = append(orders, 3)
		return func(w http.ResponseWriter, r *http.Request) {
			orders = append(orders, 4)
			// call next
			next.ServeHTTP(w, r)
		}
	}

	mid2 := func(next http.HandlerFunc) http.HandlerFunc {

		orders = append(orders, 2)
		return func(w http.ResponseWriter, r *http.Request) {
			orders = append(orders, 5)
			// call next
			next.ServeHTTP(w, r)
		}
	}

	mid3 := func(next http.HandlerFunc) http.HandlerFunc {

		orders = append(orders, 1)
		return func(w http.ResponseWriter, r *http.Request) {
			orders = append(orders, 6)
			// call next
			next.ServeHTTP(w, r)
		}
	}

	Hello := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 9)
	}

	route := MRote().Use(mid1, mid2, mid3).Before(before1, before2).After(after1, after2).HandlerFunc(Hello)
	route.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "http://localhost", nil))
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, orders)

	orders = nil
	route.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "http://localhost", nil))
	assert.ElementsMatch(t, []int{4, 5, 6, 7, 8, 9, 10, 11}, orders)
}

func TestMiddlewareOrderAbort(t *testing.T) {

	var orders []int

	before1 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 7)
	}

	before2 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 8)
		Abort(r)
	}

	after1 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 10)
	}

	after2 := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 11)
	}

	mid1 := func(next http.HandlerFunc) http.HandlerFunc {

		orders = append(orders, 3)
		return func(w http.ResponseWriter, r *http.Request) {
			orders = append(orders, 4)
			// call next
			next.ServeHTTP(w, r)
		}
	}

	mid2 := func(next http.HandlerFunc) http.HandlerFunc {

		orders = append(orders, 2)
		return func(w http.ResponseWriter, r *http.Request) {
			orders = append(orders, 5)
			// call next
			next.ServeHTTP(w, r)
		}
	}

	mid3 := func(next http.HandlerFunc) http.HandlerFunc {

		orders = append(orders, 1)
		return func(w http.ResponseWriter, r *http.Request) {
			orders = append(orders, 6)
			// call next
			next.ServeHTTP(w, r)
		}
	}

	Hello := func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 9)
	}

	route := MRote().Use(mid1, mid2, mid3).Before(before1, before2).After(after1, after2).HandlerFunc(Hello)
	route.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "http://localhost", nil))
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, orders)

	orders = nil
	route.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "http://localhost", nil))
	assert.ElementsMatch(t, []int{4, 5, 6, 7, 8}, orders)
}

func TestBindFailedButGotBodyData(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	jsonData := `{"addr":"1"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
	request = request.WithContext(NewContext(request.Context()))
	SetRequestHeader(request, HeaderContentType, ContentTypeJSON)

	response := httptest.NewRecorder()

	Bind((*V)(nil)).ServeHTTP(response, request)

	resp := response.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400
	assert.Equal(t, jsonData, string(GetContextValueBytes(request.Context(), BodyData)))
}

func TestBindFailedWithRegisterErrHandler(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	jsonData := `{"addr":"1"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
	request = request.WithContext(NewContext(request.Context()))
	SetRequestHeader(request, HeaderContentType, ContentTypeJSON)

	response := httptest.NewRecorder()

	// 使用通道模拟锁的机制防止与其他测试冲突
	BeforeTest(true)
	defer AfterTest(true)

	RegisterValidateErrorFunc(ErrBodyValidate, func(w http.ResponseWriter, r *http.Request, err error) {
		PlusPlus(w, r).JSON(map[string]interface{}{"err_message": err.Error()}, 400)
	})

	Bind((*V)(nil)).ServeHTTP(response, request)

	resp := response.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, string(body), `{"err_message":"Key: 'V.Addr' Error:Field validation for 'Addr' failed on the 'min' tag"}`)
	assert.Equal(t, jsonData, string(GetContextValueBytes(request.Context(), BodyData)))
}

func TestBindCheckBindType(t *testing.T) {
	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	assert.Panics(t, func() {
		Bind(V{})
	})

	assert.Panics(t, func() {
		Bind(map[string]interface{}{})
	})

	assert.NotPanics(t, func() {
		Bind(ValidateFunc(func(r *http.Request) (interface{}, error) {
			return nil, nil
		}))
	})

	assert.NotPanics(t, func() {
		Bind(&V{})
	})
}

func TestBindErrMediaType(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	allMethods := []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}

	allMediaType := []string{MIMEJSON, MIMEHTML, MIMEXML, MIMEXML2, MIMEPlain, MIMEPOSTForm, MIMEMultipartPOSTForm, MIMEPROTOBUF, MIMEMSGPACK, MIMEMSGPACK2, MIMEStream}
	allowMediaType := map[string]bool{MIMEJSON: true, MIMEPOSTForm: true, MIMEMultipartPOSTForm: true}

	for _, method := range allMethods {

		for _, mediaType := range allMediaType {
			t.Run("test "+method+" on "+mediaType, func(t *testing.T) {
				request := httptest.NewRequest(method, "http://localhost", nil)
				request = request.WithContext(NewContext(request.Context()))
				response := httptest.NewRecorder()

				SetRequestHeader(request, HeaderContentType, mediaType)

				Bind((*V)(nil)).ServeHTTP(response, request)

				resp := response.Result()

				if !allowMediaType[mediaType] {
					assert.Equal(t, http.StatusUnsupportedMediaType, resp.StatusCode)
				}
			})
		}
	}
}

func TestBindErrMediaTypeParse(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	t.Run("check code", func(t *testing.T) {
		jsonData := `{"addr":"1"}`
		request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, "application/")

		response := httptest.NewRecorder()

		// 使用通道模拟锁的机制防止与其他测试冲突
		BeforeTest(true)
		defer AfterTest(true)

		RegisterValidateErrorFunc(ErrMediaTypeParse, func(w http.ResponseWriter, r *http.Request, err error) {
			PlusPlus(w, r).JSON(map[string]interface{}{"err_message": err.Error()}, http.StatusUnsupportedMediaType)
		})

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		assert.Equal(t, http.StatusUnsupportedMediaType, resp.StatusCode) // must 415

	})

	t.Run("check with hook", func(t *testing.T) {
		jsonData := `{"addr":"1"}`
		request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, "application/")

		response := httptest.NewRecorder()

		// 使用通道模拟锁的机制防止与其他测试冲突
		BeforeTest(true)
		defer AfterTest(true)

		RegisterValidateErrorFunc(ErrMediaTypeParse, func(w http.ResponseWriter, r *http.Request, err error) {
			PlusPlus(w, r).JSON(map[string]interface{}{"err_message": err.Error()}, http.StatusUnsupportedMediaType)
		})

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusUnsupportedMediaType, resp.StatusCode) // must 415
		assert.Equal(t, string(body), "{\"err_message\":\"mime: expected token after slash\"}")
	})
}

func TestBindErrBodyRead(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	t.Run("check code", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodPost, "http://localhost", nil)
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, ContentTypeJSON)

		response := httptest.NewRecorder()

		// 使用通道模拟锁的机制防止与其他测试冲突
		BeforeTest(true)
		defer AfterTest(true)

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400
	})

	t.Run("check with hook", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodPost, "http://localhost", nil)
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, ContentTypeJSON)

		response := httptest.NewRecorder()

		// 使用通道模拟锁的机制防止与其他测试冲突
		BeforeTest(true)
		defer AfterTest(true)

		RegisterValidateErrorFunc(ErrBodyRead, func(w http.ResponseWriter, r *http.Request, err error) {
			PlusPlus(w, r).JSON(map[string]interface{}{"err_message": err.Error()}, 400)
		})

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		assert.Equal(t, string(body), `{"err_message":"body empty"}`)
	})

}

func TestBindErrBodyUnmarshal(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	t.Run("check code", func(t *testing.T) {
		jsonData := `{"addr":"1"` // error json bytes
		request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, ContentTypeJSON)

		response := httptest.NewRecorder()

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400
	})

	t.Run("check with hook", func(t *testing.T) {
		jsonData := `{"addr":"1"` // error json bytes
		request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, ContentTypeJSON)

		response := httptest.NewRecorder()

		// 使用通道模拟锁的机制防止与其他测试冲突
		BeforeTest(true)
		defer AfterTest(true)

		RegisterValidateErrorFunc(ErrBodyUnmarshal, func(w http.ResponseWriter, r *http.Request, err error) {
			PlusPlus(w, r).JSON(map[string]interface{}{"err_message": err.Error()}, 400)
		})

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		assert.Equal(t, string(body), `{"err_message":"unexpected end of JSON input"}`)
	})

}

func TestBindErrBodyParse(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	t.Run("check code", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "http://localhost", nil)
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, ContentTypeForm)

		response := httptest.NewRecorder()

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400
	})

	t.Run("check with hook", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(""))
		request = request.WithContext(NewContext(request.Context()))
		SetRequestHeader(request, HeaderContentType, ContentTypeMultipartPOSTForm)

		response := httptest.NewRecorder()

		// 使用通道模拟锁的机制防止与其他测试冲突
		BeforeTest(true)
		defer AfterTest(true)

		RegisterValidateErrorFunc(ErrBodyParse, func(w http.ResponseWriter, r *http.Request, err error) {
			PlusPlus(w, r).JSON(map[string]interface{}{"err_message": err.Error()}, 400)
		})

		Bind((*V)(nil)).ServeHTTP(response, request)

		resp := response.Result()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		assert.Equal(t, string(body), `{"err_message":"no multipart boundary param in Content-Type"}`)
	})
}

type V struct {
	Addr string `json:"addr"` // min len is 10
}

func (v V) Validate(r *http.Request) (ok bool /*校验是否成功*/, errMsg string /*校验失败的原因*/) {
	return false, "addr not found"
}

func TestBindErrRequestValidate(t *testing.T) {

	jsonData := `{"addr":"1"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
	request = request.WithContext(NewContext(request.Context()))
	SetRequestHeader(request, HeaderContentType, ContentTypeJSON)

	response := httptest.NewRecorder()

	// 使用通道模拟锁的机制防止与其他测试冲突
	BeforeTest(true)
	defer AfterTest(true)

	// RegisterGlobalValidateErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
	// 	cErr :=  errors.Cause(err).(errs.ValidateError)
	// 	PlusPlus(w, r).JSON(map[string]interface{}{"err_message":cErr.String()}, 400)
	// })

	RegisterValidateErrorFunc(ErrRequestValidate, func(w http.ResponseWriter, r *http.Request, err error) {
		PlusPlus(w, r).JSON(map[string]interface{}{"err_message": err.Error()}, 400)
	})

	Bind((*V)(nil)).ServeHTTP(response, request)

	resp := response.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // must 400

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, string(body), `{"err_message":"addr not found"}`)
}
