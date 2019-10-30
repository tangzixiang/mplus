package main

import (
	"net/http"
	"strings"

	"github.com/tangzixiang/mplus"
)

type V struct {
	Addr string `json:"addr" validate:"min=10"` // min len is 10
}

// implement mplus.RequestValidate
func (v V) Validate(r *http.Request) (ok bool /*校验是否成功*/, errMsg string /*校验失败的原因*/) {

	if strings.Index(v.Addr, "广东") != 0 {
		return false, "addr must begin 广东"
	}

	return true, ""
}

func main() {
	// register a hook to show err message when validate failed by model.Validate
	mplus.RegisterValidateErrorFunc(mplus.ErrRequestValidate, func(w http.ResponseWriter, r *http.Request, err error) {
		mplus.PlusPlus(w, r).JSON(map[string]string{"err_message": err.Error()}, 400)
	})

	http.ListenAndServe(":8080", mplus.MRote().Bind((*V)(nil)).HandlerFunc(Address))
}

func Address(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	// data pass to response is V instance from request data
	pp.JSONOK(pp.VO().(*V))
}
