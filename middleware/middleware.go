package middleware

import (
	"net/http"
)

// middleware interface 实现当前接口的中间件能够灵活处理中间件返回的类型 http.Handler 或则 http.HandlerFunc
type Middleware interface {
	// 封装 http.Handler 的中间件
	MidHandler(handler http.Handler) http.Handler
	// 封装 http.HandlerFunc 的中间件
	MidHandlerFunc(handler http.HandlerFunc) http.HandlerFunc
}

// MiddlewareHandler 中间件类型 该中间件执行后返回的是 http.Handler
type MiddlewareHandler func(handler http.Handler) http.Handler

// MiddlewareHandlerFunc 中间件类型 该中间件执行后返回的是 http.HandlerFunc
type MiddlewareHandlerFunc func(handler http.HandlerFunc) http.HandlerFunc

func (m MiddlewareHandler) MidHandler(handler http.Handler) http.Handler {
	return m(handler)
}

func (m MiddlewareHandler) MidHandlerFunc(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { m(handler).ServeHTTP(w, r) }
}

func (m MiddlewareHandlerFunc) MidHandler(handler http.Handler) http.Handler {
	return m(func(w http.ResponseWriter, r *http.Request) { handler.ServeHTTP(w, r) })
}

func (m MiddlewareHandlerFunc) MidHandlerFunc(handler http.HandlerFunc) http.HandlerFunc {
	return m(handler)
}