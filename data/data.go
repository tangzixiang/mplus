package data

// Data map[string]interface{} 结构的简便方式
type Data map[string]interface{}

// DynamicDataValue 回调计算 value 值的 handler,可以用于延迟计算 value
type DynamicDataValue func() interface{}

// With 根据指定内容进行初始化
func (d Data) With(m map[string]interface{}) Data {

	for key, value := range m {
		d[key] = value
	}

	return d
}

// WithNotExists 根据指定内容进行初始化，若指定键已存在则会跳过更新当前键
func (d Data) WithNotExists(m map[string]interface{}) Data {

	for key, value := range m {
		if _, e := d[key]; !e {
			d[key] = value
		}
	}

	return d
}

// Push 添加指定键值对，已存在的会被覆盖
func (d Data) Push(key string, value interface{}) Data {
	d[key] = value
	return d
}

// PushD 动态添加指定键值对，已存在的会被覆盖
func (d Data) PushD(key string, value DynamicDataValue) Data {
	d[key] = value()
	return d
}

// PushIf 添加指定键值对，已存在的会被覆盖，仅在 ensure 为 true 时执行
func (d Data) PushIf(ensure bool, key string, value interface{}) Data {
	if ensure {
		d[key] = value
	}
	return d
}

// PushIfD 动态添加指定键值对，已存在的会被覆盖，仅在 ensure 为 true 时执行
func (d Data) PushIfD(ensure bool, key string, value DynamicDataValue) Data {
	if ensure {
		d[key] = value()
	}
	return d
}

// Push 添加多对键值对，已存在的会被覆盖
//
// For example:
//  - d.PushPairs(key1,value1,key2,value2,...)
func (d Data) PushPairs(key string, value interface{}, pairs ... interface{}) Data {
	pairs = append([]interface{}{key, value}, pairs...)

	if len(pairs)%2 != 0 {
		return d
	}

	for i := 0; i < len(pairs); i += 2 {
		if key, ok := pairs[i].(string); ok {
			d[key] = pairs[i+1]
		}
	}

	return d
}

// Push 添加多对键值对，已存在的会被覆盖，仅在 ensure 为 true 时执行
//
// For example:
//  - d.PushPairsIf(GetTrue(),key1,value1,key2,value2,...)
func (d Data) PushPairsIf(ensure bool, key string, value interface{}, pairs ... interface{}) Data {
	if !ensure {
		return d
	}

	pairs = append([]interface{}{key, value}, pairs...)

	if len(pairs)%2 != 0 {
		return d
	}

	for i := 0; i < len(pairs); i += 2 {
		if key, ok := pairs[i].(string); ok {
			d[key] = pairs[i+1]
		}
	}

	return d
}

// PushNotExists 添加指定键值对，若指定键已存在则会跳过更新当前键
func (d Data) PushNotExists(key string, value interface{}) Data {
	if _, e := d[key]; !e {
		d[key] = value
	}
	return d
}

// PushD 动态添加指定键值对，若指定键已存在则会跳过更新当前键
func (d Data) PushNotExistsD(key string, value DynamicDataValue) Data {
	if _, e := d[key]; !e {
		d[key] = value()
	}
	return d
}

// PushIf 添加指定键值对，若指定键已存在则会跳过更新当前键，仅在 ensure 为 true 时执行
func (d Data) PushNotExistsIf(ensure bool, key string, value interface{}) Data {
	if ensure {
		if _, e := d[key]; !e {
			d[key] = value
		}
	}
	return d
}

// PushNotExistsIfD 动态添加指定键值对，若指定键已存在则会跳过更新当前键，仅在 ensure 为 true 时执行
func (d Data) PushNotExistsIfD(ensure bool, key string, value DynamicDataValue) Data {
	if ensure {
		if _, e := d[key]; !e {
			d[key] = value()
		}
	}
	return d
}

// PushPairsNotExists 添加多对键值对，若指定键已存在则会跳过更新当前键
//
// For example:
//  - d.PushPairsNotExists(key1,value1,key2,value2,...)
func (d Data) PushPairsNotExists(key string, value interface{}, pairs ... interface{}) Data {
	pairs = append([]interface{}{key, value}, pairs...)

	if len(pairs)%2 != 0 {
		return d
	}

	for i := 0; i < len(pairs); i += 2 {
		if key, ok := pairs[i].(string); ok {
			if _, e := d[key]; !e {
				d[key] = pairs[i+1]
			}
		}
	}

	return d
}

// PushPairsNotExistsIf 添加多对键值对，若指定键已存在则会跳过更新当前键，仅在 ensure 为 true 时执行
//
// For example:
//  - d.PushPairsNotExistsIf(GetTrue(),key1,value1,key2,value2,...)
func (d Data) PushPairsNotExistsIf(ensure bool, key string, value interface{}, pairs ... interface{}) Data {
	if !ensure {
		return d
	}

	pairs = append([]interface{}{key, value}, pairs...)

	if len(pairs)%2 != 0 {
		return d
	}

	for i := 0; i < len(pairs); i += 2 {
		if key, ok := pairs[i].(string); ok {
			if _, e := d[key]; !e {
				d[key] = pairs[i+1]
			}
		}
	}

	return d
}

// Del 删除指定键
func (d Data) Del(key string) Data {
	delete(d, key)
	return d
}

// DelAll 删除多个键
func (d Data) DelAll(keys ... string) Data {
	for _, key := range keys {
		delete(d, key)
	}
	return d
}

// Exists 判断指定键是否存在
func (d Data) Exists(key string) bool {
	_, e := d[key]
	return e
}

// ForEach 遍历 Data
func (d Data) ForEach(f func(key string, value interface{}, data Data)) {
	for key, value := range d {
		f(key, value, d)
	}
}

// Keys 获取 Data 的所有 key
func (d Data) Keys() (keys []string) {
	for key := range d {
		keys = append(keys, key)
	}
	return
}

// Len 获取 Data 的长度
func (d Data) Len()int{
	return len(d)
}