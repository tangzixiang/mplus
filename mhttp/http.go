package mhttp

import (
	"bytes"
	"encoding/json"
	"github.com/tangzixiang/mplus/context"
	"github.com/tangzixiang/mplus/header"
	"github.com/tangzixiang/mplus/message"
	"io/ioutil"
	"net/http"
)

const (
	requestAbortKey = "__abort"
	responseStatus  = "__status"
)

// use for http.Request.ParseMultipartForm, unit is bytes
var defaultMemory = int64(32 * 1024 * 1024)

// EmptyRespData 空响应体
var EmptyRespData = map[string]interface{}{}

// SetDefaultMemorySize 设置 http.Request.ParseMultipartForm 参数 maxMemory
func SetDefaultMemorySize(size int64) {
	defaultMemory = size
}

// DefaultMemorySize 获取 http.Request.ParseMultipartForm 参数 maxMemory
func DefaultMemorySize() int64 {
	return defaultMemory
}

// Abort 标识当前请求链已中断
func Abort(r *http.Request) *http.Request {
	context.SetContextValue(r.Context(), requestAbortKey, true)
	return r
}

// NotAbort 取消当前请求标识中断状态
func NotAbort(r *http.Request) *http.Request {
	context.SetContextValue(r.Context(), requestAbortKey, false)
	return r
}

// IsAbort 判断当前请求链是否已中断
func IsAbort(r *http.Request) bool {
	return context.GetContextValueBool(r.Context(), requestAbortKey)
}

// Error 返回异常信息，当前方法触发的请求响应内容将是文本格式
func Error(w http.ResponseWriter, m message.Message) {
	http.Error(SetHTTPRespStatus(w, m.Status(), false), m.Default(), m.Status())
}

// ErrorEmpty 返回异常信息，当前方法触发的请求无响应内容
func ErrorEmpty(w http.ResponseWriter, m message.Message) {
	http.Error(SetHTTPRespStatus(w, m.Status(), false), "", m.Status())
}

// Plain 返回一个 text/plain 格式的响应
func Plain(w http.ResponseWriter, m message.Message) {

	header.SetResponseHeaderIf(
		header.GetResponseHeader(w, header.ContentType) == "", w, header.ContentType, "text/plain; charset=utf-8")

	header.SetResponseHeaderIf(
		header.GetResponseHeader(w, header.ContentTypeOptions) == "", w, header.ContentTypeOptions, "nosniff")

	SetHTTPRespStatus(w, m.Status(), true).Write([]byte(m.Default()))
}

// PlainEmpty 返回一个 text/plain 格式的响应，当前方法触发的请求无响应内容
func PlainEmpty(w http.ResponseWriter, m message.Message) {

	header.SetResponseHeaderIf(
		header.GetResponseHeader(w, header.ContentType) == "", w, header.ContentType, "text/plain; charset=utf-8")

	header.SetResponseHeaderIf(
		header.GetResponseHeader(w, header.ContentTypeOptions) == "", w, header.ContentTypeOptions, "nosniff")

	SetHTTPRespStatus(w, m.Status(), true)
}

// AbortError 终止请求链并返回异常信息，当前方法触发的请求响应内容将是文本格式
func AbortError(w http.ResponseWriter, r *http.Request, m message.Message) {
	Abort(r)
	Error(w, m)
}

// AbortEmptyError 终止请求链并返回异常信息，当前方法触发的请求响应内容将是文本格式
func AbortEmptyError(w http.ResponseWriter, r *http.Request, m message.Message) {
	Abort(r)
	ErrorEmpty(w, m)
}

// AbortEmptyPlain 终止请求链，当前方法触发的请求响应内容将是文本格式
func AbortEmptyPlain(w http.ResponseWriter, r *http.Request, m message.Message) {
	Abort(r)
	PlainEmpty(w, m)
}

// AbortPlain 终止请求链，当前方法触发的请求响应内容将是文本格式
func AbortPlain(w http.ResponseWriter, r *http.Request, m message.Message) {
	Abort(r)
	Plain(w, m)
}

// JSONOK 以 JSON 格式输出请求状态码为 200 的响应
func JSONOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	JSON(w, r, data, http.StatusOK)
}

// JSON 正常响应 JSON 请求
//
// 如果在序列化的过程中发生异常则响应服务器异常状态
func JSON(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	Abort(r)

	if data == nil {
		data = EmptyRespData
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		InternalServerError(w, r)
		return
	}

	header.SetResponseHeader(w, "Content-Type", "application/json; charset=utf-8")
	_, err = SetHTTPRespStatus(w, status).Write(jsonBytes)
	if err != nil {
		InternalServerError(w, r)
	}
}

// Redirect 重定向
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func Redirect(w http.ResponseWriter, r *http.Request, url string, statusCode int) {
	http.Redirect(SetHTTPRespStatus(w, statusCode, false), Abort(r), url, statusCode)
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
