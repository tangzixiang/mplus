package mplus

import (
	"github.com/tangzixiang/mplus/testdata"
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
type User testdata.User

// Genders 测试数据
var Genders = testdata.Genders
