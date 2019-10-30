package mime

import (
	"mime"
	"net/http"

	"github.com/tangzixiang/mplus/header"
)

// mime 媒体类型
const (
	MIMEJSON              = header.ContentTypeJSON
	MIMEHTML              = header.ContentTypeHTML
	MIMEXML               = header.ContentTypeXML
	MIMEXML2              = header.ContentTypeXML2
	MIMEPlain             = header.ContentTypePlain
	MIMEPOSTForm          = header.ContentTypeForm
	MIMEMultipartPOSTForm = header.ContentTypeMultipartPOSTForm
	MIMEPROTOBUF          = header.ContentTypePROTOBUF
	MIMEMSGPACK           = header.ContentTypeMSGPACK
	MIMEMSGPACK2          = header.ContentTypeMSGPACK2
	MIMEStream            = header.ContentTypeStream
)

// ParseMediaType 解析 Header 中的 Content-Type
func ParseMediaType(ct string) (string, error) {
	var err error

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

// ParseRequestMediaType 解析请求 Header 中的 Content-Type
func ParseRequestMediaType(r *http.Request) (string, error) {
	return ParseMediaType(header.GetHeader(r, header.ContentType))
}

// ParseResponseMediaType 解析响应 Header 中的 Content-Type
func ParseResponseMediaType(w http.ResponseWriter) (string, error) {
	return ParseMediaType(header.GetResponseHeader(w, header.ContentType))
}
