package middleware

import (
	"net/http"

	"github.com/tangzixiang/mplus/mhttp"
)

// Use 使用 MiddlewareHandlerFunc 系列中间件
func Use(ms ...MiddlewareHandlerFunc) MiddlewareHandlerFunc {

	return func(handler http.HandlerFunc) http.HandlerFunc {

		// Wrapped like an onion
		for _, m := range ms {
			handler = m.MidHandlerFunc(handler)
		}

		return handler
	}
}

// UseHandlerMiddleware 使用 MiddlewareHandler 系列中间件
func UseHandlerMiddleware(ms ...MiddlewareHandler) MiddlewareHandler {

	return func(handler http.Handler) http.Handler {

		// Wrapped like an onion
		for _, m := range ms {
			handler = m.MidHandler(handler)
		}

		return handler
	}
}

// Thunk 创建一个请求链，请求链中的任意请求处理函数可以通过 Abort() 方法中断后续请求
func Thunk(handlers ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, httpHandler := range handlers {
			if httpHandler == nil {
				continue
			}

			if mhttp.IsAbort(r) {
				return
			}

			httpHandler.ServeHTTP(w, r)
		}
	}
}

// ThunkHandler 创建一个请求链，请求链中的任意请求处理函数可以通过 Abort() 方法中断后续请求
func ThunkHandler(handlers ...http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, httpHandler := range handlers {
			if httpHandler == nil {
				continue
			}

			if mhttp.IsAbort(r) {
				return
			}

			httpHandler.ServeHTTP(w, r)
		}
	}
}
