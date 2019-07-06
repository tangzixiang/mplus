package mplus

import (
	"testing"

	"net/http"
	"net/http/httptest"

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