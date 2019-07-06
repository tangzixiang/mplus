package mplus

import (
	"net/http"
	"net/url"
)

// Queries 获取 URL 上的请求字段
func Queries(r *http.Request) url.Values {

	values,err := ParseQuery(r)
	if err != nil {
		return make(url.Values)
	}

	return values
}
