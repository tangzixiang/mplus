package middleware

import (
	"net/http"

	"github.com/satori/go.uuid"
	"github.com/tangzixiang/mplus/context"
	"github.com/tangzixiang/mplus/header"
	"github.com/tangzixiang/mplus/mhttp"
)

// Pre 初始话上下文的中间件，必须作为第一个中间件使用，使用 mplus 路由功能必须初始化上下文
func Pre(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(mhttp.NewResponseWrite(w), r.WithContext(context.NewContext(r.Context())))
	}
}

// PreHandler 初始话上下文的中间件，必须作为第一个中间件使用，使用 mplus 路由功能必须初始化上下文
func PreHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(mhttp.NewResponseWrite(w), r.WithContext(context.NewContext(r.Context())))
	})
}

// RequestID 为每个请求配置 request-id
func RequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request-ID
		requestID := header.GetHeaderRequestID(r)
		if requestID == "" {
			requestID = uuid.Must(uuid.NewV4()).String()
			header.SetRequestHeaderRequestID(r, requestID)
			header.SetResponseHeaderRequestID(w, requestID)
		}

		next.ServeHTTP(w, r)
	}
}

// RequestIDHandler 为每个请求配置 request-id
func RequestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request-ID
		requestID := header.GetHeaderRequestID(r)
		if requestID == "" {
			requestID = uuid.Must(uuid.NewV4()).String()
			header.SetRequestHeaderRequestID(r, requestID)
			header.SetResponseHeaderRequestID(w, requestID)
		}

		next.ServeHTTP(w, r)
	})
}

var (
	_ MiddlewareHandlerFunc = Pre
	_ MiddlewareHandler     = PreHandler
	_ MiddlewareHandlerFunc = RequestID
	_ MiddlewareHandler     = RequestIDHandler
)
