package route

import (
	"net/http"

	"github.com/tangzixiang/mplus/middleware"
)

// mRote 可复用型中间件路由
type mRote struct {
	middlewares   []middleware.Middleware
	before, after []http.Handler
}

type Route = mRote

// MRote 获取一个可复用型中间件路由实例
func MRote() *mRote {
	return new(mRote).UseHandlerMiddleware(middleware.PreHandler).Use(middleware.Pre) // 初始化上下文
}

// EmptyMRote 获取一个非内置初始化上下文的可复用型中间件路由实例，可以用于配合其他框架使用，需要在外部对上下文进行初始化
func EmptyMRote() *mRote {
	return new(mRote) // 初始化上下文
}

// Handler 获取通过前置或后置请求处理器及中间件进行封装后的 handler
func (mr *mRote) Handler(handler http.Handler) http.Handler {

	// 1. 包裹所有前置请求处理器宝及当前 handler
	// 2. 包裹当前 handler 及所有后置请求处理器
	// 3. 将当前 handler 放入 middleware 层层封装

	if len(mr.before) > 0 {
		handler = middleware.ThunkHandler(append(mr.before, handler)...)
	}

	if len(mr.after) > 0 {
		handler = middleware.ThunkHandler(append([]http.Handler{handler}, mr.after...)...)
	}

	if len(mr.middlewares) > 0 {

		for i := len(mr.middlewares); i > 0; i-- {
			handler = mr.middlewares[i-1].MidHandler(handler)
		}
	}

	return handler
}

// HandlerFunc 获取通过前置或后置请求处理器及中间件进行封装后的 HandlerFunc
func (mr *mRote) HandlerFunc(handler http.HandlerFunc) http.HandlerFunc {

	// 1. 包裹所有前置请求处理器宝及当前 handler
	// 2. 包裹当前 handler 及所有后置请求处理器
	// 3. 将当前 handler 放入 middleware 层层封装

	if len(mr.before) > 0 {
		handler = middleware.ThunkHandler(append(mr.before, handler)...)
	}

	if len(mr.after) > 0 {
		handler = middleware.ThunkHandler(append([]http.Handler{handler}, mr.after...)...)
	}

	if len(mr.middlewares) > 0 {

		for i := len(mr.middlewares); i > 0; i-- {
			handler = mr.middlewares[i-1].MidHandlerFunc(handler)
		}
	}

	return handler
}

// Use 使用 MiddlewareHandlerFunc 系列中间件
func (mr *mRote) Use(ms ...middleware.MiddlewareHandlerFunc) *mRote {

	for _, m := range ms {
		mr.middlewares = append(mr.middlewares, m)
	}
	return mr
}

// UseHandlerMiddleware 使用 MiddlewareHandler 系列中间件
func (mr *mRote) UseHandlerMiddleware(ms ...middleware.MiddlewareHandler) *mRote {
	for _, m := range ms {
		mr.middlewares = append(mr.middlewares, m)
	}
	return mr
}

// Before 添加一个前置请求处理器
func (mr *mRote) Before(handler ...http.HandlerFunc) *mRote {

	for _, h := range handler {
		mr.before = append(mr.before, h)
	}

	return mr
}

// BeforeHandler 添加一个前置请求处理器
func (mr *mRote) BeforeHandler(handler ...http.Handler) *mRote {

	for _, h := range handler {
		mr.before = append(mr.before, h)
	}

	return mr
}

// AfterHandler 添加一个后置请求处理器
func (mr *mRote) After(handler ...http.HandlerFunc) *mRote {

	for _, h := range handler {
		mr.after = append(mr.after, h)
	}

	return mr
}

// AfterHandler 添加一个后置请求处理器
func (mr *mRote) AfterHandler(handler ...http.Handler) *mRote {

	for _, h := range handler {
		mr.after = append(mr.after, h)
	}

	return mr
}

// Bind 将请求数据绑定至 validateData，validateData 只能是对象指针或则 ValidateFunc, 返回的为当前路由的拷贝
func (mr *mRote) Bind(validateData interface{}) *mRote {
	return mr.Copy().BeforeHandler(middleware.Bind(validateData))
}

// Copy 获取一份当前配置的拷贝
func (mr *mRote) Copy() *mRote {

	_mr := &mRote{}
	for _, b := range mr.before {
		_mr.before = append(_mr.before, b)
	}

	for _, a := range mr.after {
		_mr.after = append(_mr.after, a)
	}

	for _, m := range mr.middlewares {
		_mr.middlewares = append(_mr.middlewares, m)
	}

	return _mr
}
