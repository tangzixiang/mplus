package mhttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tangzixiang/mplus/context"
)

// CopyRequest 拷贝一个请求
func CopyRequest(r *http.Request) *http.Request {

	r2 := &http.Request{
		Proto:            r.Proto,
		ProtoMajor:       r.ProtoMajor,
		ProtoMinor:       r.ProtoMinor,
		Method:           r.Method,
		ContentLength:    r.ContentLength,
		TransferEncoding: r.TransferEncoding,
		Host:             r.Host,
		RemoteAddr:       r.RemoteAddr,
		RequestURI:       r.RequestURI,
	}

	// copy header
	if r.Header != nil {
		r2.Header = http.Header{}

		for key, values := range r.Header {
			for _, value := range values {
				r2.Header.Add(key, value)
			}
		}
	}

	// copy form
	if r.Form != nil {
		r2.Form = url.Values{}

		for key, values := range r.Form {
			for _, value := range values {
				r2.Form.Add(key, value)
			}
		}
	}

	// copy post form
	if r.PostForm != nil {
		r2.PostForm = url.Values{}

		for key, values := range r.PostForm {
			for _, value := range values {
				r2.PostForm.Add(key, value)
			}
		}
	}

	if r.MultipartForm != nil {
		newForm := *r.MultipartForm

		r2.MultipartForm = &newForm
	}

	if r.URL != nil {
		newURL := *r.URL
		r2.URL = &newURL
	}

	body, _ := ioutil.ReadAll(r.Body)

	// Reset req.Body so it can be use again
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	r2.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return r2.WithContext(context.CopyContext(r.Context()))
}
