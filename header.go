package mplus

import (
	"net"
	"net/http"
	"strings"
)

// 请求头常量
const (
	RequestIDHeader           = "X-Request-Id" // 当前请求的 request-id
	ContentTypeHeader         = "Content-Type"
	ForwardedForHeader        = "X-Forwarded-For"
	RealIPHeader              = "X-Real-Ip"
	AppEngineRemoteAddrHeader = "X-Appengine-Remote-Addr"
)

// 请求头分割字符
const HeaderSplitSep = " "

// GetHeader 获取指定请求头
func GetHeader(r *http.Request, header string) string {
	return r.Header.Get(header)
}

// SplitHeader 分割指定请求头内容
func SplitHeader(r *http.Request, header string) []string {
	return strings.Split(GetHeader(r, header), HeaderSplitSep)
}

// GetHeaderRequestID 获取 RequestID
func GetHeaderRequestID(r *http.Request) string {
	return GetHeader(r, RequestIDHeader)
}

// SetRequestHeader 添加指定请求头
func SetRequestHeader(r *http.Request, key, value string) *http.Request {
	r.Header.Set(key, value)
	return r
}

// SetRequestHeaders 添加指定请求头
func SetRequestHeaders(r *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		r.Header.Set(key, value)
	}
	return r
}

// GetResponseHeader 获取指定请求头
func GetResponseHeader(w http.ResponseWriter, header string) string {
	return w.Header().Get(header)
}

// SetResponseHeader 将指定请求头添加到响应中
func SetResponseHeader(w http.ResponseWriter, key, value string) http.ResponseWriter {
	w.Header().Set(key, value)
	return w
}

// SetResponseHeaders 将指定请求头添加到响应中
func SetResponseHeaders(w http.ResponseWriter, headers map[string]string) http.ResponseWriter {
	for key, value := range headers {
		w.Header().Set(key, value)
	}
	return w
}

// SetRequestHeaderRequestID 在 request 中添加 RequestID
func SetRequestHeaderRequestID(r *http.Request, id string) *http.Request {
	SetRequestHeader(r, RequestIDHeader, id)
	return r
}

// GetClientIP 获取客户端 IP 地址
// ClientIP implements a best effort algorithm to return the real client IP, it parses
// X-Real-IP and X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
// Use X-Forwarded-For before X-Real-Ip as nginx uses X-Real-Ip with the proxy's IP.
// from gin
func GetClientIP(r *http.Request) string {
	clientIP := GetHeader(r, ForwardedForHeader)
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])

	if clientIP == "" {
		clientIP = strings.TrimSpace(GetHeader(r, RealIPHeader))
	}
	if clientIP != "" {
		return clientIP
	}

	// #726 #755 If enabled, it will thrust some headers starting with
	// 'X-AppEngine...' for better integration with that PaaS.
	if addr := GetHeader(r, AppEngineRemoteAddrHeader); addr != "" {
		return addr
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

const (
	ContentTypeJSON   = "application/json"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeText   = "text/plain"
	ContentTypeXML    = "application/xml"
	ContentTypeStream = "application/octet-stream"
)
