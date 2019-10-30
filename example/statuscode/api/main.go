package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/tangzixiang/mplus"
)

func main() {
	mux := http.NewServeMux()
	mr := mplus.MRote() // It's very easy to reuse

	//  response's status code is 200
	mux.Handle("/200", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r)
	}))

	//  response's status code is 400
	mux.Handle("/400", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r).BadRequest()
	}))

	//  response's status code is 401
	mux.Handle("/401", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// WWW-Authenticate: Basic
		// see https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401
		mplus.PlusPlus(w, r).WriteRespHeader(mplus.HeaderWWWAuthenticate, `Basic`).Unauthorized()
	}))

	//  response's status code is 403
	mux.Handle("/403", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// chang default Date header
		// Date: Sat, 28 Sep 2019 01:58:43 GMT

		// TimeFormat is the time format to use when generating times in HTTP
		// headers. It is like time.RFC1123 but hard-codes GMT as the time
		// zone. The time being formatted must be in UTC for Format to
		// generate the correct format.
		mplus.PlusPlus(w, r).WriteRespHeader(mplus.HeaderDate, time.Now().Add(time.Hour).UTC().Format(http.TimeFormat)).Forbidden()
	}))

	//  response's status code is 404
	mux.Handle("/404", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r).NotFound()
	}))

	//  response's status code is 405
	mux.Handle("/405", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Allow: GET, POST, HEAD
		// see https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
		mplus.PlusPlus(w, r).WriteRespHeader(
			mplus.HeaderAllow, strings.Join([]string{http.MethodGet, http.MethodPost, http.MethodPut}, mplus.SplitSepComma)).NotAllowed()
	}))

	//  response's status code is 408
	mux.Handle("/408", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Connection: close
		// see https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408
		mplus.PlusPlus(w, r).WriteRespHeader(mplus.HeaderConnection, "close").RequestTimeout()
	}))

	//  response's status code is 409
	mux.Handle("/409", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r).Conflict()
	}))

	//  response's status code is 415
	mux.Handle("/415", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r).UnsupportedMediaType()
	}))

	//  response's status code is 500
	mux.Handle("/500", mr.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r).InternalServerError()
	}))

	http.ListenAndServe(":8080", mux)
}
