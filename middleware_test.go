package mplus

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cacheLock sync.Mutex
	cacheNums []int
)

func TestThunk(t *testing.T) {

	handlers := []http.Handler{
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

func handlerFunc(num int, abort bool) http.HandlerFunc {
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

func TestBindFailedButGotBodyData(t *testing.T) {

	type V struct {
		Addr string `json:"addr" validate:"min=10"` // min len is 10
	}

	jsonData := `{"addr":"1"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost", strings.NewReader(jsonData))
	request = request.WithContext(NewContext(request.Context()))
	SetRequestHeader(request, ContentTypeHeader, ContentTypeJSON)

	response := httptest.NewRecorder()

	Bind((*V)(nil)).ServeHTTP(response, request)

	resp := response.Result()

	if !assert.Equal(t, http.StatusBadRequest, resp.StatusCode) { // must 400
		return
	}

	if !assert.Equal(t, jsonData, string(GetContextValueBytes(request.Context(), BodyData))) {
		return
	}
}
