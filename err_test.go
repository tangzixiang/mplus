package mplus

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidateError_Error(t *testing.T) {
	err := ValidateErrorWrap(errors.New("body parse error"), ErrBodyParse)

	if assert.NotPanics(t, func() { fmt.Println(err.Error()) }) {
		return
	}
}

func TestRegisterValidateErrorFunc(t *testing.T) {

	// read body failed
	RegisterValidateErrorFunc(ErrBodyRead, func(w http.ResponseWriter, r *http.Request, err error) {
		Abort(r)
	})

	SetStrictJSONBodyCheck(true) // 设置严格校验模式

	w := http.ResponseWriter(httptest.NewRecorder())

	r := SetRequestHeader(
		httptest.NewRequest(http.MethodPost, "http://127.0.0.1", strings.NewReader("")), // body len is 0
		ContentType, MIMEJSON) // set parse by json

	// init context
	r = r.WithContext(NewContext(r.Context()))

	type emptyBody struct{}

	// 主动触发
	Bind((*emptyBody)(nil)).ServeHTTP(w, r)

	assert.True(t, IsAbort(r))
}
