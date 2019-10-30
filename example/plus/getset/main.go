package main

import (
	"net/http"

	"github.com/satori/go.uuid"
	"github.com/tangzixiang/mplus"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	http.ListenAndServe(":8080", mplus.MRote().Use(SetRequestID, PreSearchUser).HandlerFunc(Hello))
}

// use middleware to set requestID per request
func SetRequestID(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		pp := mplus.PlusPlus(w, r).Handler(next)

		// add "X-Request-Id" to request's context and header at the same time
		pp.WriteRespHeader(mplus.HeaderRequestID, pp.SetStringR(mplus.HeaderRequestID, uuid.Must(uuid.NewV4()).String()))

		// call next
		pp.ServeHTTP()
	}
}

const KeyUser = "user"

// use middleware to check user before handler
func PreSearchUser(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		pp := mplus.PlusPlus(w, r).Handler(next)

		// search user by id which get from url then add to request's context
		pp.Set(KeyUser, SearchUserService(pp.Query("id")))

		// call next
		pp.ServeHTTP()
	}
}

func SearchUserService(id string) *User {
	return &User{"tom"}
}

func Hello(w http.ResponseWriter, r *http.Request) {

	pp := mplus.PlusPlus(w, r)

	requestID := pp.Get(mplus.HeaderRequestID) // get requestID from request's context
	userName := pp.Get(KeyUser).(*User).Name   // get user from request's context

	pp.JSONOK(mplus.Data{"request_id": requestID, "message": "hello " + userName})
}
