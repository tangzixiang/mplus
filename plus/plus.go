package plus

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/tangzixiang/mplus/context"
	"github.com/tangzixiang/mplus/header"
	"github.com/tangzixiang/mplus/message"
	"github.com/tangzixiang/mplus/mhttp"
	"github.com/tangzixiang/mplus/query"
)

// PP 缓存了当前请求的 w ResponseWriter 及 r Request
type PP struct {
	w http.ResponseWriter
	r *http.Request

	h http.Handler
	q url.Values
}

// PlusPlus 返回一个当前请求上下文托管对象
func PlusPlus(w http.ResponseWriter, r *http.Request) *PP {
	return &PP{w, r, nil, nil}
}

// Req 获取 Request
func (p *PP) Req() *http.Request {
	return p.r
}

// Resp 获取 ResponseWriter
func (p *PP) Resp() http.ResponseWriter {
	return p.w
}

// RequestID 获取 RequestID
func (p *PP) RequestID() string {
	return header.GetHeaderRequestID(p.r)
}

// VO 获取请求对象
func (p *PP) VO() interface{} {
	return context.GetContextValue(p.r.Context(), context.ReqData)
}

// Abort 将当前请求标识为中断
func (p *PP) Abort() *PP {
	mhttp.Abort(p.r)
	return p
}

// NotAbort 取消当前请求标识中断状态
func (p *PP) NotAbort() *PP {
	mhttp.NotAbort(p.r)
	return p
}

// IsAbort 判断当前请求是否中断
func (p *PP) IsAbort() bool {
	return mhttp.IsAbort(p.r)
}

// Error 响应异常信息
func (p *PP) Error(statusCode int, msg string) *PP {
	mhttp.Error(p.w, message.NewMessage(statusCode, msg))
	return p
}

// EmptyError 响应空的异常信息
func (p *PP) EmptyError(statusCode int) *PP {
	mhttp.ErrorEmpty(p.w, message.NewMessage(statusCode, ""))
	return p
}

// AbortEmptyError 将当前请求标识为中断并响应空的异常信息
func (p *PP) AbortEmptyError(statusCode int) *PP {
	mhttp.AbortEmptyError(p.w, p.r, message.NewMessage(statusCode, ""))
	return p
}

// ErrorMsg 响应异常信息
func (p *PP) ErrorMsg(message message.Message) *PP {
	mhttp.Error(p.w, message)
	return p
}

// AbortErrorMsg 将当前请求标识为中断并响应异常信息
func (p *PP) AbortErrorMsg(message message.Message) *PP {
	mhttp.AbortError(p.w, p.r, message)
	return p
}

// Plain 返回一个 text/plain 格式的响应
func (p *PP) Plain(statusCode int, msg string) *PP {
	mhttp.Plain(p.w, message.NewMessage(statusCode, msg))
	return p
}

// EmptyPlain 返回一个 text/plain 格式的响应，当前方法触发的请求无响应内容
func (p *PP) EmptyPlain(statusCode int) *PP {
	mhttp.PlainEmpty(p.w, message.NewMessage(statusCode, ""))
	return p
}

// Redirect 3xx 重定向,底层调用的是 http.Redirect
func (p *PP) Redirect(url string, statusCode int) *PP {
	mhttp.Redirect(p.w, p.r, url, statusCode)
	return p
}

// JSON 响应指定状态码的 JSON 数据
func (p *PP) JSON(data interface{}, status int) *PP {
	mhttp.JSON(p.w, p.r, data, status)
	return p
}

// JSONOK 响应指定状态码为 200 的 JSON 数据
func (p *PP) JSONOK(data interface{}) *PP {
	mhttp.JSONOK(p.w, p.r, data)
	return p
}

// OK 200 OK
func (p *PP) OK() *PP {
	mhttp.OK(p.w, p.r)
	return p
}

// Created 201 Created
func (p *PP) Created() *PP {
	mhttp.Created(p.w, p.r)
	return p
}

// Accepted 202 Accepted
func (p *PP) Accepted() *PP {
	mhttp.Accepted(p.w, p.r)
	return p
}

// NonAuthoritativeInfo 203 Non-Authoritative Information
func (p *PP) NonAuthoritativeInfo() *PP {
	mhttp.NonAuthoritativeInfo(p.w, p.r)
	return p
}

// NoContent 204 No Content
func (p *PP) NoContent() *PP {
	mhttp.NoContent(p.w, p.r)
	return p
}

// ResetContent 205 Reset Content
func (p *PP) ResetContent() *PP {
	mhttp.ResetContent(p.w, p.r)
	return p
}

// PartialContent 206 Partial Content
func (p *PP) PartialContent() *PP {
	mhttp.PartialContent(p.w, p.r)
	return p
}

// MultiStatus 207 Multi-Status
func (p *PP) MultiStatus() *PP {
	mhttp.MultiStatus(p.w, p.r)
	return p
}

// AlreadyReported 208 Already Reported
func (p *PP) AlreadyReported() *PP {
	mhttp.AlreadyReported(p.w, p.r)
	return p
}

// IMUsed 226 IM Used
func (p *PP) IMUsed() *PP {
	mhttp.IMUsed(p.w, p.r)
	return p
}

// MultipleChoices 300 Multiple Choices
func (p *PP) MultipleChoices() *PP {
	mhttp.MultipleChoices(p.w, p.r)
	return p
}

// MovedPermanently 301 Moved Permanently
func (p *PP) MovedPermanently() *PP {
	mhttp.MovedPermanently(p.w, p.r)
	return p
}

// Found 302 Found
func (p *PP) Found() *PP {
	mhttp.Found(p.w, p.r)
	return p
}

// SeeOther 303 See Other
func (p *PP) SeeOther() *PP {
	mhttp.SeeOther(p.w, p.r)
	return p
}

// NotModified 304 Not Modifie
func (p *PP) NotModified() *PP {
	mhttp.NotModified(p.w, p.r)
	return p
}

// UseProxy 305 Use Proxy
func (p *PP) UseProxy() *PP {
	mhttp.UseProxy(p.w, p.r)
	return p
}

// TemporaryRedirect 307 Temporary Redirect
func (p *PP) TemporaryRedirect() *PP {
	mhttp.TemporaryRedirect(p.w, p.r)
	return p
}

// PermanentRedirect 308 Permanent Redirect
func (p *PP) PermanentRedirect() *PP {
	mhttp.PermanentRedirect(p.w, p.r)
	return p
}

// BadRequest 400 Bad Request
func (p *PP) BadRequest() *PP {
	mhttp.BadRequest(p.w, p.r)
	return p
}

// Unauthorized 401 Unauthorized
func (p *PP) Unauthorized() *PP {
	mhttp.Unauthorized(p.w, p.r)
	return p
}

// PaymentRequired 402 Payment Required
func (p *PP) PaymentRequired() *PP {
	mhttp.PaymentRequired(p.w, p.r)
	return p
}

// Forbidden 403 Forbidden
func (p *PP) Forbidden() *PP {
	mhttp.Forbidden(p.w, p.r)
	return p
}

// NotFound 404 Not Found
func (p *PP) NotFound() *PP {
	mhttp.NotFound(p.w, p.r)
	return p
}

// MethodNotAllowed 405 Method Not Allowed
func (p *PP) MethodNotAllowed() *PP {
	mhttp.MethodNotAllowed(p.w, p.r)
	return p
}

// NotAcceptable 406 Not Acceptable
func (p *PP) NotAcceptable() *PP {
	mhttp.NotAcceptable(p.w, p.r)
	return p
}

// ProxyAuthRequired 407 Proxy Authentication Required
func (p *PP) ProxyAuthRequired() *PP {
	mhttp.ProxyAuthRequired(p.w, p.r)
	return p
}

// RequestTimeout 408 Request Timeout
func (p *PP) RequestTimeout() *PP {
	mhttp.RequestTimeout(p.w, p.r)
	return p
}

// Conflict 409 Conflict
func (p *PP) Conflict() *PP {
	mhttp.Conflict(p.w, p.r)
	return p
}

// Gone 410 Gone
func (p *PP) Gone() *PP {
	mhttp.Gone(p.w, p.r)
	return p
}

// LengthRequired 411 Length Required
func (p *PP) LengthRequired() *PP {
	mhttp.LengthRequired(p.w, p.r)
	return p
}

// PreconditionFailed 412 Precondition Failed
func (p *PP) PreconditionFailed() *PP {
	mhttp.PreconditionFailed(p.w, p.r)
	return p
}

// RequestEntityTooLarge 413 Request Entity Too Large
func (p *PP) RequestEntityTooLarge() *PP {
	mhttp.RequestEntityTooLarge(p.w, p.r)
	return p
}

// RequestURITooLong 414 Request URI Too Long
func (p *PP) RequestURITooLong() *PP {
	mhttp.RequestURITooLong(p.w, p.r)
	return p
}

// UnsupportedMediaType 415 Unsupported Media Type
func (p *PP) UnsupportedMediaType() *PP {
	mhttp.UnsupportedMediaType(p.w, p.r)
	return p
}

// RequestedRangeNotSatisfiable 416 Requested Range Not Satisfiable
func (p *PP) RequestedRangeNotSatisfiable() *PP {
	mhttp.RequestedRangeNotSatisfiable(p.w, p.r)
	return p
}

// ExpectationFailed 417 Expectation Failed
func (p *PP) ExpectationFailed() *PP {
	mhttp.ExpectationFailed(p.w, p.r)
	return p
}

// Teapot 418 I'm a teapot
func (p *PP) Teapot() *PP {
	mhttp.Teapot(p.w, p.r)
	return p
}

// MisdirectedRequest 421 Misdirected Request
func (p *PP) MisdirectedRequest() *PP {
	mhttp.MisdirectedRequest(p.w, p.r)
	return p
}

// UnprocessableEntity 422 Unprocessable Entity
func (p *PP) UnprocessableEntity() *PP {
	mhttp.UnprocessableEntity(p.w, p.r)
	return p
}

// Locked 423 Locked
func (p *PP) Locked() *PP {
	mhttp.Locked(p.w, p.r)
	return p
}

// FailedDependency 424 Failed Dependency
func (p *PP) FailedDependency() *PP {
	mhttp.FailedDependency(p.w, p.r)
	return p
}

// TooEarly 425 Too Early
func (p *PP) TooEarly() *PP {
	mhttp.TooEarly(p.w, p.r)
	return p
}

// UpgradeRequired 426 Upgrade Required
func (p *PP) UpgradeRequired() *PP {
	mhttp.UpgradeRequired(p.w, p.r)
	return p
}

// PreconditionRequired 428 Precondition Required
func (p *PP) PreconditionRequired() *PP {
	mhttp.PreconditionRequired(p.w, p.r)
	return p
}

// TooManyRequests 429 Too Many Requests
func (p *PP) TooManyRequests() *PP {
	mhttp.TooManyRequests(p.w, p.r)
	return p
}

// RequestHeaderFieldsTooLarge 431 Request Header Fields Too Large
func (p *PP) RequestHeaderFieldsTooLarge() *PP {
	mhttp.RequestHeaderFieldsTooLarge(p.w, p.r)
	return p
}

// UnavailableForLegalReasons 451 Unavailable For Legal Reasons
func (p *PP) UnavailableForLegalReasons() *PP {
	mhttp.UnavailableForLegalReasons(p.w, p.r)
	return p
}

// InternalServerError 500 Internal Server Error
func (p *PP) InternalServerError() *PP {
	mhttp.InternalServerError(p.w, p.r)
	return p
}

// NotImplemented 501 Not Implemented
func (p *PP) NotImplemented() *PP {
	mhttp.NotImplemented(p.w, p.r)
	return p
}

// BadGateway 502 Bad Gateway
func (p *PP) BadGateway() *PP {
	mhttp.BadGateway(p.w, p.r)
	return p
}

// ServiceUnavailable 503 Service Unavailable
func (p *PP) ServiceUnavailable() *PP {
	mhttp.ServiceUnavailable(p.w, p.r)
	return p
}

// GatewayTimeout 504 Gateway Timeout
func (p *PP) GatewayTimeout() *PP {
	mhttp.GatewayTimeout(p.w, p.r)
	return p
}

// HTTPVersionNotSupported 505 HTTP Version Not Supported
func (p *PP) HTTPVersionNotSupported() *PP {
	mhttp.HTTPVersionNotSupported(p.w, p.r)
	return p
}

// VariantAlsoNegotiates 506 Variant Also Negotiates
func (p *PP) VariantAlsoNegotiates() *PP {
	mhttp.VariantAlsoNegotiates(p.w, p.r)
	return p
}

// InsufficientStorage 507 Insufficient Storage
func (p *PP) InsufficientStorage() *PP {
	mhttp.InsufficientStorage(p.w, p.r)
	return p
}

// LoopDetected 508 Loop Detected
func (p *PP) LoopDetected() *PP {
	mhttp.LoopDetected(p.w, p.r)
	return p
}

// NotExtended  510 Not Extended
func (p *PP) NotExtended() *PP {
	mhttp.NotExtended(p.w, p.r)
	return p
}

// NetworkAuthenticationRequired 511 Network Authentication Required
func (p *PP) NetworkAuthenticationRequired() *PP {
	mhttp.NetworkAuthenticationRequired(p.w, p.r)
	return p
}

// Get 获取上下文内容
func (p *PP) Get(key string) interface{} {
	return context.GetContextValue(p.r.Context(), key)
}

// Set 设置上下文内容
func (p *PP) Set(key string, value interface{}) *PP {
	context.SetContextValue(p.r.Context(), key, value)
	return p
}

// SetR 设置上下文内容,并返回原内容
func (p *PP) SetR(key string, value interface{}) interface{} {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetStringR 设置上下文内容,并返回原内容
func (p *PP) SetStringR(key string, value string) string {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetIntR 设置上下文内容,并返回原内容
func (p *PP) SetIntR(key string, value int) int {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt8R 设置上下文内容,并返回原内容
func (p *PP) SetInt8R(key string, value int8) int8 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt16R 设置上下文内容,并返回原内容
func (p *PP) SetInt16R(key string, value int16) int16 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt32R 设置上下文内容,并返回原内容
func (p *PP) SetInt32R(key string, value int32) int32 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt64R 设置上下文内容,并返回原内容
func (p *PP) SetInt64R(key string, value int64) int64 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUintR 设置上下文内容,并返回原内容
func (p *PP) SetUintR(key string, value uint) uint {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint8R 设置上下文内容,并返回原内容
func (p *PP) SetUint8R(key string, value uint8) uint8 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint16R 设置上下文内容,并返回原内容
func (p *PP) SetUint16R(key string, value uint16) uint16 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint32R 设置上下文内容,并返回原内容
func (p *PP) SetUint32R(key string, value uint32) uint32 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint64R 设置上下文内容,并返回原内容
func (p *PP) SetUint64R(key string, value uint64) uint64 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetBoolR 设置上下文内容,并返回原内容
func (p *PP) SetBoolR(key string, value bool) bool {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetByteR 设置上下文内容,并返回原内容
func (p *PP) SetByteR(key string, value byte) byte {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetBytesR 设置上下文内容,并返回原内容
func (p *PP) SetBytesR(key string, value []byte) []byte {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetTimeR 设置上下文内容,并返回原内容
func (p *PP) SetTimeR(key string, value time.Time) time.Time {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetFloat32R 设置上下文内容,并返回原内容
func (p *PP) SetFloat32R(key string, value float32) float32 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// SetFloat64R 设置上下文内容,并返回原内容
func (p *PP) SetFloat64R(key string, value float64) float64 {
	context.SetContextValue(p.r.Context(), key, value)
	return value
}

// GetString 获取 string 类型的上下文内容，获取失败返回其零值
func (p *PP) GetString(key string) string {
	return context.GetContextValueString(p.r.Context(), key)
}

// GetInt 获取 int 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt(key string) int {
	return context.GetContextValueInt(p.r.Context(), key)
}

// GetInt8 获取 int8 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt8(key string) int8 {
	return context.GetContextValueInt8(p.r.Context(), key)
}

// GetInt16 获取 int16 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt16(key string) int16 {
	return context.GetContextValueInt16(p.r.Context(), key)
}

// GetInt32 获取 int32 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt32(key string) int32 {
	return context.GetContextValueInt32(p.r.Context(), key)
}

// GetInt64 获取 int64 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt64(key string) int64 {
	return context.GetContextValueInt64(p.r.Context(), key)
}

// GetUInt 获取 uint 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt(key string) uint {
	return context.GetContextValueUInt(p.r.Context(), key)
}

// GetUInt8 获取 uint8 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt8(key string) uint8 {
	return context.GetContextValueUInt8(p.r.Context(), key)
}

// GetUInt16 获取 uint16 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt16(key string) uint16 {
	return context.GetContextValueUInt16(p.r.Context(), key)
}

// GetUInt32 获取 uint32 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt32(key string) uint32 {
	return context.GetContextValueUInt32(p.r.Context(), key)
}

// GetUInt64 获取 uint64 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt64(key string) uint64 {
	return context.GetContextValueUInt64(p.r.Context(), key)
}

// GetBool 获取 bool 类型的上下文内容，获取失败返回其零值
func (p *PP) GetBool(key string) bool {
	return context.GetContextValueBool(p.r.Context(), key)
}

// GetByte 获取 byte 类型的上下文内容，获取失败返回其零值
func (p *PP) GetByte(key string) byte {
	return context.GetContextValueByte(p.r.Context(), key)
}

// GetBytes 获取 []byte 类型的上下文内容，获取失败返回其零值
func (p *PP) GetBytes(key string) []byte {
	return context.GetContextValueBytes(p.r.Context(), key)
}

// GetTime 获取 time.Time 类型的上下文内容，获取失败返回其零值
func (p *PP) GetTime(key string) time.Time {
	return context.GetContextValueTime(p.r.Context(), key)
}

// GetUInt32 获取 float32 类型的上下文内容，获取失败返回其零值
func (p *PP) GetFloat32(key string) float32 {
	return context.GetContextValueFloat32(p.r.Context(), key)
}

// GetFloat64 获取 float64 类型的上下文内容，获取失败返回其零值
func (p *PP) GetFloat64(key string) float64 {
	return context.GetContextValueFloat64(p.r.Context(), key)
}

// GetDf 获取上下文内容，获取失败返回默认值
func (p *PP) GetDf(key string, defaultValue interface{}) interface{} {
	rv := context.GetContextValue(p.r.Context(), key)
	if rv != nil {
		return rv
	}
	return defaultValue
}

// GetStringDf 获取 string 类型的上下文内容，获取失败返回默认值
func (p *PP) GetStringDf(key string, defaultValue string) string {
	return context.GetContextValueString(p.r.Context(), key, defaultValue)
}

// GetIntDf 获取 int 类型的上下文内容，获取失败返回默认值
func (p *PP) GetIntDf(key string, defaultValue int) int {
	return context.GetContextValueInt(p.r.Context(), key, defaultValue)
}

// GetInt8Df 获取 int8 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt8Df(key string, defaultValue int8) int8 {
	return context.GetContextValueInt8(p.r.Context(), key, defaultValue)
}

// GetInt16Df 获取 int16 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt16Df(key string, defaultValue int16) int16 {
	return context.GetContextValueInt16(p.r.Context(), key, defaultValue)
}

// GetInt32Df 获取 int32 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt32Df(key string, defaultValue int32) int32 {
	return context.GetContextValueInt32(p.r.Context(), key, defaultValue)
}

// GetInt64Df 获取 int64 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt64Df(key string, defaultValue int64) int64 {
	return context.GetContextValueInt64(p.r.Context(), key, defaultValue)
}

// GetUIntDf 获取 uint 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUIntDf(key string, defaultValue uint) uint {
	return context.GetContextValueUInt(p.r.Context(), key, defaultValue)
}

// GetUInt8Df 获取 uint8 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt8Df(key string, defaultValue uint8) uint8 {
	return context.GetContextValueUInt8(p.r.Context(), key, defaultValue)
}

// GetUInt16Df 获取 uint16 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt16Df(key string, defaultValue uint16) uint16 {
	return context.GetContextValueUInt16(p.r.Context(), key, defaultValue)
}

// GetUInt32Df 获取 uint32 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt32Df(key string, defaultValue uint32) uint32 {
	return context.GetContextValueUInt32(p.r.Context(), key, defaultValue)
}

// GetUInt64Df 获取 uint64 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt64Df(key string, defaultValue uint64) uint64 {
	return context.GetContextValueUInt64(p.r.Context(), key, defaultValue)
}

// GetBoolDf 获取 bool 类型的上下文内容，获取失败返回默认值
func (p *PP) GetBoolDf(key string, defaultValue bool) bool {
	return context.GetContextValueBool(p.r.Context(), key, defaultValue)
}

// GetByteDf 获取 byte 类型的上下文内容，获取失败返回默认值
func (p *PP) GetByteDf(key string, defaultValue byte) byte {
	return context.GetContextValueByte(p.r.Context(), key, defaultValue)
}

// GetBytesDf 获取 []byte 类型的上下文内容，获取失败返回默认值
func (p *PP) GetBytesDf(key string, defaultValue []byte) []byte {
	return context.GetContextValueBytes(p.r.Context(), key, defaultValue)
}

// GetTimeDf 获取 time.Time 类型的上下文内容，获取失败返回默认值
func (p *PP) GetTimeDf(key string, defaultValue time.Time) time.Time {
	return context.GetContextValueTime(p.r.Context(), key, defaultValue)
}

// GetUInt32 获取 float32 类型的上下文内容，获取失败返回默认值
func (p *PP) GetFloat32Df(key string, defaultValue float32) float32 {
	return context.GetContextValueFloat32(p.r.Context(), key, defaultValue)
}

// GetFloat64 获取 float64 类型的上下文内容，获取失败返回默认值
func (p *PP) GetFloat64Df(key string, defaultValue float64) float64 {
	return context.GetContextValueFloat64(p.r.Context(), key, defaultValue)
}

// DoCallback 查询指定错误嘛注册的回调并执行
func (p *PP) CallbackByCode(errorCode int, respData interface{}) *PP {
	mByCode := message.Messages.Get(errorCode)

	if mByCode != nil {
		if m, ok := mByCode.(message.Callback); ok {
			m.Do(p.w, p.r, mByCode, respData)
			return p
		}

		switch mByCode.Status() {
		case http.StatusOK:
			p.JSONOK(respData)
		case http.StatusBadRequest:
			p.BadRequest()

		}
	}

	return p
}

// Queries 获取 URL 上的请求字段
func (p *PP) Queries() url.Values {
	if p.q == nil {
		p.q = query.Queries(p.r)
	}
	return p.q
}

// GetQuery 获取 URL 上的指定的请求字段
func (p *PP) GetQuery(key string) (value string, exists bool) {
	if p.q == nil {
		p.q = query.Queries(p.r)
	}
	values, exists := p.q[key]
	if !exists || len(values) <= 0 {
		return "", false
	}

	return values[0], true
}

// GetQuery 获取 URL 上的指定的请求字段
func (p *PP) Query(key string) (value string) {
	if p.q == nil {
		p.q = query.Queries(p.r)
	}

	values, exists := p.q[key]
	if !exists || len(values) <= 0 {
		return ""
	}

	return values[0]
}

// GetQuery 获取 URL 上的指定的请求字段，获取失败返回默认值
func (p *PP) GetQueryDf(key, defaultValue string) (value string) {
	if p.q == nil {
		p.q = query.Queries(p.r)
	}

	values, exists := p.q[key]
	if !exists || len(values) <= 0 {
		return defaultValue
	}

	return values[0]
}

// GetHeader 获取指定请求头信息
func (p *PP) GetHeader(key string) string {
	return header.GetHeader(p.r, key)
}

// GetHeaderDf 获取指定请求头信息，获取失败返回默认值
func (p *PP) GetHeaderDf(key string, defaultValue string) string {
	value := header.GetHeader(p.r, key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetHeaderValues 获取指定请求头信息集
func (p *PP) GetHeaderValues(key string) []string {
	return header.GetHeaderValues(p.r, key)
}

// SplitHeader 分割指定请求头内容
func (p *PP) SplitHeader(key string, splitSep string) []string {
	return header.SplitHeader(p.r, key, splitSep)
}

// WriteHeader 将指定头信息写入到请求头中
func (p *PP) WriteHeader(key, value string) *PP {
	header.SetRequestHeader(p.r, key, value)
	return p
}

// WriteHeaders 将指定头信息写入到请求头中
func (p *PP) WriteHeaders(headers map[string]string) *PP {
	header.SetRequestHeaders(p.r, headers)
	return p
}

// AppendHeader 将指定头信息追加写入到请求头中
func (p *PP) AppendHeader(key string, values ...string) *PP {
	header.AddRequestHeader(p.r, key, values...)
	return p
}

// AppendHeaders 将指定头信息追加写入到请求头中
func (p *PP) AppendHeaders(headers map[string]string) *PP {
	header.AddRequestHeaders(p.r, headers)
	return p
}

// GetRespHeader 获取指定响应头信息
func (p *PP) GetRespHeader(key string) string {
	return header.GetResponseHeader(p.w, key)
}

// GetRespHeaderDf 获取指定响应头信息，获取失败返回默认值
func (p *PP) GetRespHeaderDf(key string, defaultValue string) string {
	value := header.GetResponseHeader(p.w, key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetRespHeaderValues 获取指定响应头信息集
func (p *PP) GetRespHeaderValues(key string) []string {
	return header.GetResponseHeaderValues(p.w, key)
}

// WriteRespHeader 将指定头信息写入到响应头中
func (p *PP) WriteRespHeader(key, value string) *PP {
	header.SetResponseHeader(p.w, key, value)
	return p
}

// WriteRespHeaders 将指定头信息写入到响应头中
func (p *PP) WriteRespHeaders(headers map[string]string) *PP {
	header.SetResponseHeaders(p.w, headers)
	return p
}

// AppendRespHeader 将指定头信息追加写入到响应头中
func (p *PP) AppendRespHeader(key string, values ...string) *PP {
	header.AddResponseHeader(p.w, key, values...)
	return p
}

// AppendRespHeaders 将指定头信息追加写入到响应头中
func (p *PP) AppendRespHeaders(headers map[string]string) *PP {
	header.AddResponseHeaders(p.w, headers)
	return p
}

// ReqHeader 获取请求头
func (p *PP) ReqHeader() http.Header {
	return p.r.Header
}

// ReqURL 获取请求的 URL
func (p *PP) ReqURL() *url.URL {
	return p.r.URL
}

// ReqHost 获取请求的 Host
func (p *PP) ReqHost() string {
	return p.r.Host
}

// Handler 缓存当前请求的 Handler
//
// mplus.PlusPlus(w,r).Handler(handlerFunc).ServeHTTP()
func (p *PP) Handler(h http.Handler) *PP {
	p.h = h
	return p
}

// ServeHTTP 回调执行缓存的 handlerFunc
func (p *PP) ServeHTTP() *PP {
	if p.h != nil {
		p.h.ServeHTTP(p.w, p.r)
	}
	return p
}

// FormFile returns the first file for the provided form key.
func (p *PP) FormFile(name string) (*multipart.FileHeader, error) {
	_, fh, err := p.r.FormFile(name)
	return fh, err
}

// Status 设置请求响应状态码
func (p *PP) Status(statusCode int) *PP {
	mhttp.SetHTTPRespStatus(p.w, statusCode)
	return p
}

// GetStatus 获取响应状态码
func (p *PP) GetStatus() int {
	return mhttp.GetHTTPRespStatus(p.w)
}

// CopyReq 拷贝一个请求，body 不会被拷贝，因为 body 是一个数据流
func (p *PP) CopyReq() *http.Request {
	return mhttp.CopyRequest(p.r)
}

// Method 获取请求方式
func (p *PP) Method() string {
	return p.r.Method
}

// GetClientIP 获取客户端 IP 地址
func (p *PP) GetClientIP() string {
	return header.GetClientIP(p.r)
}

// ReqBody 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用
func (p *PP) ReqBody() string {
	return mhttp.DumpRequest(p.r)
}

// ReqBody 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用
func (p *PP) ReqBodyPure() []byte {
	return mhttp.DumpRequestPure(p.r)
}

// ReqBodyMap 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用,body 内容会被序列化成 map[string] interface{}
func (p *PP) ReqBodyMap() (map[string]interface{}, error) {
	body := mhttp.DumpRequestPure(p.r)
	m := map[string]interface{}{}

	return m, json.Unmarshal(body, &m)
}

// ReqBodyMap 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用,body 内容会被序列化至 unmarshaler
func (p *PP) ReqBodyToUnmarshaler(unmarshaler json.Unmarshaler) error {
	return unmarshaler.UnmarshalJSON(mhttp.DumpRequestPure(p.r))
}
