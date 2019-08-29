package mplus

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/pkg/errors"
)

// Thunk 创建一个请求链，请求链中的任意请求处理函数可以通过 Abort() 方法中断后续请求
func Thunk(handlers ...http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, httpHandler := range handlers {
			if httpHandler == nil {
				continue
			}

			if IsAbort(r) {
				return
			}

			httpHandler.ServeHTTP(w, r)
		}
	}
}

// Bind 通用型校验请求数据有效性的 request middleware,该方法会校验请求对象并置于请求上下文中
//
// 第一步 如果 validateData 是 ValidateFunc 类型，则调用该函数返回 VO (请求数据对象)，执行第二步
//
// 第二步 如果 validateData 是 VO 则根据 tag 进行解析校验，或则根据第一步返回的 VO 根据 tag 进行解析校验
//
// 第三步 校验失败则终止请求链，校验成功则将 VO 对象放入 request context
//
// example: Bind((*VO)(nil))
//
func Bind(validateData interface{}) http.HandlerFunc {

	// 入参只允许指针及函数类型
	checkBindType(validateData)

	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var vo interface{}
		var vr ValidateResult

		// 1. 解析数据
		if parse(r, &vr); vr.Err != nil {
			fmt.Printf("[muxplus] [request:%v] parse data failed : %v\n", GetHeaderRequestID(r), vr.Err)
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		if len(vr.BodyBytes) != 0 { // 放这里是因为有可能序列化失败但是 body 读取成功，可以供后续使用
			SetContextValue(r.Context(), BodyData, vr.BodyBytes)
		}

		// 得到一个全新的 VO 的对象
		if vo, err = checkValidateData(w, r, validateData); err != nil {
			fmt.Printf("[muxplus] [request:%v] get validate object failed : %v\n", GetHeaderRequestID(r), err)
			dealValidateResultErr(w, r, err)
			return
		}

		// 绑定数据
		if decodeTo(r, vo, &vr); vr.Err != nil {
			fmt.Printf("[muxplus] [request:%v] decode data to object failed : %v\n", GetHeaderRequestID(r), vr.Err)
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		// 数据校验
		if bindValidate(r, vo, &vr); vr.Err != nil {
			fmt.Printf("[muxplus] [request:%v] decode data to object failed : %v\n", GetHeaderRequestID(r), vr.Err)
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		SetContextValue(r.Context(), ReqData, vo)
	}
}

func checkBindType(validateData interface{}) {
	vType := reflect.TypeOf(validateData)

	if vType.Kind() != reflect.Ptr && vType.Kind() != reflect.Func {
		panic(errors.New("[muxplus] bind data must be ptr or func"))
	}
}

func checkValidateData(w http.ResponseWriter, r *http.Request, validateData interface{}) (interface{}, error) {

	// 如果是函数则获取函数返回的值,返回值只能是指针
	if voTypeFunc, ok := validateData.(ValidateFunc); ok {
		var err error
		if validateData, err = voTypeFunc(r); err != nil {
			return nil, ValidateError{lastErr: err, errType: ErrDecode}
		}

		if reflect.TypeOf(validateData).Kind() != reflect.Ptr {
			return nil, errors.New("[muxplus] bind data must be ptr or func")
		}
	}

	return reflect.New( // new vo
		reflect.TypeOf(validateData). // get type ptr
			Elem(), // get type
	).Interface(), nil
}

func dealValidateResultErr(w http.ResponseWriter, r *http.Request, err error) {
	cErr, ok := errors.Cause(err).(ValidateError)

	if !ok {
		if errHandler, exists := validateErrorHub[ErrDefault]; exists {
			errHandler(w, r, err)
			return
		}

		if errHandler, exists := defaultValidateErrorHub[ErrDefault]; exists {
			errHandler(w, r)
			return
		}
	}

	errType := cErr.Type()

	if errHandler, exists := validateErrorHub[errType]; exists {
		errHandler(w, r, cErr)
		return
	}

	if errHandler, exists := defaultValidateErrorHub[errType]; exists {
		errHandler(w, r)
		return
	}

	BadRequest(w, r)
	return
}
