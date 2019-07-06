package mplus

import (
	"net/http"
	"net/url"
)

// CopyRequest 拷贝一个请求，body 不会被拷贝，因为 body 是一个数据流
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

	// cope header
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

	return r2.WithContext(CopyContext(r.Context()))
}
