package mplus

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbort(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "", nil)
	assert.Equal(t, err, nil)

	req = req.WithContext(NewContext(req.Context()))
	assert.Equal(t, IsAbort(req), false)
	assert.Equal(t, IsAbort(Abort(req)), true)
}

func TestError(t *testing.T) {
	respR := httptest.NewRecorder()

	respW := NewResponseWrite(respR)
	assert.Equal(t, respW.Status(), http.StatusOK)

	Error(respW, MessageStatusInternalServerError)
	assert.Equal(t, respW.Status(), http.StatusInternalServerError)
	assert.Equal(t, respR.Result().StatusCode, http.StatusInternalServerError)
}

func TestDumpRequest(t *testing.T) {

	content := NewQuery().AddPairs("name", "tom", "age", "18").Encode()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(content))

	if !assert.Equal(t, content, DumpRequest(req)) {
		return
	}

	bodyBytes,err := ioutil.ReadAll(req.Body)
	if !assert.Nil(t,err) {
		return
	}

	if !assert.Equal(t,content,string(bodyBytes)) {
		return
	}
}
