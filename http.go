package mplus

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	requestAbortKey = "__abort"
	responseStatus  = "__status"
)

const defaultMemory = 32 * 1024 * 1024

// EmptyRespData 空响应体
var EmptyRespData = map[string]interface{}{}

// Abort 标识当前请求链已中断
func Abort(r *http.Request) *http.Request {
	SetContextValue(r.Context(), requestAbortKey, true)
	return r
}

// IsAbort 判断当前请求链是否已中断
func IsAbort(r *http.Request) bool {
	return GetContextValueBool(r.Context(), requestAbortKey)
}

// Error 返回异常信息，当前方法触发的请求响应内容将是文本格式
func Error(w http.ResponseWriter, message Message) {
	http.Error(SetHTTPRespStatus(w, message.Status(), false), message.Default(), message.Status())
}

// AbortError 终止请求链并返回异常信息，当前方法触发的请求响应内容将是文本格式
func AbortError(w http.ResponseWriter, r *http.Request, message Message) {
	Abort(r)
	Error(w, message)
}

// InternalServerError 500 系统内部错误
func InternalServerError(w http.ResponseWriter, r *http.Request) {

	registerFunc, exists := httpStatusMethodHub[http.StatusInternalServerError]

	if !exists {
		AbortError(w, r, MessageStatusInternalServerError)
		return
	}

	registerFunc(w, Abort(r), MessageStatusInternalServerError, http.StatusInternalServerError)
}

// Redirect 重定向
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, Abort(r), url, code)
}

// Forbidden 403 拒绝服务
func Forbidden(w http.ResponseWriter, r *http.Request) {

	registerFunc, exists := httpStatusMethodHub[http.StatusForbidden]

	if !exists {
		AbortError(w, r, MessageStatusForbidden)
		return
	}

	registerFunc(w, Abort(r), MessageStatusForbidden, http.StatusForbidden)
}

// NotFound 404 资源不存在
func NotFound(w http.ResponseWriter, r *http.Request) {

	registerFunc, exists := httpStatusMethodHub[http.StatusNotFound]

	if !exists {
		AbortError(w, r, MessageStatusNotFound)
		return
	}

	registerFunc(w, Abort(r), MessageStatusNotFound, http.StatusNotFound)
}

// Unauthorized 401 未认证
func Unauthorized(w http.ResponseWriter, r *http.Request) {

	registerFunc, exists := httpStatusMethodHub[http.StatusUnauthorized]

	if !exists {
		AbortError(w, r, MessageStatusUnauthorized)
		return
	}

	registerFunc(w, Abort(r), MessageStatusUnauthorized, http.StatusUnauthorized)
}

// BadRequest 400 请求异常
func BadRequest(w http.ResponseWriter, r *http.Request) {

	registerFunc, exists := httpStatusMethodHub[http.StatusBadRequest]

	if !exists {
		AbortError(w, r, MessageStatusBadRequest)
		return
	}

	registerFunc(w, Abort(r), MessageStatusBadRequest, http.StatusBadRequest)
}

// UnsupportedMediaType 415 不支持的媒体类型
func UnsupportedMediaType(w http.ResponseWriter, r *http.Request) {

	registerFunc, exists := httpStatusMethodHub[http.StatusUnsupportedMediaType]

	if !exists {
		AbortError(w, r, MessageUnsupportedMediaType)
		return
	}

	registerFunc(w, Abort(r), MessageUnsupportedMediaType, http.StatusUnsupportedMediaType)
}

// JSON 正常响应 JSON 请求
//
// 如果在序列化的过程中发生异常则响应服务器异常状态
func JSON(w http.ResponseWriter, r *http.Request, data interface{}, status int) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		InternalServerError(w, r)
		return
	}

	SetResponseHeader(w, "Content-Type", "application/json; charset=utf-8")
	_, err = SetHTTPRespStatus(w, status).Write(jsonBytes)
	if err != nil {
		InternalServerError(w, r)
	}
}

// DumpRequest 读取 r 的 body 内容并保持 r.Body 可持续使用
// 一般用于请求 handler 中读取 body 数据后，并保证后续代码可再次通过 r.Body 读取数据
func DumpRequest(r *http.Request) string {

	body, _ := ioutil.ReadAll(r.Body)

	// Reset resp.Body so it can be use again
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return string(body)
}

// DumpRequestPure 读取 r 的 body 内容并保持 r.Body 可持续使用
// 一般用于请求 handler 中读取 body 数据后，并保证后续代码可再次通过 r.Body 读取数据
func DumpRequestPure(r *http.Request) []byte {

	body, _ := ioutil.ReadAll(r.Body)
	// Reset resp.Body so it can be use again
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body
}

type StatusMethodCallback func(w http.ResponseWriter, r *http.Request, message Message, code int)

var (
	httpStatusMethodHubLock sync.Mutex
	httpStatusMethodHub     = map[int]StatusMethodCallback{}
)

// RegisterHttpStatusMethod 注册默认的请求状态回调
func RegisterHttpStatusMethod(code int, f StatusMethodCallback) {
	httpStatusMethodHubLock.Lock()
	httpStatusMethodHub[code] = f
	httpStatusMethodHubLock.Unlock()
}
