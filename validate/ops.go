package validate


// strictJSONBodyCheck 是否严格校验 json 请求，若 StrictJSONBodyCheck 值为 true 且请求为 json 请求，则读取到空的数据将抛出异常
var strictJSONBodyCheck = false

// StrictJSONBodyCheck 获取 strictJSONBodyCheck
// 若 strictJSONBodyCheck 值为 true 且请求为 json 请求，则读取到空的数据将抛出异常
func StrictJSONBodyCheck()bool{
	return strictJSONBodyCheck
}

// SetStrictJSONBodyCheck 设置 strictJSONBodyCheck
// 若 strictJSONBodyCheck 值为 true 且请求为 json 请求，则读取到空的数据将抛出异常
func SetStrictJSONBodyCheck(b bool){
	strictJSONBodyCheck = b
}

// SetStrictJSONBodyCheckLockChan 由于测试存在多次调用及还原的情况，为了避免互相干扰，引入锁的机制
// 不应该在实际项目中使用
var SetStrictJSONBodyCheckLockChan = make(chan struct{}, 1)