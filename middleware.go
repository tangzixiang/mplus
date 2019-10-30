package mplus

import (
	"github.com/tangzixiang/mplus/middleware"
)

type MiddlewareHandler = middleware.MiddlewareHandler
type MiddlewareHandlerFunc = middleware.MiddlewareHandlerFunc

var (
	PreMiddleware              = middleware.Pre
	PreHandlerMiddleware       = middleware.PreHandler
	RequestIDMiddleware        = middleware.RequestID
	RequestIDHandlerMiddleware = middleware.RequestIDHandler
	Use                        = middleware.Use
	UseHandlerMiddleware       = middleware.UseHandlerMiddleware
	Thunk                      = middleware.Thunk
	ThunkHandler               = middleware.ThunkHandler
	Bind                       = middleware.Bind
)
