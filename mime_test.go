package mplus

import (
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestParseMediaType(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	SetRequestHeader(request, HeaderContentType, "application/json;charset=utf-8")

	jsonMedia, err := ParseMediaType(request.Header.Get(HeaderContentType))

	assert.Equal(t, err, nil)
	assert.Equal(t, jsonMedia, MIMEJSON)
}

func TestParseRequestMediaType(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	SetRequestHeader(request, HeaderContentType, "application/json;charset=utf-8")

	jsonMedia, err := ParseRequestMediaType(request)

	assert.Equal(t, err, nil)
	assert.Equal(t, jsonMedia, MIMEJSON)
}

func TestParseResponseMediaType(t *testing.T) {

	resp := httptest.NewRecorder()
	SetResponseHeader(resp, HeaderContentType, "application/json;charset=utf-8")

	jsonMedia, err := ParseResponseMediaType(resp)

	assert.Equal(t, err, nil)
	assert.Equal(t, jsonMedia, MIMEJSON)
}
