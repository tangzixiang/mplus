package mplus

import (
	"net/http"

	"github.com/pkg/errors"
)

// ValidateErrorType 异常类型
type ValidateErrorType uint8

const (
	// ErrBodyRead 请求体读取失败
	ErrBodyRead ValidateErrorType = iota
	// ErrBodyUnmarshal 请求体序列化失败
	ErrBodyUnmarshal
	// ErrBodyParse 请求体解析失败
	ErrBodyParse
	// ErrMediaTypeParse 媒体类型解析失败
	ErrMediaTypeParse
	// ErrMediaType 不支持的媒体类型
	ErrMediaType
	// ErrDecode 请求参数解析失败
	ErrDecode
	// ErrParseQuery 请求参数获取失败
	ErrParseQuery
	// ErrBodyValidate 请求体内容校验失败
	ErrBodyValidate
	// ErrRequestValidate 自定义请求体内容校验失败
	ErrRequestValidate
	// ErrDefault 默认
	ErrDefault
)

// ValidateErrorTypeMsg 异常与描述信息
var ValidateErrorTypeMsg = map[ValidateErrorType]string{
	ErrBodyRead:        "read request body failed",
	ErrBodyUnmarshal:   "unmarshal request body failed",
	ErrBodyParse:       "parse request body failed",
	ErrMediaType:       "parse request media type failed",
	ErrMediaTypeParse:  "parse request media type failed",
	ErrDecode:          "decode request query failed",
	ErrParseQuery:      "parse request query failed",
	ErrBodyValidate:    "validate request body failed",
	ErrRequestValidate: "validate request body failed",
	ErrDefault:         "validate request failed",
}

// ValidateError 校验异常
type ValidateError struct {
	errType ValidateErrorType
	lastErr error
}

func (ve ValidateError) LastErr() error {
	return ve.lastErr
}

// Error 等效于  ve.LastErr().Error()
func (ve ValidateError) Error() string {
	return ve.lastErr.Error()
}

// IsErr 判断是否指定类型异常
func (ve ValidateError) IsErr(errType ValidateErrorType) bool {
	return errType == ve.errType
}

// Type 查看异常类型
func (ve ValidateError) Type() ValidateErrorType {
	return ve.errType
}

// ValidateErrorWrap 包装一个解析异常
func ValidateErrorWrap(err error, errType ValidateErrorType) error {
	return errors.Wrap(ValidateError{errType: errType, lastErr: err}, ValidateErrorTypeMsg[errType])
}

// ValidateErrorFunc 请求解析失败的处理器
type ValidateErrorFunc func(w http.ResponseWriter, r *http.Request, err error)

// 存在于 validateErrorHub 的处理器将会覆盖 defaultValidateErrorHub 的处理器
var validateErrorHub = map[ValidateErrorType]ValidateErrorFunc{}

// 默认解析错误处理器
var defaultValidateErrorHub = map[ValidateErrorType]http.HandlerFunc{
	ErrBodyRead:        InternalServerError,
	ErrBodyUnmarshal:   InternalServerError,
	ErrMediaType:       UnsupportedMediaType,
	ErrMediaTypeParse:  UnsupportedMediaType,
	ErrBodyParse:       BadRequest,
	ErrDecode:          BadRequest,
	ErrParseQuery:      BadRequest,
	ErrBodyValidate:    BadRequest,
	ErrRequestValidate: BadRequest,
	ErrDefault:         BadRequest,
}

// RegisterValidateErrorFunc 注册一个解析失败的处理器
func RegisterValidateErrorFunc(errType ValidateErrorType, fun ValidateErrorFunc) {
	validateErrorHub[errType] = fun
}
