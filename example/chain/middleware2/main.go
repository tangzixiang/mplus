package main

import (
	"net/http"

	"github.com/satori/go.uuid"
	"github.com/tangzixiang/mplus"
)

func main() {
	http.ListenAndServe(":8080", mplus.MRote().Use(SetRequestID).HandlerFunc(Hello))
}

// use middleware to set requestID per request
func SetRequestID(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r).WriteRespHeader(mplus.HeaderRequestID, uuid.Must(uuid.NewV4()).String()).Handler(next).ServeHTTP()
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	mplus.PlusPlus(w, r).JSONOK(mplus.Data{"data": "hello world"})
}
