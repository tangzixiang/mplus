package mplus

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

// PP 缓存了当前请求的 w ResponseWriter 及 r Request
type PP struct {
	w http.ResponseWriter
	r *http.Request

	h http.Handler
}

// PlusPlus 返回一个当前请求上下文托管对象
func PlusPlus(w http.ResponseWriter, r *http.Request) *PP {
	return &PP{w, r, nil}
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
	return GetHeaderRequestID(p.r)
}

// VO 获取请求对象
func (p *PP) VO() interface{} {
	return GetContextValue(p.r.Context(), ReqData)
}

// Abort 将当前请求标识为中断
func (p *PP) Abort() *PP {
	Abort(p.r)
	return p
}

// IsAbort 判断当前请求是否中断
func (p *PP) IsAbort() bool {
	return IsAbort(p.r)
}

// ErrorMsg 响应异常信息
func (p *PP) ErrorMsg(message Message) *PP {
	Error(p.w, message)
	return p
}

// AbortErrorMsg 将当前请求标识为中断并响应异常信息
func (p *PP) AbortErrorMsg(message Message) *PP {
	AbortError(p.w, p.r, message)
	return p
}

// Redirect 3xx 重定向
func (p *PP) Redirect(url string, code int) *PP {
	Redirect(p.w, p.r, url, code)
	return p
}

// BadRequest 400 请求失败
func (p *PP) BadRequest() *PP {
	BadRequest(p.w, p.r)
	return p
}

// Forbidden 403 请求拒绝
func (p *PP) Forbidden() *PP {
	Forbidden(p.w, p.r)
	return p
}

// NotFound 404 请求资源不存在
func (p *PP) NotFound() *PP {
	NotFound(p.w, p.r)
	return p
}

// Unauthorized 401 未认证请求
func (p *PP) Unauthorized() *PP {
	Unauthorized(p.w, p.r)
	return p
}

// UnsupportedMediaType 415 不支持的媒体类型
func (p *PP) UnsupportedMediaType() *PP {
	UnsupportedMediaType(p.w, p.r)
	return p
}

// JSON 响应指定状态码的 JSON 数据
func (p *PP) JSON(data interface{}, status int) *PP {
	JSON(p.w, p.r, data, status)
	return p
}

// InternalServerError 500 内部服务器异常
func (p *PP) InternalServerError() *PP {
	InternalServerError(p.w, p.r)
	return p
}

// Get 获取上下文内容
func (p *PP) Get(key string) interface{} {
	return GetContextValue(p.r.Context(), key)
}

// Set 设置上下文内容
func (p *PP) Set(key string, value interface{}) *PP {
	SetContextValue(p.r.Context(), key, value)
	return p
}

// SetR 设置上下文内容,并返回原内容
func (p *PP) SetR(key string, value interface{}) interface{} {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetStringR 设置上下文内容,并返回原内容
func (p *PP) SetStringR(key string, value string) string {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetIntR 设置上下文内容,并返回原内容
func (p *PP) SetIntR(key string, value int) int {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt8R 设置上下文内容,并返回原内容
func (p *PP) SetInt8R(key string, value int8) int8 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt16R 设置上下文内容,并返回原内容
func (p *PP) SetInt16R(key string, value int16) int16 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt32R 设置上下文内容,并返回原内容
func (p *PP) SetInt32R(key string, value int32) int32 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetInt64R 设置上下文内容,并返回原内容
func (p *PP) SetInt64R(key string, value int64) int64 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUintR 设置上下文内容,并返回原内容
func (p *PP) SetUintR(key string, value uint) uint {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint8R 设置上下文内容,并返回原内容
func (p *PP) SetUint8R(key string, value uint8) uint8 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint16R 设置上下文内容,并返回原内容
func (p *PP) SetUint16R(key string, value uint16) uint16 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint32R 设置上下文内容,并返回原内容
func (p *PP) SetUint32R(key string, value uint32) uint32 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetUint64R 设置上下文内容,并返回原内容
func (p *PP) SetUint64R(key string, value uint64) uint64 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetBoolR 设置上下文内容,并返回原内容
func (p *PP) SetBoolR(key string, value bool) bool {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetByteR 设置上下文内容,并返回原内容
func (p *PP) SetByteR(key string, value byte) byte {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetBytesR 设置上下文内容,并返回原内容
func (p *PP) SetBytesR(key string, value []byte) []byte {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetTimeR 设置上下文内容,并返回原内容
func (p *PP) SetTimeR(key string, value time.Time) time.Time {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetFloat32R 设置上下文内容,并返回原内容
func (p *PP) SetFloat32R(key string, value float32) float32 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// SetFloat64R 设置上下文内容,并返回原内容
func (p *PP) SetFloat64R(key string, value float64) float64 {
	SetContextValue(p.r.Context(), key, value)
	return value
}

// GetString 获取 string 类型的上下文内容，获取失败返回其零值
func (p *PP) GetString(key string) string {
	return GetContextValueString(p.r.Context(), key)
}

// GetInt 获取 int 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt(key string) int {
	return GetContextValueInt(p.r.Context(), key)
}

// GetInt8 获取 int8 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt8(key string) int8 {
	return GetContextValueInt8(p.r.Context(), key)
}

// GetInt16 获取 int16 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt16(key string) int16 {
	return GetContextValueInt16(p.r.Context(), key)
}

// GetInt32 获取 int32 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt32(key string) int32 {
	return GetContextValueInt32(p.r.Context(), key)
}

// GetInt64 获取 int64 类型的上下文内容，获取失败返回其零值
func (p *PP) GetInt64(key string) int64 {
	return GetContextValueInt64(p.r.Context(), key)
}

// GetUInt 获取 uint 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt(key string) uint {
	return GetContextValueUInt(p.r.Context(), key)
}

// GetUInt8 获取 uint8 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt8(key string) uint8 {
	return GetContextValueUInt8(p.r.Context(), key)
}

// GetUInt16 获取 uint16 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt16(key string) uint16 {
	return GetContextValueUInt16(p.r.Context(), key)
}

// GetUInt32 获取 uint32 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt32(key string) uint32 {
	return GetContextValueUInt32(p.r.Context(), key)
}

// GetUInt64 获取 uint64 类型的上下文内容，获取失败返回其零值
func (p *PP) GetUInt64(key string) uint64 {
	return GetContextValueUInt64(p.r.Context(), key)
}

// GetBool 获取 bool 类型的上下文内容，获取失败返回其零值
func (p *PP) GetBool(key string) bool {
	return GetContextValueBool(p.r.Context(), key)
}

// GetByte 获取 byte 类型的上下文内容，获取失败返回其零值
func (p *PP) GetByte(key string) byte {
	return GetContextValueByte(p.r.Context(), key)
}

// GetBytes 获取 []byte 类型的上下文内容，获取失败返回其零值
func (p *PP) GetBytes(key string) []byte {
	return GetContextValueBytes(p.r.Context(), key)
}

// GetTime 获取 time.Time 类型的上下文内容，获取失败返回其零值
func (p *PP) GetTime(key string) time.Time {
	return GetContextValueTime(p.r.Context(), key)
}

// GetUInt32 获取 float32 类型的上下文内容，获取失败返回其零值
func (p *PP) GetFloat32(key string) float32 {
	return GetContextValueFloat32(p.r.Context(), key)
}

// GetFloat64 获取 float64 类型的上下文内容，获取失败返回其零值
func (p *PP) GetFloat64(key string) float64 {
	return GetContextValueFloat64(p.r.Context(), key)
}

// GetStringDf 获取 string 类型的上下文内容，获取失败返回默认值
func (p *PP) GetStringDf(key string, defaultValue string) string {
	return GetContextValueString(p.r.Context(), key, defaultValue)
}

// GetIntDf 获取 int 类型的上下文内容，获取失败返回默认值
func (p *PP) GetIntDf(key string, defaultValue int) int {
	return GetContextValueInt(p.r.Context(), key, defaultValue)
}

// GetInt8Df 获取 int8 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt8Df(key string, defaultValue int8) int8 {
	return GetContextValueInt8(p.r.Context(), key, defaultValue)
}

// GetInt16Df 获取 int16 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt16Df(key string, defaultValue int16) int16 {
	return GetContextValueInt16(p.r.Context(), key, defaultValue)
}

// GetInt32Df 获取 int32 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt32Df(key string, defaultValue int32) int32 {
	return GetContextValueInt32(p.r.Context(), key, defaultValue)
}

// GetInt64Df 获取 int64 类型的上下文内容，获取失败返回默认值
func (p *PP) GetInt64Df(key string, defaultValue int64) int64 {
	return GetContextValueInt64(p.r.Context(), key, defaultValue)
}

// GetUIntDf 获取 uint 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUIntDf(key string, defaultValue uint) uint {
	return GetContextValueUInt(p.r.Context(), key, defaultValue)
}

// GetUInt8Df 获取 uint8 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt8Df(key string, defaultValue uint8) uint8 {
	return GetContextValueUInt8(p.r.Context(), key, defaultValue)
}

// GetUInt16Df 获取 uint16 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt16Df(key string, defaultValue uint16) uint16 {
	return GetContextValueUInt16(p.r.Context(), key, defaultValue)
}

// GetUInt32Df 获取 uint32 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt32Df(key string, defaultValue uint32) uint32 {
	return GetContextValueUInt32(p.r.Context(), key, defaultValue)
}

// GetUInt64Df 获取 uint64 类型的上下文内容，获取失败返回默认值
func (p *PP) GetUInt64Df(key string, defaultValue uint64) uint64 {
	return GetContextValueUInt64(p.r.Context(), key, defaultValue)
}

// GetBoolDf 获取 bool 类型的上下文内容，获取失败返回默认值
func (p *PP) GetBoolDf(key string, defaultValue bool) bool {
	return GetContextValueBool(p.r.Context(), key, defaultValue)
}

// GetByteDf 获取 byte 类型的上下文内容，获取失败返回默认值
func (p *PP) GetByteDf(key string, defaultValue byte) byte {
	return GetContextValueByte(p.r.Context(), key, defaultValue)
}

// GetBytesDf 获取 []byte 类型的上下文内容，获取失败返回默认值
func (p *PP) GetBytesDf(key string, defaultValue []byte) []byte {
	return GetContextValueBytes(p.r.Context(), key, defaultValue)
}

// GetTimeDf 获取 time.Time 类型的上下文内容，获取失败返回默认值
func (p *PP) GetTimeDf(key string, defaultValue time.Time) time.Time {
	return GetContextValueTime(p.r.Context(), key, defaultValue)
}

// GetUInt32 获取 float32 类型的上下文内容，获取失败返回默认值
func (p *PP) GetFloat32Df(key string, defaultValue float32) float32 {
	return GetContextValueFloat32(p.r.Context(), key, defaultValue)
}

// GetFloat64 获取 float64 类型的上下文内容，获取失败返回默认值
func (p *PP) GetFloat64Df(key string, defaultValue float64) float64 {
	return GetContextValueFloat64(p.r.Context(), key, defaultValue)
}

// DoCallback 查询指定错误嘛注册的回调并执行
func (p *PP) CallbackByCode(errorCode int, respData interface{}) *PP {
	mByCode := Messages.Get(errorCode)

	if m, ok := mByCode.(*message); ok && m != nil {
		m.callback(p, mByCode, respData)
	}

	return p
}

// Queries 获取 URL 上的请求字段
func (p *PP) Queries() url.Values {
	return Queries(p.r)
}

// GetQuery 获取 URL 上的指定的请求字段
func (p *PP) GetQuery(key string) (value string, exists bool) {
	values, exists := Queries(p.r)[key]
	if !exists || len(values) <= 0 {
		return "", false
	}

	return values[0], true
}

// GetQuery 获取 URL 上的指定的请求字段
func (p *PP) Query(key string) (value string) {
	values, exists := Queries(p.r)[key]
	if !exists || len(values) <= 0 {
		return ""
	}

	return values[0]
}

// GetQuery 获取 URL 上的指定的请求字段，获取失败返回默认值
func (p *PP) GetQueryDf(key, defaultValue string) (value string) {
	values, exists := Queries(p.r)[key]
	if !exists || len(values) <= 0 {
		return defaultValue
	}

	return values[0]
}

// GetHeader 获取指定请求头信息
func (p *PP) GetHeader(key string) string {
	return GetHeader(p.r, key)
}

// GetHeaderDf 获取指定请求头信息，获取失败返回默认值
func (p *PP) GetHeaderDf(key string, defaultValue string) string {
	value := GetHeader(p.r, key)
	if value == "" {
		return defaultValue
	}
	return value
}

// SplitHeader 分割指定请求头内容
func (p *PP) SplitHeader(key string) [] string {
	return SplitHeader(p.r, key)
}

// WriteRespHeader 将指定头信息写入到响应中
func (p *PP) WriteRespHeader(key, value string) *PP {
	SetResponseHeader(p.w, key, value)
	return p
}

// WriteRespHeaders 将指定头信息写入到响应中
func (p *PP) WriteRespHeaders(headers map[string]string) *PP {
	SetResponseHeaders(p.w, headers)
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
// handlerFunc.ServeHTTP(w,r)
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
func (p *PP) Status(code int) *PP {
	SetHTTPRespStatus(p.w, code)
	return p
}

// GetStatus 获取响应状态码
func (p *PP) GetStatus() int {
	return GetHTTPRespStatus(p.w)
}

// CopyReq 拷贝一个请求，body 不会被拷贝，因为 body 是一个数据流
func (p *PP) CopyReq() *http.Request {
	return CopyRequest(p.r)
}

// Method 获取请求方式
func (p *PP) Method() string {
	return p.r.Method
}

// GetClientIP 获取客户端 IP 地址
func (p *PP) GetClientIP() string {
	return GetClientIP(p.r)
}

// ReqBody 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用
func (p *PP) ReqBody() string {
	return DumpRequest(p.r)
}

// ReqBody 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用
func (p *PP) ReqBodyPure() []byte {
	return DumpRequestPure(p.r)
}

// ReqBodyMap 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用,body 内容会被序列化成 map[string] interface{}
func (p *PP) ReqBodyMap() (map[string]interface{}, error) {
	body := DumpRequestPure(p.r)
	m := map[string]interface{}{}

	return m, json.Unmarshal(body, &m)
}

// ReqBodyMap 读取 p.r 的 body 内容并保持 p.r.Body 可持续使用,body 内容会被序列化至 unmarshaler
func (p *PP) ReqBodyToUnmarshaler(unmarshaler json.Unmarshaler) (error) {
	return unmarshaler.UnmarshalJSON(DumpRequestPure(p.r))
}
