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
