package mplus

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// ParseQuery 解析 URL 上的参数
func ParseQuery(r *http.Request) (url.Values, error) {
	var newValues url.Values

	if r.URL == nil {
		return nil,errors.New("url is empty")
	}

	newValues = r.URL.Query()
	if newValues != nil {
		return newValues,nil
	}

	return  url.ParseQuery(r.URL.RawQuery)
}
