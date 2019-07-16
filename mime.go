package mplus

import (
	"mime"
	"net/http"
)

// mime 媒体类型
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEStream            = "application/octet-stream"
)

// ContentType 请求体 ContentType 标识
const ContentType = "Content-Type"

// ParseMediaType 解析 Header 中的 Content-Type
func ParseMediaType(r *http.Request) (string, error) {
	var err error

	ct := GetHeader(r, ContentType)

	// RFC 7231, section 3.1.1.5 - empty type
	//   MAY be treated as application/octet-stream
	// if ct == "" {
	// 	ct = MIMEStream
	// }

	// 允许 post 请求不带请求体
	if ct == "" {
		ct = MIMEPOSTForm
	}

	ct, _, err = mime.ParseMediaType(ct)
	return ct, err
}
