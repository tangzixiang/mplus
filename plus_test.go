package mplus

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPP_ReqBody(t *testing.T) {
	content := NewQuery().AddPairs("name", "tom", "age", "18").Encode()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(content))
	resp := httptest.NewRecorder()
	resp.Body = &bytes.Buffer{}
	resp.Body.WriteString(content)

	pp := PlusPlus(resp, req)

	assert.Equal(t, content, pp.ReqBody())

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	assert.Nil(t, err)
	assert.Equal(t, content, string(bodyBytes))
}

func TestPP_ReqBodyMap(t *testing.T) {
	contentBytes := []byte(`{"name":"tom","age":18}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewReader(contentBytes))
	resp := httptest.NewRecorder()

	m := map[string]interface{}{}
	assert.Nil(t, json.Unmarshal(contentBytes, &m))

	pp := PlusPlus(resp, req)
	bodyM, err := pp.ReqBodyMap()
	assert.Nil(t, err)

	for key, value := range m {
		assert.Equal(t, value, bodyM[key])
	}

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	assert.Nil(t, err)

	assert.Equal(t, contentBytes, bodyBytes)
}

type TestUnmarshaler struct {
	m map[string]interface{}
}

func (t *TestUnmarshaler) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &t.m)
}

func TestPP_ReqBodyToUnmarshaler(t *testing.T) {
	contentBytes := []byte(`{"name":"tom","age":18}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewReader(contentBytes))
	resp := httptest.NewRecorder()

	m := map[string]interface{}{}
	assert.Nil(t, json.Unmarshal(contentBytes, &m))

	pp := PlusPlus(resp, req)
	mer := &TestUnmarshaler{m: map[string]interface{}{}}
	assert.Nil(t, pp.ReqBodyToUnmarshaler(mer))

	for key, value := range m {
		assert.Equal(t, value, mer.m[key])
	}

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	assert.Nil(t, err)

	assert.Equal(t, contentBytes, bodyBytes)
}
