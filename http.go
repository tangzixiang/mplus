package mplus

import (
	"github.com/tangzixiang/mplus/mhttp"
)

type StatusMethodCallback = mhttp.StatusMethodCallback
type ResponseWriter = mhttp.ResponseWriter

var (
	EmptyRespData                     = mhttp.EmptyRespData
	DefaultMemorySize                 = mhttp.DefaultMemorySize
	SetDefaultMemorySize              = mhttp.SetDefaultMemorySize
	Abort                             = mhttp.Abort
	NotAbort                          = mhttp.NotAbort
	IsAbort                           = mhttp.IsAbort
	Error                             = mhttp.Error
	ErrorEmpty                        = mhttp.ErrorEmpty
	Plain                             = mhttp.Plain
	PlainEmpty                        = mhttp.PlainEmpty
	AbortError                        = mhttp.AbortError
	AbortEmptyError                   = mhttp.AbortEmptyError
	AbortPlain                        = mhttp.AbortPlain
	AbortEmptyPlain                   = mhttp.AbortEmptyPlain
	Redirect                          = mhttp.Redirect
	JSON                              = mhttp.JSON
	JSONOK                            = mhttp.JSONOK
	DumpRequest                       = mhttp.DumpRequest
	DumpRequestPure                   = mhttp.DumpRequestPure
	RegisterHttpStatusMethod          = mhttp.RegisterHttpStatusMethod
	NewResponseWrite                  = mhttp.NewResponseWrite
	GetHTTPRespStatus                 = mhttp.GetHTTPRespStatus
	SetHTTPRespStatus                 = mhttp.SetHTTPRespStatus
	UnWrapResponseWriter              = mhttp.UnWrapResponseWriter
	CopyRequest                       = mhttp.CopyRequest
	OK                                = mhttp.OK
	Created                           = mhttp.Created
	Accepted                          = mhttp.Accepted
	NonAuthoritativeInfo              = mhttp.NonAuthoritativeInfo
	NoContent                         = mhttp.NoContent
	ResetContent                      = mhttp.ResetContent
	PartialContent                    = mhttp.PartialContent
	MultiStatus                       = mhttp.MultiStatus
	AlreadyReported                   = mhttp.AlreadyReported
	IMUsed                            = mhttp.IMUsed
	MultipleChoices                   = mhttp.MultipleChoices
	MovedPermanently                  = mhttp.MovedPermanently
	Found                             = mhttp.Found
	SeeOther                          = mhttp.SeeOther
	NotModified                       = mhttp.NotModified
	UseProxy                          = mhttp.UseProxy
	TemporaryRedirect                 = mhttp.TemporaryRedirect
	PermanentRedirect                 = mhttp.PermanentRedirect
	BadRequest                        = mhttp.BadRequest
	Unauthorized                      = mhttp.Unauthorized
	PaymentRequired                   = mhttp.PaymentRequired
	Forbidden                         = mhttp.Forbidden
	NotFound                          = mhttp.NotFound
	MethodNotAllowed                  = mhttp.MethodNotAllowed
	NotAcceptable                     = mhttp.NotAcceptable
	ProxyAuthRequired                 = mhttp.ProxyAuthRequired
	RequestTimeout                    = mhttp.RequestTimeout
	Conflict                          = mhttp.Conflict
	Gone                              = mhttp.Gone
	LengthRequired                    = mhttp.LengthRequired
	PreconditionFailed                = mhttp.PreconditionFailed
	RequestEntityTooLarge             = mhttp.RequestEntityTooLarge
	RequestURITooLong                 = mhttp.RequestURITooLong
	UnsupportedMediaType              = mhttp.UnsupportedMediaType
	RequestedRangeNotSatisfiable      = mhttp.RequestedRangeNotSatisfiable
	ExpectationFailed                 = mhttp.ExpectationFailed
	Teapot                            = mhttp.Teapot
	MisdirectedRequest                = mhttp.MisdirectedRequest
	UnprocessableEntity               = mhttp.UnprocessableEntity
	Locked                            = mhttp.Locked
	FailedDependency                  = mhttp.FailedDependency
	TooEarly                          = mhttp.TooEarly
	UpgradeRequired                   = mhttp.UpgradeRequired
	PreconditionRequired              = mhttp.PreconditionRequired
	TooManyRequests                   = mhttp.TooManyRequests
	RequestHeaderFieldsTooLarge       = mhttp.RequestHeaderFieldsTooLarge
	UnavailableForLegalReasons        = mhttp.UnavailableForLegalReasons
	InternalServerError               = mhttp.InternalServerError
	NotImplemented                    = mhttp.NotImplemented
	BadGateway                        = mhttp.BadGateway
	ServiceUnavailable                = mhttp.ServiceUnavailable
	GatewayTimeout                    = mhttp.GatewayTimeout
	HTTPVersionNotSupported           = mhttp.HTTPVersionNotSupported
	VariantAlsoNegotiates             = mhttp.VariantAlsoNegotiates
	InsufficientStorage               = mhttp.InsufficientStorage
	LoopDetected                      = mhttp.LoopDetected
	NotExtended                       = mhttp.NotExtended
	NetworkAuthenticationRequired     = mhttp.NetworkAuthenticationRequired
	CallRegisterFuncOrAbortEmptyError = mhttp.CallRegisterFuncOrAbortEmptyError
	CallRegisterFuncOrAbortEmptyPlain = mhttp.CallRegisterFuncOrAbortEmptyPlain
	CallRegisterFuncOrAbortError      = mhttp.CallRegisterFuncOrAbortError
	CallRegisterFuncOrAbortPlain      = mhttp.CallRegisterFuncOrAbortPlain
)
