package header

import (
	"net"
	"net/http"
	"strings"
)

// GetHeader 获取指定请求头
func GetHeader(r *http.Request, header string) string {
	return r.Header.Get(header)
}

// GetHeaderValues 获取指定 key 的请求头集合
func GetHeaderValues(r *http.Request, header string) []string {
	return r.Header[header]
}

// SplitHeader 分割指定请求头内容
func SplitHeader(r *http.Request, header string, splitSep string) []string {
	return strings.Split(GetHeader(r, header), splitSep)
}

// GetHeaderRequestID 获取 RequestID
func GetHeaderRequestID(r *http.Request) string {
	return GetHeader(r, RequestID)
}

// SetRequestHeader 添加指定请求头
func SetRequestHeader(r *http.Request, key, value string) *http.Request {
	r.Header.Set(key, value)
	return r
}

// SetRequestHeaderIf 添加指定请求头，仅当 ensure 为 true 时生效
func SetRequestHeaderIf(ensure bool, r *http.Request, key, value string) *http.Request {
	if ensure {
		r.Header.Set(key, value)
	}
	return r
}

// SetRequestHeaders 添加指定请求头
func SetRequestHeaders(r *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		r.Header.Set(key, value)
	}
	return r
}

// SetRequestHeadersIf 添加指定请求头，仅当 ensure 为 true 时生效
func SetRequestHeadersIf(ensure bool, r *http.Request, headers map[string]string) *http.Request {
	if ensure {
		for key, value := range headers {
			r.Header.Set(key, value)
		}
	}
	return r
}

// AddRequestHeader 添加指定请求头到 r
func AddRequestHeader(r *http.Request, key string, values ...string) *http.Request {
	for _, value := range values {
		r.Header.Add(key, value)
	}
	return r
}

// AddRequestHeaderIf 添加指定请求头到 r，仅当 ensure 为 true 时生效
func AddRequestHeaderIf(ensure bool, r *http.Request, key string, values ...string) *http.Request {
	if ensure {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	return r
}

// AddRequestHeaders 添加指定请求头到 r
func AddRequestHeaders(r *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		r.Header.Add(key, value)
	}
	return r
}

// AddRequestHeadersIf 添加指定请求头到 r，仅当 ensure 为 true 时生效
func AddRequestHeadersIf(ensure bool, r *http.Request, headers map[string]string) *http.Request {
	if ensure {
		for key, value := range headers {
			r.Header.Add(key, value)
		}
	}
	return r
}

// GetResponseHeader 获取指定响应头
func GetResponseHeader(w http.ResponseWriter, header string) string {
	return w.Header().Get(header)
}

// GetResponseHeaderValues 获取指定 key 响应头集合
func GetResponseHeaderValues(w http.ResponseWriter, header string) []string {
	return w.Header()[header]
}

// SetResponseHeader 将指定请求头添加到响应中
func SetResponseHeader(w http.ResponseWriter, key, value string) http.ResponseWriter {
	w.Header().Set(key, value)
	return w
}

// SetResponseHeaderIf 将指定请求头添加到响应中，仅当 ensure 为 true 时生效
func SetResponseHeaderIf(ensure bool, w http.ResponseWriter, key, value string) http.ResponseWriter {
	if ensure {
		w.Header().Set(key, value)
	}
	return w
}

// SetResponseHeaders 将指定请求头添加到响应中
func SetResponseHeaders(w http.ResponseWriter, headers map[string]string) http.ResponseWriter {
	for key, value := range headers {
		w.Header().Set(key, value)
	}
	return w
}

// SetResponseHeadersIf 将指定请求头添加到响应中，仅当 ensure 为 true 时生效
func SetResponseHeadersIf(ensure bool, w http.ResponseWriter, headers map[string]string) http.ResponseWriter {
	if ensure {
		for key, value := range headers {
			w.Header().Set(key, value)
		}
	}
	return w
}

// AddResponseHeader 将指定响应头添加到 w 中
func AddResponseHeader(w http.ResponseWriter, key string, values ...string) http.ResponseWriter {
	for _, value := range values {
		w.Header().Add(key, value)
	}
	return w
}

// AddResponseHeaderIf 将指定响应头添加到 w 中，仅当 ensure 为 true 时生效
func AddResponseHeaderIf(ensure bool, w http.ResponseWriter, key string, values ...string) http.ResponseWriter {
	if ensure {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	return w
}

// AddResponseHeaders 将指定响应头添加到 w 中
func AddResponseHeaders(w http.ResponseWriter, headers map[string]string) http.ResponseWriter {
	for key, value := range headers {
		w.Header().Add(key, value)
	}
	return w
}

// AddResponseHeadersIf 将指定响应头添加到 w 中，仅当 ensure 为 true 时生效
func AddResponseHeadersIf(ensure bool, w http.ResponseWriter, headers map[string]string) http.ResponseWriter {
	if ensure {
		for key, value := range headers {
			w.Header().Add(key, value)
		}
	}
	return w
}

// SetRequestHeaderRequestID 在 request 中添加 RequestID
func SetRequestHeaderRequestID(r *http.Request, id string) *http.Request {
	SetRequestHeader(r, RequestID, id)
	return r
}

// SetResponseHeaderRequestID 在 response 中添加 RequestID
func SetResponseHeaderRequestID(w http.ResponseWriter, id string) http.ResponseWriter {
	SetResponseHeader(w, RequestID, id)
	return w
}

// GetClientIP 获取客户端 IP 地址
// ClientIP implements a best effort algorithm to return the real client IP, it parses
// X-Real-IP and X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
// Use X-Forwarded-For before X-Real-Ip as nginx uses X-Real-Ip with the proxy's IP.
func GetClientIP(r *http.Request) string {
	clientIP := GetHeader(r, ForwardedFor)
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])

	if clientIP == "" {
		clientIP = strings.TrimSpace(GetHeader(r, RealIP))
	}

	if clientIP != "" {
		return clientIP
	}

	// #726 #755 If enabled, it will thrust some headers starting with
	// 'X-AppEngine...' for better integration with that PaaS.
	if addr := GetHeader(r, AppEngineRemoteAddr); addr != "" {
		return addr
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
