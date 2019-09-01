package mplus

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_Encode(t *testing.T) {

	urlValues := make(url.Values)
	urlValues.Add("name", "tom")
	urlValues.Add("age", "15")
	queryString := urlValues.Encode()

	if !assert.Equal(t, queryString, NewQuery().Add("name", "tom").Add("age", "15").Encode()) {
		return
	}

	if !assert.Equal(t, queryString, NewQuery().AddPairs("name", "tom", "age", "15").Encode()) {
		return
	}

	if !assert.Equal(t, queryString, NewQuery().AddByM(map[string]string{"name": "tom", "age": "15"}).Encode()) {
		return
	}

}

func TestQuery_AppendToURI(t *testing.T) {

	urlValues := make(url.Values)
	urlValues.Add("name", "tom")
	urlValues.Add("age", "15")
	path := "http://localhost?" + urlValues.Encode()

	if !assert.Equal(t, path, NewQuery().Add("name", "tom").Add("age", "15").AppendToURI("http://localhost?")) {
		return
	}

	if !assert.Equal(t, path, NewQuery().AddPairs("name", "tom", "age", "15").AppendToURI("http://localhost")) {
		return
	}
}

func TestQuery_AppendToURIf(t *testing.T) {
	urlValues := make(url.Values)
	urlValues.Add("name", "tom")
	urlValues.Add("age", "15")
	path := "http://localhost/users/1?" + urlValues.Encode()

	if !assert.Equal(t, path, NewQuery().Add("name", "tom").Add("age", "15").AppendToURIFormat("http://localhost/users/%v?", 1)) {
		return
	}

	if !assert.Equal(t, path, NewQuery().AddPairs("name", "tom", "age", "15").AppendToURIFormat("http://localhost/users/%v", 1)) {
		return
	}
}

func TestQuery_SetIf(t *testing.T) {
	values := map[string]interface{}{
		"name": nil, // do not add
		"age":  "15",
	}

	urlValues := make(url.Values)
	for key, value := range values {
		if value != nil {
			urlValues.Set(key, value.(string))
		}
	}

	path := "http://localhost/users/1?" + urlValues.Encode()
	generatePath := NewQuery().
		SetIf(values["name"] != nil, "name", "tom").
		SetIf(values["age"] != nil, "age", "15").
		AppendToURIFormat("http://localhost/users/%v?", 1)

	if !assert.Equal(t, path, generatePath) {
		return
	}
}
