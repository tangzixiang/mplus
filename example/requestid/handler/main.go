package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

func main() {
	http.ListenAndServe(":8080", mplus.MRote().Use(mplus.RequestIDMiddleware).HandlerFunc(RequestID))
}

func RequestID(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	// {"request_id":"ab03f8c8-8187-4ad7-897d-97620a2d081f"} 200 OK
	pp.JSONOK(mplus.Data{"request_id": pp.RequestID()})
}