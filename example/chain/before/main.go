package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

var num = 0

func main() {
	http.ListenAndServe(":8080", mplus.MRote().Before(AddNum).HandlerFunc(Num))
}

// increase num per request
func AddNum(w http.ResponseWriter, r *http.Request) {
	num ++
}

func Num(w http.ResponseWriter, r *http.Request) {
	mplus.PlusPlus(w, r).JSONOK(mplus.Data{"num": num})
}
