package mplus

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

// 为当前异常错误消息注册一个处理回调
type CallBack func(pp *PP, m Message, respData interface{})

type message struct {
	StatusCode int
	ErrorCode  int

	callback   CallBack
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
	SetStatus(code int) Message
	// Status 获取当前消息的状态码
	Status() int
	// SetErrCode 设置当前消息的错误码
	SetErrCode(code int) Message
	// ErrCode 获取当前消息的错误码
	ErrCode() int
	// Copy 拷贝当前消息
	Copy() Message
	// Zh 获取中文消息
	Zh() string
	// En 获取英文消息
	En() string
	// Default 获取默认语言类型消息
	Default() string
	// Set 设置默认消息
	Set(string) Message
	// SetZh 设置中文消息
	SetZh(string) Message
	// SetEn 设置英文消息
	SetEn(string) Message
}

var _ Message = &message{}

// NewMessage 新建一个消息体
func NewMessage(statusCode, errorCode int, zh, en string, back CallBack) Message {
	m := &message{
		MessageStr: map[MSGType]string{},
		StatusCode: statusCode, ErrorCode: errorCode,
	}

	if zh != "" {
		m.SetZh(zh)
	}

	if en != "" {
		m.SetEn(en)
	}

	if back != nil {
		m.callback = back
	}

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

func (m *message) Zh() string {
	return m.I18nMessage(MSGLangZH)
}

func (m *message) En() string {
	return m.I18nMessage(MSGLangEN)
}

func (m *message) Default() string {
	msg := m.I18nMessage(defaultLang)
	if msg == "" {
		msg = m.En()
	}
	return msg
}

func (m *message) Set(value string) Message {
	return m.AddI18Message(defaultLang, value)
}

func (m *message) SetZh(value string) Message {
	return m.AddI18Message(MSGLangZH, value)
}

func (m *message) SetEn(value string) Message {
	return m.AddI18Message(MSGLangEN, value)
}

func (m *message) SetStatus(code int) Message {
	m.StatusCode = code

	if m.En() == "" {
		m.SetEn(http.StatusText(code))
	}

	if m.Zh() == "" {
		m.SetZh(http.StatusText(code))
	}

	return m
}

func (m *message) Copy() Message {
	newMessage := NewMessage(m.Status(), m.ErrCode(), "", "", nil)

	for key, value := range m.MessageStr {
		newMessage.AddI18Message(key, value)
	}

	return newMessage
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
	MessageStatusOK                  = NewMessage(http.StatusOK, 0, "请求成功", http.StatusText(http.StatusOK), nil)
	MessageStatusBadRequest          = NewMessage(http.StatusBadRequest, 0, "请求体错误", http.StatusText(http.StatusBadRequest), nil)
	MessageStatusForbidden           = NewMessage(http.StatusForbidden, 0, "无权限请求", http.StatusText(http.StatusForbidden), nil)
	MessageStatusNotFound           = NewMessage(http.StatusNotFound, 0, "资源不存在", http.StatusText(http.StatusNotFound), nil)
	MessageStatusUnauthorized        = NewMessage(http.StatusUnauthorized, 0, "认证信息无效或已失效", http.StatusText(http.StatusUnauthorized), nil)
	MessageUnsupportedMediaType      = NewMessage(http.StatusUnsupportedMediaType, 0, "非法的 media-type", http.StatusText(http.StatusUnsupportedMediaType), nil)
	MessageStatusInternalServerError = NewMessage(http.StatusInternalServerError, 0, "服务器内部错误", http.StatusText(http.StatusInternalServerError), nil)
)
