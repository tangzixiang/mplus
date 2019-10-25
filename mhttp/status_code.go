package mhttp

import (
	"net/http"
	"sync"

	"github.com/tangzixiang/mplus/message"
)

// OK 200 OK
func OK(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusOK, http.StatusOK)
}

// Created 201 Created
func Created(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusCreated, http.StatusCreated)
}

// Accepted 202 Accepted
func Accepted(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusAccepted, http.StatusAccepted)
}

// NonAuthoritativeInfo 203 Non-Authoritative Information
func NonAuthoritativeInfo(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusNonAuthoritativeInfo, http.StatusNonAuthoritativeInfo)
}

// NoContent 204 No Content
func NoContent(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusNoContent, http.StatusNoContent)
}

// ResetContent 205 Reset Content
func ResetContent(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusResetContent, http.StatusResetContent)
}

// PartialContent 206 Partial Content
func PartialContent(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusPartialContent, http.StatusPartialContent)
}

// MultiStatus 207 Multi-Status
func MultiStatus(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusMultiStatus, http.StatusMultiStatus)
}

// AlreadyReported 208 Already Reported
func AlreadyReported(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusAlreadyReported, http.StatusAlreadyReported)
}

// IMUsed 226 IM Used
func IMUsed(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusIMUsed, http.StatusIMUsed)
}

// MultipleChoices 300 Multiple Choices
func MultipleChoices(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusMultipleChoices, http.StatusMultipleChoices)
}

// MovedPermanently 301 Moved Permanently
func MovedPermanently(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusMovedPermanently, http.StatusMovedPermanently)
}

// Found 302 Found
func Found(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusFound, http.StatusFound)
}

// SeeOther 303 See Other
func SeeOther(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusSeeOther, http.StatusSeeOther)
}

// NotModified 304 Not Modifie
func NotModified(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusNotModified, http.StatusNotModified)
}

// UseProxy 305 Use Proxy
func UseProxy(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusUseProxy, http.StatusUseProxy)
}

// TemporaryRedirect 307 Temporary Redirect
func TemporaryRedirect(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusTemporaryRedirect, http.StatusTemporaryRedirect)
}

// PermanentRedirect 308 Permanent Redirect
func PermanentRedirect(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyPlain(w, r, message.MessageStatusPermanentRedirect, http.StatusPermanentRedirect)
}

// BadRequest 400 Bad Request
func BadRequest(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusBadRequest, http.StatusBadRequest)
}

// Unauthorized 401 Unauthorized
func Unauthorized(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusUnauthorized, http.StatusUnauthorized)
}

// PaymentRequired 402 Payment Required
func PaymentRequired(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusPaymentRequired, http.StatusPaymentRequired)
}

// Forbidden 403 Forbidden
func Forbidden(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusForbidden, http.StatusForbidden)
}

// NotFound 404 Not Found
func NotFound(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusNotFound, http.StatusNotFound)
}

// MethodNotAllowed 405 Method Not Allowed
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusMethodNotAllowed, http.StatusMethodNotAllowed)
}

// NotAcceptable 406 Not Acceptable
func NotAcceptable(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusNotAcceptable, http.StatusNotAcceptable)
}

// ProxyAuthRequired 407 Proxy Authentication Required
func ProxyAuthRequired(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusProxyAuthRequired, http.StatusProxyAuthRequired)
}

// RequestTimeout 408 Request Timeout
func RequestTimeout(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusRequestTimeout, http.StatusRequestTimeout)
}

// Conflict 409 Conflict
func Conflict(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusConflict, http.StatusConflict)
}

// Gone 410 Gone
func Gone(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusGone, http.StatusGone)
}

// LengthRequired 411 Length Required
func LengthRequired(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusLengthRequired, http.StatusLengthRequired)
}

// PreconditionFailed 412 Precondition Failed
func PreconditionFailed(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusPreconditionFailed, http.StatusPreconditionFailed)
}

// RequestEntityTooLarge 413 Request Entity Too Large
func RequestEntityTooLarge(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusRequestEntityTooLarge, http.StatusRequestEntityTooLarge)
}

// RequestURITooLong 414 Request URI Too Long
func RequestURITooLong(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusRequestURITooLong, http.StatusRequestURITooLong)
}

// UnsupportedMediaType 415 Unsupported Media Type
func UnsupportedMediaType(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusUnsupportedMediaType, http.StatusUnsupportedMediaType)
}

// RequestedRangeNotSatisfiable 416 Requested Range Not Satisfiable
func RequestedRangeNotSatisfiable(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusRequestedRangeNotSatisfiable, http.StatusRequestedRangeNotSatisfiable)
}

// ExpectationFailed 417 Expectation Failed
func ExpectationFailed(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusExpectationFailed, http.StatusExpectationFailed)
}

// Teapot 418 I'm a teapot
func Teapot(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusTeapot, http.StatusTeapot)
}

// MisdirectedRequest 421 Misdirected Request
func MisdirectedRequest(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusMisdirectedRequest, http.StatusMisdirectedRequest)
}

// UnprocessableEntity 422 Unprocessable Entity
func UnprocessableEntity(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusUnprocessableEntity, http.StatusUnprocessableEntity)
}

// Locked 423 Locked
func Locked(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusLocked, http.StatusLocked)
}

// FailedDependency 424 Failed Dependency
func FailedDependency(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusFailedDependency, http.StatusFailedDependency)
}

// TooEarly 425 Too Early
func TooEarly(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusTooEarly, http.StatusTooEarly)
}

// UpgradeRequired 426 Upgrade Required
func UpgradeRequired(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusUpgradeRequired, http.StatusUpgradeRequired)
}

// PreconditionRequired 428 Precondition Required
func PreconditionRequired(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusPreconditionRequired, http.StatusPreconditionRequired)
}

// TooManyRequests 429 Too Many Requests
func TooManyRequests(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusTooManyRequests, http.StatusTooManyRequests)
}

// RequestHeaderFieldsTooLarge 431 Request Header Fields Too Large
func RequestHeaderFieldsTooLarge(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusRequestHeaderFieldsTooLarge, http.StatusRequestHeaderFieldsTooLarge)
}

// UnavailableForLegalReasons 451 Unavailable For Legal Reasons
func UnavailableForLegalReasons(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusUnavailableForLegalReasons, http.StatusUnavailableForLegalReasons)
}

// InternalServerError 500 Internal Server Error
func InternalServerError(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusInternalServerError, http.StatusInternalServerError)
}

// NotImplemented 501 Not Implemented
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusNotImplemented, http.StatusNotImplemented)
}

// BadGateway 502 Bad Gateway
func BadGateway(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusBadGateway, http.StatusBadGateway)
}

// ServiceUnavailable 503 Service Unavailable
func ServiceUnavailable(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusServiceUnavailable, http.StatusServiceUnavailable)
}

// GatewayTimeout 504 Gateway Timeout
func GatewayTimeout(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusGatewayTimeout, http.StatusGatewayTimeout)
}

// HTTPVersionNotSupported 505 HTTP Version Not Supported
func HTTPVersionNotSupported(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusHTTPVersionNotSupported, http.StatusHTTPVersionNotSupported)
}

// VariantAlsoNegotiates 506 Variant Also Negotiates
func VariantAlsoNegotiates(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusVariantAlsoNegotiates, http.StatusVariantAlsoNegotiates)
}

// InsufficientStorage 507 Insufficient Storage
func InsufficientStorage(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusInsufficientStorage, http.StatusInsufficientStorage)
}

// LoopDetected 508 Loop Detected
func LoopDetected(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusLoopDetected, http.StatusLoopDetected)
}

// NotExtended  510 Not Extended
func NotExtended(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusNotExtended, http.StatusNotExtended)
}

// NetworkAuthenticationRequired 511 Network Authentication Required
func NetworkAuthenticationRequired(w http.ResponseWriter, r *http.Request) {
	CallRegisterFuncOrAbortEmptyError(w, r, message.MessageStatusNetworkAuthenticationRequired, http.StatusNetworkAuthenticationRequired)
}

// CallRegisterFuncOrAbortEmptyError 调用已注册的状态回调，状态回调不存在则使用默认方式终止请求链
func CallRegisterFuncOrAbortEmptyError(w http.ResponseWriter, r *http.Request, m message.Message, statusCode int) {

	registeredFunc, exists := httpStatusMethodHub[statusCode]

	if !exists {
		AbortEmptyError(w, r, m)
		return
	}

	registeredFunc(w, Abort(r), m, statusCode)
}


// CallRegisterFuncOrAbortEmptyPlain 调用已注册的状态回调，状态回调不存在则使用默认方式终止请求链
func CallRegisterFuncOrAbortEmptyPlain(w http.ResponseWriter, r *http.Request, m message.Message, statusCode int) {

	registeredFunc, exists := httpStatusMethodHub[statusCode]

	if !exists {
		AbortEmptyPlain(w, r, m)
		return
	}

	registeredFunc(w, Abort(r), m, statusCode)
}

// CallRegisterFuncOrAbortError 调用已注册的状态回调，状态回调不存在则使用默认方式终止请求链，resp body 内容取决于 m 持有的信息
func CallRegisterFuncOrAbortError(w http.ResponseWriter, r *http.Request, m message.Message, statusCode int) {

	registeredFunc, exists := httpStatusMethodHub[statusCode]

	if !exists {
		AbortError(w, r, m)
		return
	}

	registeredFunc(w, Abort(r), m, statusCode)
}

// CallRegisterFuncOrAbortPlain 调用已注册的状态回调，状态回调不存在则使用默认方式终止请求链，resp body 内容取决于 m 持有的信息
func CallRegisterFuncOrAbortPlain(w http.ResponseWriter, r *http.Request, m message.Message, statusCode int) {

	registeredFunc, exists := httpStatusMethodHub[statusCode]

	if !exists {
		AbortPlain(w, r, m)
		return
	}

	registeredFunc(w, Abort(r), m, statusCode)
}

type StatusMethodCallback func(w http.ResponseWriter, r *http.Request, m message.Message, statusCode int)

var (
	httpStatusMethodHubLock sync.Mutex
	httpStatusMethodHub     = map[int]StatusMethodCallback{}
)

// RegisterHttpStatusMethod 注册默认的请求状态回调
func RegisterHttpStatusMethod(statusCode int, f StatusMethodCallback) {
	httpStatusMethodHubLock.Lock()
	httpStatusMethodHub[statusCode] = f
	httpStatusMethodHubLock.Unlock()
}


