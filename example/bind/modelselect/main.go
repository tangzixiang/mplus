package main

import (
	"errors"
	"net/http"

	"github.com/tangzixiang/mplus"
)

// InsertUser use for insert a new one
type InsertUser struct {
	Account  string `json:"account"     form:"account"`
	Name     string `json:"name"        form:"name"`
	Password string `json:"password"    form:"password"`
}

//  UpdateUser use for update name
type UpdateUser struct {
	Account string `json:"account"  form:"account"`
	Name    string `json:"name"     form:"name"`
}

// SelectBindModel will calculate which type of model to bind when handler got a request
func SelectBindModel(r *http.Request) (interface{}, error) {

	switch mplus.GetHeader(r, "X-Do-What") {
	case "update":
		return (*UpdateUser)(nil), nil
	case "insert":
		return (*InsertUser)(nil), nil
	}

	return nil, errors.New("bind type not found")
}

func main() {
	// register a hook to show err message when select model failed
	mplus.RegisterValidateErrorFunc(mplus.ErrModelSelect, func(w http.ResponseWriter, r *http.Request, err error) {

		// {"err_message":"bind type not found"} 400
		mplus.PlusPlus(w, r).JSON(mplus.Data{"err_message": err.Error()}, 400)
	})

	handler := mplus.MRote().Bind(mplus.ValidateFunc(SelectBindModel)).HandlerFunc(Whatever)
	http.ListenAndServe(":8080",handler)
}

// Handler
func Whatever(w http.ResponseWriter, r *http.Request) {
	pp := mplus.PlusPlus(w, r)

	// got pp.VO() do anything

	// pass data to response
	pp.JSONOK(pp.VO())
}
