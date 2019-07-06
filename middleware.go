package mplus

import (
	"fmt"
	"net/http"

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
// example: Check(schema.DeviceLoginValidate{})
func Bind(validateData interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		if validateData == nil {
			fmt.Println("[muxplus] validate object can't be nil")
			InternalServerError(w, r)
			return
		}

		if validateData, err = checkValidateData(w, r, validateData); err != nil {
			fmt.Printf("[muxplus] get validate object failed : %v\n", err)
			dealValidateResultErr(w, r, err)
			return
		}

		vr := ParseValidate(r, &validateData)
		if vr.Err != nil {
			fmt.Printf("[muxplus] parse validate object failed: %v\n", err)
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		SetContextValue(r.Context(), BodyData, vr.BodyBytes)
		SetContextValue(r.Context(), ReqData, validateData)
	}
}

func checkValidateData(w http.ResponseWriter, r *http.Request, validateData interface{}) (interface{}, error) {
	switch voType := validateData.(type) {
	case ValidateFunc:
		var err error
		if validateData, err = voType(r); err != nil {
			return nil, err
		}
	default:
		// do nothing
	}
	return validateData, nil
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
