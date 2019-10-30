package mplus

import (
	"github.com/tangzixiang/mplus/mime"
)

// mime 媒体类型
const (
	MIMEJSON              = mime.MIMEJSON
	MIMEHTML              = mime.MIMEHTML
	MIMEXML               = mime.MIMEXML
	MIMEXML2              = mime.MIMEXML2
	MIMEPlain             = mime.MIMEPlain
	MIMEPOSTForm          = mime.MIMEPOSTForm
	MIMEMultipartPOSTForm = mime.MIMEMultipartPOSTForm
	MIMEPROTOBUF          = mime.MIMEPROTOBUF
	MIMEMSGPACK           = mime.MIMEMSGPACK
	MIMEMSGPACK2          = mime.MIMEMSGPACK2
	MIMEStream            = mime.MIMEStream
)

var ParseMediaType = mime.ParseMediaType
var ParseRequestMediaType = mime.ParseRequestMediaType
var ParseResponseMediaType = mime.ParseResponseMediaType
