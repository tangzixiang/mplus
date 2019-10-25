package message

import (
	"net/http"
	"sync"
)

// MSGType 消息语言类别
type MSGType string

// 支持的语言类别
const (
	MSGLangZH MSGType = "zh"
	MSGLangEN MSGType = "en"
)

var defaultLang = MSGLangEN

type message struct {
	StatusCode int
	ErrorCode  int

	Callback

	lock       sync.Mutex
	MessageStr map[MSGType]string
}

// Message 消息体
type Message interface {
	// AddI18Message 新增一个指定语言的消息
	AddI18Message(msgType MSGType, message string) Message
	// I18nMessage 获取一个指定语言的消息
	I18nMessage(msgType MSGType) string
	// SetStatus 设置当前消息的状态嘛
	SetStatus(statusCode int) Message
	// Status 获取当前消息的状态码
	Status() int
	// SetErrCode 设置当前消息的错误码
	SetErrCode(code int) Message
	// ErrCode 获取当前消息的错误码
	ErrCode() int
	// Copy 拷贝当前消息
	Copy() Message
	// En 获取英文消息
	En() string
	// Default 获取默认语言类型消息
	Default() string
	// Set 设置默认消息
	Set(string) Message
	// SetEn 设置英文消息
	SetEn(string) Message

	Callback
}

var _ Message = &message{}

// NewCallbackMessage 新建一个附带回调的 ErrCode 消息体
func NewCallbackMessage(statusCode, errorCode int, en string, back CallbackMessage) Message {
	m := &message{
		MessageStr: map[MSGType]string{},
		StatusCode: statusCode, ErrorCode: errorCode,
	}

	if en != "" {
		m.SetEn(en)
	}

	if back != nil {
		m.Callback = back
	} else {
		m.Callback = EmptyCallback // prevent panic
	}

	return m
}

// NewErrCodeMessage 新建一个 ErrCode 消息体
func NewErrCodeMessage(statusCode, errorCode int, en string) Message {
	m := &message{
		MessageStr: map[MSGType]string{},
		StatusCode: statusCode, ErrorCode: errorCode,
	}

	if en != "" {
		m.SetEn(en)
	}

	m.Callback = EmptyCallback // prevent panic

	return m
}

// NewMessage 新建一个消息体
func NewMessage(statusCode int, en string) Message {
	m := &message{
		MessageStr: map[MSGType]string{},
		StatusCode: statusCode,
	}

	if en != "" {
		m.SetEn(en)
	}

	m.Callback = EmptyCallback // prevent panic

	return m
}

func (m *message) AddI18Message(msgType MSGType, message string) Message {
	m.lock.Lock()
	m.MessageStr[msgType] = message
	m.lock.Unlock()
	return m
}

func (m *message) I18nMessage(msgType MSGType) string {
	return m.MessageStr[msgType]
}

func (m *message) En() string {
	return m.I18nMessage(MSGLangEN)
}

func (m *message) Default() string {
	return m.I18nMessage(defaultLang)
}

func (m *message) Set(value string) Message {
	return m.AddI18Message(defaultLang, value)
}

func (m *message) SetEn(value string) Message {
	return m.AddI18Message(MSGLangEN, value)
}

func (m *message) SetStatus(statusCode int) Message {
	m.StatusCode = statusCode

	if m.En() == "" {
		m.SetEn(http.StatusText(statusCode))
	}

	return m
}

func (m *message) Copy() Message {
	return NewCallbackMessage(m.Status(), m.ErrCode(), m.Default(), m.Callback.(CallbackMessage))
}

func (m *message) Status() int {
	return m.StatusCode
}

func (m *message) SetErrCode(code int) Message {
	m.ErrorCode = code
	return m
}

func (m *message) ErrCode() int {
	return m.ErrorCode
}

type messages map[int]Message

// Messages Message 消息集合
var (
	lock     = sync.Mutex{}
	Messages = messages{}
)

// Add 添加指定 Message
func (ms messages) Add(msg Message) {
	lock.Lock()
	ms[msg.ErrCode()] = msg
	lock.Unlock()
}

// Get 获取指定 Message
func (ms messages) Get(errCode int) Message {
	return ms[errCode]
}

// SetDefaultLang 设置当前项目默认的语言
func SetDefaultLang(msgType MSGType) {
	defaultLang = msgType
}

// 通用型话术
var (
	MessageStatusOK                   = NewMessage(http.StatusOK, http.StatusText(http.StatusOK))                                     // 200 OK
	MessageStatusCreated              = NewMessage(http.StatusCreated, http.StatusText(http.StatusCreated))                           // 201 Created
	MessageStatusAccepted             = NewMessage(http.StatusAccepted, http.StatusText(http.StatusAccepted))                         // 202 Accepted
	MessageStatusNonAuthoritativeInfo = NewMessage(http.StatusNonAuthoritativeInfo, http.StatusText(http.StatusNonAuthoritativeInfo)) // 203 Non-Authoritative Information
	MessageStatusNoContent            = NewMessage(http.StatusNoContent, http.StatusText(http.StatusNoContent))                       // 204 No Content
	MessageStatusResetContent         = NewMessage(http.StatusResetContent, http.StatusText(http.StatusResetContent))                 // 205 Reset Content
	MessageStatusPartialContent       = NewMessage(http.StatusPartialContent, http.StatusText(http.StatusPartialContent))             // 206 Partial Content
	MessageStatusMultiStatus          = NewMessage(http.StatusMultiStatus, http.StatusText(http.StatusMultiStatus))                   // 207 Multi-Status
	MessageStatusAlreadyReported      = NewMessage(http.StatusAlreadyReported, http.StatusText(http.StatusAlreadyReported))           // 208 Already Reported
	MessageStatusIMUsed               = NewMessage(http.StatusIMUsed, http.StatusText(http.StatusIMUsed))                             // 226 IM Used

	MessageStatusMultipleChoices   = NewMessage(http.StatusMultipleChoices, http.StatusText(http.StatusMultipleChoices))     // 300 Multiple Choices
	MessageStatusMovedPermanently  = NewMessage(http.StatusMovedPermanently, http.StatusText(http.StatusMovedPermanently))   // 301 Moved Permanently
	MessageStatusFound             = NewMessage(http.StatusFound, http.StatusText(http.StatusFound))                         // 302 Found
	MessageStatusSeeOther          = NewMessage(http.StatusSeeOther, http.StatusText(http.StatusSeeOther))                   // 303 See Other
	MessageStatusNotModified       = NewMessage(http.StatusNotModified, http.StatusText(http.StatusNotModified))             // 304 Not Modified
	MessageStatusUseProxy          = NewMessage(http.StatusUseProxy, http.StatusText(http.StatusUseProxy))                   // 305 Use Proxy
	MessageStatusTemporaryRedirect = NewMessage(http.StatusTemporaryRedirect, http.StatusText(http.StatusTemporaryRedirect)) // 307 Temporary Redirect
	MessageStatusPermanentRedirect = NewMessage(http.StatusPermanentRedirect, http.StatusText(http.StatusPermanentRedirect)) // 308 Permanent Redirect

	MessageStatusBadRequest                   = NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))                                     // 400 Bad Request
	MessageStatusUnauthorized                 = NewMessage(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))                                 // 401 Unauthorized
	MessageStatusPaymentRequired              = NewMessage(http.StatusPaymentRequired, http.StatusText(http.StatusPaymentRequired))                           // 402 Payment Required
	MessageStatusForbidden                    = NewMessage(http.StatusForbidden, http.StatusText(http.StatusForbidden))                                       // 403 Forbidden
	MessageStatusNotFound                     = NewMessage(http.StatusNotFound, http.StatusText(http.StatusNotFound))                                         // 404 Not Found
	MessageStatusMethodNotAllowed             = NewMessage(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))                         // 405 Method Not Allowed
	MessageStatusNotAcceptable                = NewMessage(http.StatusNotAcceptable, http.StatusText(http.StatusNotAcceptable))                               // 406 Not Acceptable
	MessageStatusProxyAuthRequired            = NewMessage(http.StatusProxyAuthRequired, http.StatusText(http.StatusProxyAuthRequired))                       // 407 Proxy Authentication Required
	MessageStatusRequestTimeout               = NewMessage(http.StatusRequestTimeout, http.StatusText(http.StatusRequestTimeout))                             // 408 Request Timeout
	MessageStatusConflict                     = NewMessage(http.StatusConflict, http.StatusText(http.StatusConflict))                                         // 409 Conflict
	MessageStatusGone                         = NewMessage(http.StatusGone, http.StatusText(http.StatusGone))                                                 // 410 Gone
	MessageStatusLengthRequired               = NewMessage(http.StatusLengthRequired, http.StatusText(http.StatusLengthRequired))                             // 411 Length Required
	MessageStatusPreconditionFailed           = NewMessage(http.StatusPreconditionFailed, http.StatusText(http.StatusPreconditionFailed))                     // 412 Precondition Failed
	MessageStatusRequestEntityTooLarge        = NewMessage(http.StatusRequestEntityTooLarge, http.StatusText(http.StatusRequestEntityTooLarge))               // 413 Request Entity Too Large
	MessageStatusRequestURITooLong            = NewMessage(http.StatusRequestURITooLong, http.StatusText(http.StatusRequestURITooLong))                       // 414 Request URI Too Long
	MessageStatusUnsupportedMediaType         = NewMessage(http.StatusUnsupportedMediaType, http.StatusText(http.StatusUnsupportedMediaType))                 // 415 Unsupported Media Type
	MessageStatusRequestedRangeNotSatisfiable = NewMessage(http.StatusRequestedRangeNotSatisfiable, http.StatusText(http.StatusRequestedRangeNotSatisfiable)) // 416 Requested Range Not Satisfiable
	MessageStatusExpectationFailed            = NewMessage(http.StatusExpectationFailed, http.StatusText(http.StatusExpectationFailed))                       // 417 Expectation Failed
	MessageStatusTeapot                       = NewMessage(http.StatusTeapot, http.StatusText(http.StatusTeapot))                                             // 418 I'm a teapot
	MessageStatusMisdirectedRequest           = NewMessage(http.StatusMisdirectedRequest, http.StatusText(http.StatusMisdirectedRequest))                     // 421 Misdirected Request
	MessageStatusUnprocessableEntity          = NewMessage(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))                   // 422 Unprocessable Entity
	MessageStatusLocked                       = NewMessage(http.StatusLocked, http.StatusText(http.StatusLocked))                                             // 423 Locked
	MessageStatusFailedDependency             = NewMessage(http.StatusFailedDependency, http.StatusText(http.StatusFailedDependency))                         // 424 Failed Dependency
	MessageStatusTooEarly                     = NewMessage(http.StatusTooEarly, http.StatusText(http.StatusTooEarly))                                         // 425 Too Early
	MessageStatusUpgradeRequired              = NewMessage(http.StatusUpgradeRequired, http.StatusText(http.StatusUpgradeRequired))                           // 426 Upgrade Required
	MessageStatusPreconditionRequired         = NewMessage(http.StatusPreconditionRequired, http.StatusText(http.StatusPreconditionRequired))                 // 428 Precondition Required
	MessageStatusTooManyRequests              = NewMessage(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))                           // 429 Too Many Requests
	MessageStatusRequestHeaderFieldsTooLarge  = NewMessage(http.StatusRequestHeaderFieldsTooLarge, http.StatusText(http.StatusRequestHeaderFieldsTooLarge))   // 431 Request Header Fields Too Large
	MessageStatusUnavailableForLegalReasons   = NewMessage(http.StatusUnavailableForLegalReasons, http.StatusText(http.StatusUnavailableForLegalReasons))     // 451 Unavailable For Legal Reasons

	MessageStatusInternalServerError           = NewMessage(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))                     // 500 Internal Server Error
	MessageStatusNotImplemented                = NewMessage(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))                               // 501 Not Implemented
	MessageStatusBadGateway                    = NewMessage(http.StatusBadGateway, http.StatusText(http.StatusBadGateway))                                       // 502 Bad Gateway
	MessageStatusServiceUnavailable            = NewMessage(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable))                       // 503 Service Unavailable
	MessageStatusGatewayTimeout                = NewMessage(http.StatusGatewayTimeout, http.StatusText(http.StatusGatewayTimeout))                               // 504 Gateway Timeout
	MessageStatusHTTPVersionNotSupported       = NewMessage(http.StatusHTTPVersionNotSupported, http.StatusText(http.StatusHTTPVersionNotSupported))             // 505 HTTP Version Not Supported
	MessageStatusVariantAlsoNegotiates         = NewMessage(http.StatusVariantAlsoNegotiates, http.StatusText(http.StatusVariantAlsoNegotiates))                 // 506 Variant Also Negotiates
	MessageStatusInsufficientStorage           = NewMessage(http.StatusInsufficientStorage, http.StatusText(http.StatusInsufficientStorage))                     // 507 Insufficient Storage
	MessageStatusLoopDetected                  = NewMessage(http.StatusLoopDetected, http.StatusText(http.StatusLoopDetected))                                   // 508 Loop Detected
	MessageStatusNotExtended                   = NewMessage(http.StatusNotExtended, http.StatusText(http.StatusNotExtended))                                     // 510 Not Extended
	MessageStatusNetworkAuthenticationRequired = NewMessage(http.StatusNetworkAuthenticationRequired, http.StatusText(http.StatusNetworkAuthenticationRequired)) // 511 Network Authentication Required
)
