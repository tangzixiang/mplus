package mplus

import (
	"bytes"
	"net/http"
	"net/http/httptest"
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
