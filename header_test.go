package mplus

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/satori/go.uuid"
	assert "github.com/stretchr/testify/require"
)

func TestGetHeaderValues(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)

	AddRequestHeader(r, HeaderAcceptEncoding, "gzip")
	AddRequestHeader(r, HeaderAcceptEncoding, "deflate")

	// Accept-Encoding: gzip
	// Accept-Encoding: deflate
	assert.ElementsMatch(t, []string{"gzip", "deflate"}, GetHeaderValues(r, HeaderAcceptEncoding))
}

func TestSplitHeader(t *testing.T) {
	// Authorization: Bearer xxx
	headerKey := "Authorization"
	headerValues := []string{"Bearer", "xxx"}

	r := SetRequestHeader( // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		headerKey, strings.Join(headerValues, " "),
	)

	for i, value := range SplitHeader(r, headerKey, SplitSepBlankSpace) {
		assert.Equal(t, value, headerValues[i])
	}
}

func TestGetHeaderRequestID(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)

	requestID := uuid.Must(uuid.NewV4()).String()

	SetRequestHeader(r, HeaderRequestID, requestID)
	assert.Equal(t, requestID, GetHeaderRequestID(r))
}

func TestSetRequestHeaderIf(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)

	requestID := uuid.Must(uuid.NewV4()).String()
	SetRequestHeader(r, HeaderRequestID, requestID)

	newRequestID := uuid.Must(uuid.NewV4()).String()
	SetRequestHeaderIf(GetHeaderRequestID(r) == "", r, HeaderRequestID, newRequestID)
	assert.Equal(t, requestID, GetHeaderRequestID(r))
}

func TestSetRequestHeaders(t *testing.T) {
	headers := map[string]string{
		HeaderContentType:   MIMEJSON,
		HeaderAuthorization: "Bearer xxx",
	}

	r := SetRequestHeaders( // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		headers,
	)

	for key, value := range headers {
		assert.Equal(t, value, GetHeader(r, key))
	}
}

func TestSetRequestHeadersIf(t *testing.T) {
	headers := map[string]string{
		HeaderContentType:   MIMEJSON,
		HeaderAuthorization: "Bearer xxx",
	}

	r := SetRequestHeadersIf(false, // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		headers,
	)

	for key, value := range headers {
		assert.NotEqual(t, value, GetHeader(r, key))
	}

	r = SetRequestHeadersIf(true, // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		headers,
	)

	for key, value := range headers {
		assert.Equal(t, value, GetHeader(r, key))
	}
}

func TestAddRequestHeader(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)

	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderAcceptEncoding, value: "gzip"}},
		{name: "", args: args{key: HeaderAcceptEncoding, value: "deflate"}},
	}

	for _, item := range tests {
		r = AddRequestHeader(r, item.args.key, item.args.value)
	}

	assert.Equal(t, "gzip", GetHeader(r, HeaderAcceptEncoding))

	// Accept-Encoding: gzip
	// Accept-Encoding: deflate
	assert.ElementsMatch(t, []string{"gzip", "deflate"}, GetHeaderValues(r, HeaderAcceptEncoding))
}

func TestAddRequestHeaderIf(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)

	type args struct {
		ensure bool
		key    string
		values []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderAcceptEncoding, values: []string{"gzip"}, ensure: true}},
		{name: "", args: args{key: HeaderAcceptEncoding, values: []string{"deflate"}, ensure: false}},
	}

	for _, item := range tests {
		r = AddRequestHeaderIf(item.args.ensure, r, item.args.key, item.args.values...)
	}

	assert.Equal(t, "gzip", GetHeader(r, HeaderAcceptEncoding))

	assert.Len(t, GetHeaderValues(r, HeaderAcceptEncoding), 1)
}

func TestAddRequestHeaders(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)

	headers01 := map[string]string{
		HeaderContentType:    MIMEJSON,
		HeaderAcceptEncoding: "gzip",
	}

	headers02 := map[string]string{
		HeaderAcceptEncoding: "deflate",
	}

	r = AddRequestHeaders(r, headers01)
	r = AddRequestHeaders(r, headers02)

	assert.Equal(t, "gzip", GetHeader(r, HeaderAcceptEncoding))

	// Accept-Encoding: gzip
	// Accept-Encoding: deflate
	assert.ElementsMatch(t, []string{"gzip", "deflate"}, GetHeaderValues(r, HeaderAcceptEncoding))
}

func TestAddRequestHeadersIf(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)

	headers01 := map[string]string{
		HeaderContentType:    MIMEJSON,
		HeaderAcceptEncoding: "gzip",
	}

	headers02 := map[string]string{
		HeaderAcceptEncoding: "deflate",
	}

	r = AddRequestHeadersIf(true, r, headers01)
	r = AddRequestHeadersIf(false, r, headers02)

	assert.Equal(t, "gzip", GetHeader(r, HeaderAcceptEncoding))
	assert.Len(t, GetHeaderValues(r, HeaderAcceptEncoding), 1)
}

func TestGetResponseHeader(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderRequestID, value: uuid.Must(uuid.NewV4()).String()}},
		{name: "", args: args{key: HeaderContentType, value: ContentTypeJSON}},
	}

	for _, item := range tests {
		w = SetResponseHeader(w, item.args.key, item.args.value)
	}

	for _, item := range tests {
		assert.Equal(t, item.args.value, GetResponseHeader(w, item.args.key))
	}
}

func TestGetResponseHeaderValues(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderContentEncoding, value: "gzip"}},
		{name: "", args: args{key: HeaderContentEncoding, value: "deflate"}},
	}

	for _, item := range tests {
		w = AddResponseHeader(w, item.args.key, item.args.value)
	}

	assert.Equal(t, "gzip", GetResponseHeader(w, HeaderContentEncoding))
	assert.ElementsMatch(t, []string{"gzip", "deflate"}, GetResponseHeaderValues(w, HeaderContentEncoding))
}

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
		{name: "", args: args{key: HeaderRequestID, value: uuid.Must(uuid.NewV4()).String()}},
		{name: "", args: args{key: HeaderContentType, value: ContentTypeJSON}},
	}

	for _, item := range tests {
		w = SetResponseHeader(w, item.args.key, item.args.value)
	}

	for _, item := range tests {
		assert.Equal(t, item.args.value, GetResponseHeader(w, item.args.key))
	}
}

func TestSetResponseHeaders(t *testing.T) {

	headers := map[string]string{
		HeaderRequestID:   uuid.Must(uuid.NewV4()).String(),
		HeaderContentType: ContentTypeStream,
	}

	w := SetResponseHeaders( // set response header and get back response
		http.ResponseWriter(httptest.NewRecorder()), headers)

	for key, value := range headers {
		assert.Equal(t, value, GetResponseHeader(w, key))
	}

}

func TestSetResponseHeaderIf(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	type args struct {
		key    string
		value  string
		ensure bool
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderRequestID, value: uuid.Must(uuid.NewV4()).String(), ensure: true}},
		{name: "", args: args{key: HeaderContentType, value: ContentTypeJSON, ensure: false}},
	}

	for _, item := range tests {
		w = SetResponseHeaderIf(item.args.ensure, w, item.args.key, item.args.value)
	}

	for _, item := range tests {
		if item.args.ensure {
			assert.Equal(t, item.args.value, GetResponseHeader(w, item.args.key))
		} else {
			assert.Equal(t, "", GetResponseHeader(w, item.args.key))
		}
	}
}

func TestSetResponseHeadersIf(t *testing.T) {
	headers := map[string]string{
		HeaderRequestID:   uuid.Must(uuid.NewV4()).String(),
		HeaderContentType: ContentTypeStream,
	}

	w := SetResponseHeadersIf(false, // set response header and get back response
		http.ResponseWriter(httptest.NewRecorder()), headers)

	for key, value := range headers {
		assert.NotEqual(t, value, GetResponseHeader(w, key))
	}

	w = SetResponseHeadersIf(true, // set response header and get back response
		http.ResponseWriter(httptest.NewRecorder()), headers)

	for key, value := range headers {
		assert.Equal(t, value, GetResponseHeader(w, key))
	}
}

func TestAddResponseHeader(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderContentEncoding, value: "gzip"}},
		{name: "", args: args{key: HeaderContentEncoding, value: "deflate"}},
	}

	for _, item := range tests {
		w = AddResponseHeader(w, item.args.key, item.args.value)
	}

	assert.Equal(t, "gzip", GetResponseHeader(w, HeaderContentEncoding))
	assert.ElementsMatch(t, []string{"gzip", "deflate"}, GetResponseHeaderValues(w, HeaderContentEncoding))
}

func TestAddResponseHeader2(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	type args struct {
		key   string
		value []string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderContentEncoding, value: []string{"gzip", "deflate"}}},
	}

	for _, item := range tests {
		w = AddResponseHeader(w, item.args.key, item.args.value...)
	}

	assert.Equal(t, "gzip", GetResponseHeader(w, HeaderContentEncoding))

	// Content-Encoding: gzip
	// Content-Encoding: deflate
	assert.ElementsMatch(t, []string{"gzip", "deflate"}, GetResponseHeaderValues(w, HeaderContentEncoding))
}

func TestAddResponseHeaderIf(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	type args struct {
		key    string
		value  string
		ensure bool
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{key: HeaderContentEncoding, value: "gzip", ensure: true}},
		{name: "", args: args{key: HeaderContentEncoding, value: "deflate", ensure: false}},
	}

	for _, item := range tests {
		w = AddResponseHeaderIf(item.args.ensure, w, item.args.key, item.args.value)
	}

	assert.Equal(t, "gzip", GetResponseHeader(w, HeaderContentEncoding))

	assert.Len(t, GetResponseHeaderValues(w, HeaderContentEncoding), 1)
}

func TestAddResponseHeaders(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	headers01 := map[string]string{
		HeaderContentType:     MIMEJSON,
		HeaderContentEncoding: "gzip",
	}

	headers02 := map[string]string{
		HeaderContentEncoding: "deflate",
	}

	w = AddResponseHeaders(w, headers01)
	w = AddResponseHeaders(w, headers02)

	assert.Equal(t, "gzip", GetResponseHeader(w, HeaderContentEncoding))

	// Content-Encoding: gzip
	// Content-Encoding: deflate
	assert.ElementsMatch(t, []string{"gzip", "deflate"}, GetResponseHeaderValues(w, HeaderContentEncoding))
}

func TestAddResponseHeadersIf(t *testing.T) {

	w := http.ResponseWriter(httptest.NewRecorder())

	headers01 := map[string]string{
		HeaderContentType:     MIMEJSON,
		HeaderContentEncoding: "gzip",
	}

	headers02 := map[string]string{
		HeaderContentEncoding: "deflate",
	}

	w = AddResponseHeadersIf(true, w, headers01)
	w = AddResponseHeadersIf(false, w, headers02)

	assert.Equal(t, "gzip", GetResponseHeader(w, HeaderContentEncoding))
	assert.Len(t, GetResponseHeaderValues(w, HeaderContentEncoding), 1)
}

func TestSetRequestHeaderRequestID(t *testing.T) {

	requestID := uuid.Must(uuid.NewV4()).String()
	r := SetRequestHeaderRequestID( // set header and get back req
		httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
		requestID,
	)

	assert.Equal(t, requestID, GetHeaderRequestID(r))
}

func TestSetResponseHeaderRequestID(t *testing.T) {
	w := http.ResponseWriter(httptest.NewRecorder())

	requestID := uuid.Must(uuid.NewV4()).String()
	SetResponseHeaderRequestID(w, requestID)

	assert.Equal(t, requestID, GetResponseHeader(w, HeaderRequestID))
}

func TestGetClientIP(t *testing.T) {

	localPath := "127.0.0.1:500"
	headers := map[string]string{
		HeaderForwardedFor:        localPath,
		HeaderRealIP:              localPath,
		HeaderAppEngineRemoteAddr: localPath,
	}

	for key, value := range headers {

		r := SetRequestHeader( // set header and get back req
			httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil),
			key, value,
		)

		assert.Equal(t, value, GetClientIP(r))
	}
}
