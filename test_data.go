package mplus

import (
	"net/http"
)

// This example is centered around a form post, but doesn't have to be
// just trying to give a well rounded real life example.

// <form method="POST">
//   <input type="text" name="Name" value="joeybloggs"/>
//   <input type="text" name="Age" value="3"/>
//   <input type="text" name="Gender" value="Male"/>
//   <input type="submit"/>
// </form>

// User contains user information
type User struct {
	Name   string `validate:"required"              json:"name"`
	Age    uint8  `validate:"required,gt=0,lt=130"  json:"age"`
	Gender string `validate:"required"              json:"gender"`
	Email  string `validate:"required,email"        json:"email"`
}

// Genders 测试数据
var Genders = map[string]bool{
	"male":   true,
	"female": true,
}

// Validate 自定义对象内容检查
func (u *User) Validate(r *http.Request) (bool, string) {

	if !Genders[u.Gender] {
		return false, "Gender only be male or female"
	}

	return true, ""
}
