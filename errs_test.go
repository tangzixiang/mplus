package mplus

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
	assert "github.com/stretchr/testify/require"
)

func TestValidateError_Error(t *testing.T) {
	assert.NotPanics(t, func() { _ = ValidateErrorWrap(errors.New("body parse error"), ErrBodyParse).Error() })
}

func TestRegisterValidateErrorFunc(t *testing.T) {

	var catchErr error

	// read body failed
	RegisterValidateErrorFunc(ErrBodyRead, func(w http.ResponseWriter, r *http.Request, err error) {
		Abort(r)
		catchErr = err
	})

	// 使用通道模拟锁的机制防止与其他测试冲突
	SetStrictJSONBodyCheckLockChan <- struct{}{}
	SetStrictJSONBodyCheck(true) // 设置严格校验模式，当前模式下若无法读取 JSON 数据则会触发异常
	defer func() {
		SetStrictJSONBodyCheck(false)
		<-SetStrictJSONBodyCheckLockChan
	}()

	w := http.ResponseWriter(httptest.NewRecorder())

	r := httptest.NewRequest(http.MethodPost, "http://127.0.0.1", strings.NewReader("")) // body len is 0
	r = SetRequestHeader(r, HeaderContentType, MIMEJSON)                                 // parse by json
	r = r.WithContext(NewContext(r.Context()))                                           // init context

	type emptyBody struct{}

	// 主动触发
	Bind((*emptyBody)(nil)).ServeHTTP(w, r)

	assert.True(t, IsAbort(r))
	assert.IsType(t, catchErr, ValidateError{})

	cErr := catchErr.(ValidateError)
	if !cErr.IsErr(ErrBodyRead) {
		return
	}

	assert.Equal(t, "body empty", cErr.Error())
}

func TestValidateErrorWrap(t *testing.T) {
	err := errors.New("TestValidateErrorWrap")
	errType := ErrBodyRead
	wrapErr := ValidateErrorWrap(err, errType)

	assert.Equal(t, ValidateErrorTypeMsg[errType]+": TestValidateErrorWrap", wrapErr.Error())
	assert.Equal(t, "TestValidateErrorWrap", errors.Cause(wrapErr).(ValidateError).Error())
	assert.Equal(t, err, errors.Cause(wrapErr).(ValidateError).LastErr())
}
