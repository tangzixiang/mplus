package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

type V struct {
	Addr string `json:"addr" validate:"min=10"` // min len is 10
}

// errCode 400001
var ErrCodeAddrNotExists = 400001

// ErrCodeCallbackFun will be perform when ErrCodeAddrNotExists be use to PP.CallbackByCode
func ErrCodeCallbackFun(w http.ResponseWriter, r *http.Request, m mplus.Message, respData interface{}) {

	// {"code":400001,"message":"addr not exists"}
	mplus.JSON(w, r, mplus.Data{"message": m.En(), "code": m.ErrCode()}, m.Status())
}

func CheckAddr(string) bool {
	return false
}

func main() {

	// register a message for ErrCodeAddrNotExists
	mplus.Messages.Add(mplus.NewCallbackMessage(http.StatusBadRequest, ErrCodeAddrNotExists, "addr not exists", ErrCodeCallbackFun))

	http.ListenAndServe(":8080", mplus.MRote().Bind((*V)(nil)).HandlerFunc(Address))
}

// Handler
func Address(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	if !CheckAddr(pp.VO().(*V).Addr) {

		//  pp search ErrCodeAddrNotExists message  pass to ErrCodeCallbackFun
		// second arg is respData pass to registered callback func on Messages -> ErrCodeCallbackFun
		pp.CallbackByCode(ErrCodeAddrNotExists, nil)
		return
	}

	pp.JSONOK(nil) // pass nil will specify mplus.EmptyRespData
}
