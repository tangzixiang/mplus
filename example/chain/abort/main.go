package main

import (
	"net/http"

	"github.com/tangzixiang/mplus"
)

func main() {
	http.ListenAndServe(":8080", mplus.MRote().Use(mid1, mid2, mid3).Before(before1, before2).After(after1, after2).HandlerFunc(Hello))
}

var orders []int

func before1(w http.ResponseWriter, r *http.Request) {
	orders = append(orders, 7)

	// 内置 Abort
	mplus.PlusPlus(w, r).JSONOK(mplus.Data{"order": orders})
	orders = nil
}

func before2(w http.ResponseWriter, r *http.Request) { // 不会执行
	orders = append(orders, 8)
}

func after1(w http.ResponseWriter, r *http.Request) { // 不会执行
	orders = append(orders, 10)
}

func after2(w http.ResponseWriter, r *http.Request) { // 不会执行

	mplus.PlusPlus(w, r).JSONOK(mplus.Data{"order": append(orders, 11)})
	orders = nil
}

func mid1(next http.HandlerFunc) http.HandlerFunc {

	orders = append(orders, 3)
	return func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 4)
		// call next
		next.ServeHTTP(w, r)
	}
}

func mid2(next http.HandlerFunc) http.HandlerFunc {

	orders = append(orders, 2)
	return func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 5)
		// call next
		next.ServeHTTP(w, r)
	}
}

func mid3(next http.HandlerFunc) http.HandlerFunc {

	orders = append(orders, 1)
	return func(w http.ResponseWriter, r *http.Request) {
		orders = append(orders, 6)
		// call next
		next.ServeHTTP(w, r)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) { // 不会执行
	orders = append(orders, 9)
}
