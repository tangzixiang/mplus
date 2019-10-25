package mplus

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
	assert "github.com/stretchr/testify/require"
)

func TestParseValidate(t *testing.T) {
	jsonStr := `{"name":"Tom","age":50,"gender":"male","email":"10086@fox.com"}`
	buffer := bytes.NewBuffer([]byte(jsonStr))

	request := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(request, HeaderContentType, MIMEJSON)

	var vr ValidateResult
	var user User

	// 1. 解析数据
	Parse(request, &vr)
	assert.Nil(t, vr.Err)

	// 绑定数据
	DecodeTo(request, &user, &vr)
	assert.Nil(t, vr.Err)

	// 数据校验
	BindValidate(request, &user, &vr)

	assert.Equal(t, vr.Err, nil)
	assert.Equal(t, user.Name, "Tom")
	assert.Equal(t, user.Age, uint8(50))
	assert.Equal(t, user.Gender, "male")
	assert.Equal(t, user.Email, "10086@fox.com")
}

func TestParseValidateErr(t *testing.T) {
	jsonStr := `{"name":"Tom","age":50,"gender":"","email":"10086@fox.com"}`
	buffer := bytes.NewBuffer([]byte(jsonStr))

	request := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(request, HeaderContentType, MIMEJSON)

	var vr ValidateResult
	var user User

	// 1. 解析数据
	Parse(request, &vr)
	assert.Nil(t, vr.Err)

	// 绑定数据
	DecodeTo(request, &user, &vr)
	assert.Nil(t, vr.Err)

	// 数据校验
	BindValidate(request, &user, &vr)
	assert.NotEqual(t, vr.Err, nil)
	assert.IsType(t, ValidateError{}, errors.Cause(vr.Err))
	assert.Equal(t, ErrBodyValidate, errors.Cause(vr.Err).(ValidateError).Type())
}

func TestParseValidateErrMediaTypeParse(t *testing.T) {
	jsonStr := `(function(){})()`
	buffer := bytes.NewBuffer([]byte(jsonStr))

	request := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(request, HeaderContentType, "application/javascript;xxx")

	var vr ValidateResult

	// 1. 解析数据
	Parse(request, &vr)

	assert.NotEqual(t, vr.Err, nil)
	assert.IsType(t, ValidateError{}, errors.Cause(vr.Err))
	assert.Equal(t, ErrMediaTypeParse, errors.Cause(vr.Err).(ValidateError).Type())
}

func TestParseValidateErrMediaType(t *testing.T) {
	jsonStr := `(function(){})()`
	buffer := bytes.NewBuffer([]byte(jsonStr))

	request := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(request, HeaderContentType, "application/javascript")

	var vr ValidateResult

	// 1. 解析数据
	Parse(request, &vr)

	assert.NotEqual(t, vr.Err, nil)
	assert.IsType(t, ValidateError{}, errors.Cause(vr.Err))
	assert.Equal(t, ErrMediaType, errors.Cause(vr.Err).(ValidateError).Type())

}

func TestValidatorStandErrMsg(t *testing.T) {

	// Parse body failed
	RegisterValidateErrorFunc(ErrBodyValidate, func(w http.ResponseWriter, r *http.Request, err error) {
		Abort(r)

		cErr, ok := err.(ValidateError)
		assert.True(t, ok)

		w.Write([]byte(cErr.Error()))
	})

	recorder := httptest.NewRecorder()
	// wrap the ResponseWriter
	w := NewResponseWrite(http.ResponseWriter(recorder))

	r := SetRequestHeader(
		httptest.NewRequest(http.MethodPost, "http://127.0.0.1", strings.NewReader("{}")),
		HeaderContentType, MIMEJSON) // set Parse by json

	// init context
	r = r.WithContext(NewContext(r.Context()))

	type body struct {
		Size int `validate:"required"`
	}

	// 主动触发
	Bind((*body)(nil)).ServeHTTP(w, r)

	errMsg := "Key: 'body.Size' Error:Field validation for 'Size' failed on the 'required' tag"
	assert.Equal(t, errMsg, recorder.Body.String())
}
