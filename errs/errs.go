package errs

import (
	"github.com/tangzixiang/mplus/message"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/tangzixiang/mplus/mhttp"
)

// ValidateErrorType 请求数据读取及校验异常类型
// 该类型的异常可以通过 RegisterValidateErrorFunc 注册 hook 进行捕获处理
type ValidateErrorType uint8

const (
	// ErrBodyRead 请求体读取失败
	// 出现于 mplus.Bind()，
	// 若当前请求为 POST/PUT/PATCH/DELETE 请求且格式为 json 时，读取 r.body 失败时触发
	// 默认不强校验 body 数据为空的情况，可以通过 mplus.SetStrictJSONBodyCheck 更改为强校验模式
	ErrBodyRead ValidateErrorType = iota
	// ErrBodyUnmarshal 请求体序列化失败
	// 出现于 mplus.Bind()，
	// 若当前请求为 POST/PUT/PATCH/DELETE 请求且格式为 json 时，反序列化数据至 model 失败时触发
	ErrBodyUnmarshal
	// ErrBodyParse 请求体解析失败
	// 出现于 mplus.Bind()，
	// 若当前请求为 POST/PUT/PATCH/DELETE 请求且格式为 x-www-form-urlencoded/form-data 时，执行 r.ParseForm 或 r.ParseMultipartForm 失败时触发
	ErrBodyParse
	// ErrMediaTypeParse 媒体类型解析失败
	// 出现于 mplus.Bind()，若请求提供的 MediaType 无法正常解析时触发
	ErrMediaTypeParse
	// ErrMediaType 不支持的媒体类型
	// 出现于 mplus.Bind()，
	// 若当前请求为 POST/PUT/PATCH/DELETE 请求且格式不为  x-www-form-urlencoded/form-data/json 时触发
	ErrMediaType
	// ErrDecode 请求参数解析失败
	// 出现于 mplus.Bind()，
	// 若当前请求为 POST/PUT/PATCH/DELETE 请求且格式为 x-www-form-urlencoded/form-data 时或
	// 若当前请求为 GET HEAD OPTION 时，
	// 其请求参数无法写入至 model 对象时触发
	ErrDecode
	// ErrParseQuery 请求参数获取失败
	// 出现于 mplus.Bind() ，
	// 当 query string 存在于 URL 上，且无法通过 r.PostForm 及 r.Form 获取时，解析 URL 失败触发，
	// r.PostForm 及 r.Form 仅在 POST/PUT/PATCH/DELETE 请求且格式为 x-www-form-urlencoded/form-data 时解析
	ErrParseQuery
	// ErrBodyValidate 请求体内容校验失败
	// 出现于 mplus.Bind() ，若 validator 校验 model 对象时失败时触发
	ErrBodyValidate
	// ErrRequestValidate 自定义请求体内容校验失败
	// 出现于 mplus.Bind() ，
	// 若 model 对象实现了 mplus.RequestValidate 接口，当该对象的 Validate 函数执行返回异常时触发
	ErrRequestValidate
	// ErrDefault 默认异常，
	// 出现于 mplus.Bind() ，当内部异常类型捕获失败时返回该异常
	ErrDefault
	// ErrModelSelect 请求数据绑定 model 选择失败，用户主动抛出异常
	// 出现于 mplus.Bind() ，
	// 当传递的参数为函数类型，且函数执行返回异常时触发，函数型参数仅在请求进入路由链路时执行
	ErrModelSelect
	// ErrModelSelectType 请求数据绑定 model 选择失败，回调函数返回的类型必须是指针类型
	// 出现于 mplus.Bind() ，
	// 当传递的参数为函数类型，且函数执行返回异常时触发，函数型参数仅在请求进入路由链路时执行
	ErrModelSelectType
)

// ValidateErrorTypeMsg ValidateErrorType 异常与描述信息
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
	ErrModelSelect:     "select request model failed",
	ErrModelSelectType: "select request model type error,must be ptr",
}

// ValidateError 校验异常
type ValidateError struct {
	errType ValidateErrorType
	lastErr error
}

func (ve ValidateError) String() string {
	return `err: ` + ve.Error() + `; type: ` + strconv.Itoa(int(uint8(ve.Type())))
}

func (ve ValidateError) LastErr() error {
	return ve.lastErr
}

// Error 等效于 ve.LastErr().Error()
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

// NewValidateError 获取一个校验异常实例
func NewValidateError(errType ValidateErrorType, lastErr error) ValidateError {
	return ValidateError{errType: errType, lastErr: lastErr}
}

// ValidateErrorWrap 包装一个解析异常，mplus 内部使用，用于像外部反馈异常及相关信息
func ValidateErrorWrap(err error, errType ValidateErrorType) error {
	return errors.Wrap(ValidateError{errType: errType, lastErr: err}, ValidateErrorTypeMsg[errType])
}

// ValidateErrorFunc 请求解析失败的处理器
type ValidateErrorFunc func(w http.ResponseWriter, r *http.Request, err error)

// ValidateErrorHub 解析错误处理器
var ValidateErrorHub = map[ValidateErrorType]ValidateErrorFunc{
	ErrBodyRead: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
	ErrBodyUnmarshal: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
	ErrMediaType: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusUnsupportedMediaType.Set(err.Error()), http.StatusUnsupportedMediaType)
	},
	ErrMediaTypeParse: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusUnsupportedMediaType.Set(err.Error()), http.StatusUnsupportedMediaType)
	},
	ErrBodyParse: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
	ErrDecode: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
	ErrParseQuery: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
	ErrBodyValidate: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
	ErrRequestValidate: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
	ErrDefault: func(w http.ResponseWriter, r *http.Request, err error) {
		mhttp.CallRegisterFuncOrAbortError(w, r, message.MessageStatusBadRequest.Set(err.Error()), http.StatusBadRequest)
	},
}

// RegisterGlobalValidateErrorHandler 全局解析异常处理器
var GlobalValidateErrorHandler ValidateErrorFunc

// RegisterGlobalValidateErrorHandler 注册全局解析异常处理器
func RegisterGlobalValidateErrorHandler(fun ValidateErrorFunc) {
	GlobalValidateErrorHandler = fun
}

// RegisterValidateErrorFunc 注册一个解析失败的处理器
func RegisterValidateErrorFunc(errType ValidateErrorType, fun ValidateErrorFunc) {
	ValidateErrorHub[errType] = fun
}
