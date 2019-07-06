package mplus

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMediaType(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	SetRequestHeader(request, ContentType, "application/json;charset=utf-8")

	jsonMedia, err := ParseMediaType(request)

	assert.Equal(t, err, nil)
	assert.Equal(t, jsonMedia, MIMEJSON)
}
