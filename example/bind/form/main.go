package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

type V struct {
	Addr string `json:"addr" form:"addr" validate:"min=10"` // min len is 10
}

func main() {
	// (*V)(nil) mean that is a nil point which hold type info
	// Bind model just need type info
	http.ListenAndServe(":8080", mplus.MRote().Bind((*V)(nil)).HandlerFunc(Address))
}

func Address(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	// data pass to response is V instance from request data
	pp.JSONOK(pp.VO().(*V))
}
