package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

type V struct {
	Addr string `json:"addr" validate:"min=10"` // min len is 10
}

func main() {
	// register a hook to show err message when validate failed
	mplus.RegisterValidateErrorFunc(mplus.ErrBodyValidate, func(w http.ResponseWriter, r *http.Request, err error) {

		// {"err_message":"Key: 'V.Addr' Error:Field validation for 'Addr' failed on the 'min' tag"} 400
		mplus.PlusPlus(w, r).JSON(mplus.Data{"err_message": err.Error()}, 400)
	})

	http.ListenAndServe(":8080", mplus.MRote().Bind((*V)(nil)).HandlerFunc(Address))
}

func Address(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	// data pass to response is V instance from request data
	pp.JSONOK(pp.VO().(*V))
}
