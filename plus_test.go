package mplus

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPP_ReqBody(t *testing.T) {

	content := NewQuery().AddPairs("name", "tom", "age", "18").Encode()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(content))
	resp := httptest.NewRecorder()
	resp.Body = &bytes.Buffer{}
	resp.Body.WriteString(content)

	pp := PlusPlus(resp, req)

	if !assert.Equal(t, content, pp.ReqBody()) {
		return
	}

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	if !assert.Nil(t, err) {
		return
	}

	if !assert.Equal(t, content, string(bodyBytes)) {
		return
	}
}

func TestPP_ReqBodyMap(t *testing.T) {

	contentBytes := []byte(`{"name":"tom","age":18}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewReader(contentBytes))
	resp := httptest.NewRecorder()

	m := map[string]interface{}{}
	if !assert.Nil(t, json.Unmarshal(contentBytes, &m)) {
		return
	}

	pp := PlusPlus(resp, req)
	bodyM, err := pp.ReqBodyMap()
	if !assert.Nil(t, err) {
		return
	}

	for key, value := range m {
		if !assert.Equal(t, value, bodyM[key]) {
			return
		}
	}

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	if !assert.Nil(t, err) {
		return
	}

	if !assert.Equal(t, contentBytes, bodyBytes) {
		return
	}
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
	if !assert.Nil(t, json.Unmarshal(contentBytes, &m)) {
		return
	}

	pp := PlusPlus(resp, req)
	mer := &TestUnmarshaler{m: map[string]interface{}{}}
	if !assert.Nil(t, pp.ReqBodyToUnmarshaler(mer)) {
		return
	}

	for key, value := range m {
		if !assert.Equal(t, value, mer.m[key]) {
			return
		}
	}

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	if !assert.Nil(t, err) {
		return
	}

	if !assert.Equal(t, contentBytes, bodyBytes) {
		return
	}

}
