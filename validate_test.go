package mplus

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestParseValidate(t *testing.T) {
	jsonstr := `{"name":"Tom","age":50,"gender":"male","email":"10086@fox.com"}`
	buffer := bytes.NewBuffer([]byte(jsonstr))

	requset := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(requset, ContentType, MIMEJSON)

	var user User
	vr := ParseValidate(requset, &user)

	if !assert.Equal(t, vr.Err, nil) {
		return
	}

	if !assert.Equal(t, user.Name, "Tom") {
		return
	}

	if !assert.Equal(t, user.Age, uint8(50)) {
		return
	}

	if !assert.Equal(t, user.Gender, "male") {
		return
	}

	if !assert.Equal(t, user.Email, "10086@fox.com") {
		return
	}
}

func TestParseValidateErr(t *testing.T) {
	jsonstr := `{"name":"Tom","age":50,"gender":"","email":"10086@fox.com"}`
	buffer := bytes.NewBuffer([]byte(jsonstr))

	requset := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(requset, ContentType, MIMEJSON)

	var user User
	vr := ParseValidate(requset, &user)

	if !assert.NotEqual(t, vr.Err, nil) {
		return
	} //

	if !assert.IsType(t, ValidateError{}, errors.Cause(vr.Err)) {
		return
	}

	if !assert.Equal(t, ErrBodyValidate, errors.Cause(vr.Err).(ValidateError).Type()) {
		return
	}
}

func TestParseValidateErrMediaTypeParse(t *testing.T) {
	jsonstr := `(function(){})()`
	buffer := bytes.NewBuffer([]byte(jsonstr))

	requset := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(requset, ContentType, "application/javascript;xxx")

	var user User
	vr := ParseValidate(requset, &user)

	if !assert.NotEqual(t, vr.Err, nil) {
		return
	}
	if !assert.IsType(t, ValidateError{}, errors.Cause(vr.Err)) {
		return
	}

	if !assert.Equal(t, ErrMediaTypeParse, errors.Cause(vr.Err).(ValidateError).Type()) {
		return
	}
}

func TestParseValidateErrMediaType(t *testing.T) {
	jsonstr := `(function(){})()`
	buffer := bytes.NewBuffer([]byte(jsonstr))

	requset := httptest.NewRequest(http.MethodPost, "http://localhost?query=test", buffer)
	SetRequestHeader(requset, ContentType, "application/javascript")

	var user User
	vr := ParseValidate(requset, &user)

	if !assert.NotEqual(t, vr.Err, nil) {
		return
	}

	if !assert.IsType(t, ValidateError{}, errors.Cause(vr.Err)) {
		return
	}

	if !assert.Equal(t, ErrMediaType, errors.Cause(vr.Err).(ValidateError).Type()) {
		return
	}
}

func TestValidatorStandErrMsg(t *testing.T) {

	// parse body failed
	RegisterValidateErrorFunc(ErrBodyValidate, func(w http.ResponseWriter, r *http.Request, err error) {
		Abort(r)

		cErr, ok := err.(ValidateError)
		if !assert.True(t, ok) {
			return
		}

		w.Write([]byte(cErr.Error()))
	})

	recorder := httptest.NewRecorder()
	// wrap the ResponseWriter
	w := NewResponseWrite(http.ResponseWriter(recorder))

	r := SetRequestHeader(
		httptest.NewRequest(http.MethodPost, "http://127.0.0.1", strings.NewReader("{}")),
		ContentType, MIMEJSON) // set parse by json

	// init context
	r = r.WithContext(NewContext(r.Context()))

	type body struct {
		Size int `validate:"required"`
	}

	// 主动触发
	Bind((*body)(nil)).ServeHTTP(w, r)

	errMsg := "Key: 'body.Size' Error:Field validation for 'Size' failed on the 'required' tag"
	if !assert.Equal(t, errMsg, recorder.Body.String()) {
		return
	}
}
