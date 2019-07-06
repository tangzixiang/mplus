package mplus

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyRequest(t *testing.T) {

	jsonstr := `{"name":"Tom","age":50,"gender":"male","email":"10086@fox.com"}`
	buffer := bytes.NewBuffer([]byte(jsonstr))

	req01 := httptest.NewRequest(http.MethodPost, "http://localhost", buffer)
	SetRequestHeader(req01, ContentType, MIMEJSON)

	req01 = req01.WithContext(NewContext(req01.Context()))
	req02 := CopyRequest(req01)

	assert.Equal(t, GetContextValue(req01.Context(), "1"), GetContextValue(req02.Context(), "1"))
	assert.Equal(t, GetHeader(req01, ContentType), GetHeader(req02, ContentType))

	SetContextValue(req01.Context(), "1", "1")
	SetContextValue(req01.Context(), "1", "2")
	assert.NotEqual(t, GetContextValue(req01.Context(), "1"), GetContextValue(req02.Context(), "1"))

	req01.Header.Del(ContentType)
	assert.NotEqual(t, GetHeader(req01, ContentType), GetHeader(req02, ContentType))
}
