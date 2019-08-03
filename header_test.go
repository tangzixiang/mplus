package mplus

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestSetResponseHeader(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: RequestIDHeader, value: uuid.NewV4().String()}},
		{name: "", args: args{key: ContentTypeHeader, value: ContentTypeJSON}},
	}

	for _, item := range tests {
		w = SetResponseHeader(w, item.args.key, item.args.value)
	}

	for _, item := range tests {
		if !assert.Equal(t, item.args.value, GetResponseHeader(w, item.args.key)) {
			return
		}
	}
}

func TestSetResponseHeaders(t *testing.T) {

	headers := map[string]string{
		RequestIDHeader:   uuid.NewV4().String(),
		ContentTypeHeader: ContentTypeStream,
	}

	w := SetResponseHeaders( // set response header and get back response
		http.ResponseWriter(httptest.NewRecorder()), headers)

	for key, value := range headers {
		if !assert.Equal(t, value, GetResponseHeader(w, key)) {
			return
		}
	}

}

func TestSplitHeader(t *testing.T) {
	// Authorization: Bearer xxx
	headerKey := "Authorization"
	headerValues := []string{"Bearer", "xxx"}

	r := SetRequestHeader( // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		headerKey, strings.Join(headerValues, " "),
	)

	for i, value := range SplitHeader(r, headerKey) {
		if !assert.Equal(t, value, headerValues[i]) {
			return
		}
	}
}

func TestSetRequestHeaders(t *testing.T) {
	headers := map[string]string{
		ContentType:     MIMEJSON,
		"Authorization": "Bearer xxx",
	}

	r := SetRequestHeaders( // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		headers,
	)

	for key, value := range headers {
		if !assert.Equal(t, value, GetHeader(r, key)) {
			return
		}
	}
}

func TestSetRequestHeaderRequestID(t *testing.T) {

	requestID := uuid.NewV4().String()
	r := SetRequestHeaderRequestID( // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		requestID,
	)

	if !assert.Equal(t, requestID, GetHeaderRequestID(r)) {
		return
	}
}

func TestGetClientIP(t *testing.T) {

	localPath := "127.0.0.1:500"
	headers := map[string]string{
		ForwardedForHeader:        localPath,
		RealIPHeader:              localPath,
		AppEngineRemoteAddrHeader: localPath,
	}

	for key, value := range headers {

		r := SetRequestHeader( // set header and get back req
			httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
			key, value,
		)

		if !assert.Equal(t, value, GetClientIP(r)) {
			return
		}
	}
}
