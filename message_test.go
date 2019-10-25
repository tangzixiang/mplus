package mplus

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/tangzixiang/mplus/message"
)

func TestNewCallbackMessage(t *testing.T) {

	m := NewCallbackMessage(http.StatusBadRequest, 400001, http.StatusText(http.StatusBadRequest), func(w http.ResponseWriter, r *http.Request, m message.Message, respData interface{}) {
		PlusPlus(w, r).JSON(Data{"data": respData}, m.Status())
	})

	assert.Equal(t, http.StatusBadRequest, m.Status())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.Default())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.I18nMessage(MSGLangEN))
	assert.Equal(t, "", m.I18nMessage(MSGLangZH))
	assert.Equal(t, 400001, m.ErrCode())

	Messages.Add(m)
	defer delete(Messages, m.ErrCode())

	respR := httptest.NewRecorder()
	w := NewResponseWrite(respR)
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	m.Do(w, r, m, Data{"code": 400001})

	assert.True(t, IsAbort(r))
	assert.Equal(t, http.StatusBadRequest, GetHTTPRespStatus(w))

	bodyBytes, err := ioutil.ReadAll(respR.Result().Body)
	assert.Nil(t, err)

	assert.Equal(t, `{"data":{"code":400001}}`, string(bodyBytes))
}

func TestNewErrCodeMessage(t *testing.T) {

	// no callback
	m := NewErrCodeMessage(http.StatusBadRequest, 400001, http.StatusText(http.StatusBadRequest))

	assert.Equal(t, http.StatusBadRequest, m.Status())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.Default())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.I18nMessage(MSGLangEN))
	assert.Equal(t, "", m.I18nMessage(MSGLangZH))
	assert.Equal(t, 400001, m.ErrCode())

	Messages.Add(m)
	defer delete(Messages, m.ErrCode())

	respR := httptest.NewRecorder()
	w := NewResponseWrite(respR)
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	// nothing happen
	m.Do(w, r, m, Data{"code": 400001})

	assert.False(t, IsAbort(r))
	assert.NotEqual(t, http.StatusBadRequest, GetHTTPRespStatus(w))
}

func TestNewMessage(t *testing.T) {
	// no callback no errorCode
	m := NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))

	assert.Equal(t, http.StatusBadRequest, m.Status())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.Default())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.I18nMessage(MSGLangEN))
	assert.Equal(t, "", m.I18nMessage(MSGLangZH))
	assert.Equal(t, 0, m.ErrCode())

	Messages.Add(m)
	defer delete(Messages, m.ErrCode())

	respR := httptest.NewRecorder()
	w := NewResponseWrite(respR)
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	// nothing happen
	m.Do(w, r, m, Data{"code": 400001})

	assert.False(t, IsAbort(r))
	assert.NotEqual(t, http.StatusBadRequest, GetHTTPRespStatus(w))
}

func Test_message_AddI18Message(t *testing.T) {

	m := NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))

	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.I18nMessage(MSGLangEN))

	m.AddI18Message(MSGLangZH, "请求体错误")
	assert.Equal(t, "请求体错误", m.I18nMessage(MSGLangZH))
}

func Test_message_En(t *testing.T) {
	assert.Equal(t, http.StatusText(http.StatusBadRequest), NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).En())
	assert.Equal(t, "", NewMessage(http.StatusBadRequest, "").En())
}

func Test_message_Default(t *testing.T) {
	assert.Equal(t, http.StatusText(http.StatusBadRequest), NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Default())
	assert.Equal(t, "", NewMessage(http.StatusBadRequest, "").Default())
}

func Test_message_Set(t *testing.T) {
	assert.Equal(t, "", NewMessage(http.StatusBadRequest, "").En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), NewMessage(http.StatusBadRequest, "").Set(http.StatusText(http.StatusBadRequest)).En())

	SetDefaultLang(MSGLangZH)
	defer SetDefaultLang(MSGLangEN)

	assert.Equal(t, "", NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Default())
	assert.Equal(t, "请求体错误", NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Set("请求体错误").Default())
}

func Test_message_SetEn(t *testing.T) {
	assert.Equal(t, "", NewMessage(http.StatusBadRequest, "").En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), NewMessage(http.StatusBadRequest, "").SetEn(http.StatusText(http.StatusBadRequest)).En())
}

func Test_message_SetStatus(t *testing.T) {

	assert.Equal(t, http.StatusOK, NewMessage(http.StatusOK, http.StatusText(http.StatusOK)).Status())
	assert.Equal(t, http.StatusBadRequest, NewMessage(http.StatusOK, http.StatusText(http.StatusOK)).SetStatus(http.StatusBadRequest).Set(http.StatusText(http.StatusBadRequest)).Status())
}

func Test_message_Copy(t *testing.T) {

	// no callback
	m := NewErrCodeMessage(http.StatusBadRequest, 400001, http.StatusText(http.StatusBadRequest))

	assert.Equal(t, http.StatusBadRequest, m.Status())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.Default())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.I18nMessage(MSGLangEN))
	assert.Equal(t, "", m.I18nMessage(MSGLangZH))
	assert.Equal(t, 400001, m.ErrCode())

	m2 := m.Copy()

	assert.NotEqual(t, m, m2)
	assert.Equal(t, http.StatusBadRequest, m2.Status())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m2.Default())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m2.En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m2.I18nMessage(MSGLangEN))
	assert.Equal(t, "", m2.I18nMessage(MSGLangZH))
	assert.Equal(t, 400001, m2.ErrCode())

}

func Test_message_SetErrCode(t *testing.T) {

	assert.Equal(t, 400001, NewErrCodeMessage(http.StatusBadRequest, 400001, http.StatusText(http.StatusBadRequest)).ErrCode())
	assert.Equal(t, 400002, NewErrCodeMessage(http.StatusBadRequest, 400001, http.StatusText(http.StatusOK)).SetErrCode(400002).ErrCode())
}

func Test_messages_Add(t *testing.T) {

	m := NewCallbackMessage(http.StatusBadRequest, 400001, http.StatusText(http.StatusBadRequest), func(w http.ResponseWriter, r *http.Request, m message.Message, respData interface{}) {
		PlusPlus(w, r).JSON(Data{"data": respData}, m.Status())
	})

	Messages.Add(m)
	assert.Equal(t, m, Messages.Get(m.ErrCode()))
}

func TestSetDefaultLang(t *testing.T) {
	assert.Equal(t, "", NewMessage(http.StatusBadRequest, "").En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), NewMessage(http.StatusBadRequest, "").Set(http.StatusText(http.StatusBadRequest)).En())

	SetDefaultLang(MSGLangZH)
	defer SetDefaultLang(MSGLangEN)

	assert.Equal(t, "", NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Default())
	assert.Equal(t, "请求体错误", NewMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).Set("请求体错误").Default())
}

func TestCallbackMessage_Do(t *testing.T) {

	// no callback
	m := NewErrCodeMessage(http.StatusBadRequest, 400001, http.StatusText(http.StatusBadRequest))

	assert.Equal(t, http.StatusBadRequest, m.Status())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.Default())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.En())
	assert.Equal(t, http.StatusText(http.StatusBadRequest), m.I18nMessage(MSGLangEN))
	assert.Equal(t, "", m.I18nMessage(MSGLangZH))
	assert.Equal(t, 400001, m.ErrCode())

	Messages.Add(m)
	defer delete(Messages, m.ErrCode())

	respR := httptest.NewRecorder()
	w := NewResponseWrite(respR)
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	// nothing happen
	m.Do(w, r, m, Data{"code": 400001})

	assert.False(t, IsAbort(r))
	assert.NotEqual(t, http.StatusBadRequest, GetHTTPRespStatus(w))
}
