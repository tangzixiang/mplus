package mplus

import (
	"github.com/tangzixiang/mplus/testdata"
	"sync"
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

// TestLockChan 用于避免多个测试互相干扰
// 不应该在实际项目中使用
var TestLockChan = make(chan struct{}, 1)

// DefaultValidateErrorHub 测试使用
// 不应该在实际项目中使用
var (
	onceDefaultValidateErrorHub sync.Once
	DefaultValidateErrorHub     = map[ValidateErrorType]ValidateErrorFunc{}
)

func BeforeTest(setStrictJSON bool) {
	onceDefaultValidateErrorHub.Do(func() {
		for key, value := range ValidateErrorHub {
			DefaultValidateErrorHub[key] = value
		}
	})

	TestLockChan <- struct{}{}

	if setStrictJSON {
		SetStrictJSONBodyCheck(true) // 设置严格校验模式，当前模式下若无法读取 JSON 数据则会触发异常
	}
}

func AfterTest(resetHub bool) {
	if resetHub {
		ResetValidateErrorHub()
	}

	SetStrictJSONBodyCheck(false)
	<-TestLockChan
}

func ResetValidateErrorHub() {
	for key, value := range DefaultValidateErrorHub {
		RegisterValidateErrorFunc(key, value)
	}
}
