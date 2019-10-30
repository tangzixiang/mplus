package message

import (
	"net/http"
)

// Callback 消息处理回调，直接和 plus.PP.CallbackByCode 配合
type Callback interface {
	Do(w http.ResponseWriter, r *http.Request, m Message, respData interface{})
}

// CallbackMessage 为当前异常错误消息注册一个处理回调
type CallbackMessage func(w http.ResponseWriter, r *http.Request, m Message, respData interface{})

func (c CallbackMessage) Do(w http.ResponseWriter, r *http.Request, m Message, respData interface{}) {
	c(w, r, m, respData)
}

// EmptyCallback 空回调
var EmptyCallback CallbackMessage = func(w http.ResponseWriter, r *http.Request, m Message, respData interface{}) {}
