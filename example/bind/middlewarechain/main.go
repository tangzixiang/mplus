package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

type V struct {
	Addr string `json:"addr" validate:"min=10"` // min len is 10
}

func main() {
	http.ListenAndServe(":8080",
		mplus.Use(mplus.PreMiddleware)( // use middleware Pre to init Context
			mplus.ThunkHandler( // compress handlers
				mplus.Bind((*V)(nil)),       // add pre handler -> Bind
				http.HandlerFunc(Address))), // last handler is target httpHandler
	)
}

func Address(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	pp.JSON(pp.VO(), 200)
}
