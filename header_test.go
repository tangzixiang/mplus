package mplus

import (
	"net/http"
	"net/http/httptest"
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

	w := SetResponseHeaders(http.ResponseWriter(httptest.NewRecorder()), headers)

	for key, value := range headers {
		if !assert.Equal(t, value, GetResponseHeader(w, key)) {
			return
		}
	}

}
