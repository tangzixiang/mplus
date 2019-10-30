package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

func main() {
	http.ListenAndServe(":8080", mplus.MRote().HandlerFunc(Hello))
}

func Hello(w http.ResponseWriter, r *http.Request) {

	// take you w,r then give you a plus
	mplus.PlusPlus(w, r).JSONOK(mplus.Data{"data": "hello world"})
}
