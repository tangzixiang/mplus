package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

type V struct {
	Addr string `json:"addr"`
}

func main() {
	// (*V)(nil) mean that is a nil point which hold type info
	// Bind model just need type info
	http.ListenAndServe(":8080", mplus.MRote().Bind((*V)(nil)).HandlerFunc(Address))
}

func Address(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	v, ok := pp.VO().(*V)
	if !ok {
		pp.BadRequest() // 400
		return
	}

	if string(pp.ReqBodyPure()) != pp.ReqBody() {

		pp.InternalServerError() // 500
		return
	}

	if bodyMap, err := pp.ReqBodyMap(); err != nil || v.Addr != bodyMap["addr"] {

		pp.InternalServerError() // 500
		return
	}

	pp.OK() // 200
}
