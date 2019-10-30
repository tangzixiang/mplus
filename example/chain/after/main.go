package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

var lastVisitIndex = 0

func main() {
	http.ListenAndServe(":8080", mplus.MRote().Before(MaxVisitTimesControl).After(AddNum).HandlerFunc(Num))
}

// control visit times
func MaxVisitTimesControl(w http.ResponseWriter, r *http.Request) {
	if lastVisitIndex > 10 {
		mplus.PlusPlus(w, r).Forbidden() // return status cod 403
	}
}

// increase num per request
func AddNum(w http.ResponseWriter, r *http.Request) {
	lastVisitIndex++
}

func Num(w http.ResponseWriter, r *http.Request) {
	mplus.PlusPlus(w, r).JSONOK(mplus.Data{"num": lastVisitIndex})
}
