package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
	"github.com/tangzixiang/mplus/message"
)

func main() {
	mux := http.NewServeMux()

	// register a StatusBadRequest hook to change default's behavior
	mplus.RegisterHttpStatusMethod(http.StatusBadRequest, func(w http.ResponseWriter, r *http.Request, m message.Message, statusCode int) {

		// {"err_message":"Bad Request"} 400 Bad Request
		mplus.PlusPlus(w, r).JSON(mplus.Data{"err_message": m.En()}, m.Status())
	})

	//  response's status code is 400
	mux.Handle("/400", mplus.MRote().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mplus.PlusPlus(w, r).BadRequest()
	}))

	http.ListenAndServe(":8080", mux)
}
