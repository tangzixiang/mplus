package context

import (
	"context"
	"time"
)

//  上下文 KEY 类型，用于在请求的 Context 中构造独立的上下文空间
type contextKey int

// Context 请求上下文，用于在请求链上传递相关信息
type Context *map[string]interface{}

const requestKey contextKey = iota

// 内部使用请求上下文键值
const (
	// ReqData 用于获取校验通过后缓存于上下文的 model 对象
	ReqData = "__req_data"

	// BodyData 用于获取校验通过后缓存于上下文的请求体内容
	BodyData = "__body_data"
)

// GetContextValue 从上下文中获取数据
func GetContextValue(ctx context.Context, key string) interface{} {
	rv := ctx.Value(requestKey)

	if rv == nil {
		return nil
	}

	return (*rv.(Context))[key]
}

// SetContextValue 在上下文中添加信息
func SetContextValue(ctx context.Context, key string, value interface{}) context.Context {
	rv := ctx.Value(requestKey)

	if rv == nil {
		return context.WithValue(ctx, requestKey, Context(&map[string]interface{}{key: value}))
	}

	(*rv.(Context))[key] = value
	return ctx
}

// NewContext 新建并初始化一个上下文
func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestKey, Context(&map[string]interface{}{}))
}

// CopyContext 在原有上下文基础上拷贝一份
func CopyContext(ctx context.Context) context.Context {
	rv := ctx.Value(requestKey)

	if rv == nil {
		return NewContext(ctx)
	}

	newContextContent := map[string]interface{}{}

	for key, value := range *rv.(Context) {
		newContextContent[key] = value
	}

	return context.WithValue(ctx, requestKey, Context(&newContextContent))
}

// GetContextValueString 获取上下文信息
func GetContextValueString(ctx context.Context, key string, defaultValue ...string) string {
	value := GetContextValue(ctx, key)
	rv, ok := value.(string)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return ""
		}
	}

	return rv
}

// GetContextValueInt 获取上下文信息
func GetContextValueInt(ctx context.Context, key string, defaultValue ...int) int {
	value := GetContextValue(ctx, key)
	rv, ok := value.(int)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return int(0)
		}
	}

	return rv
}

// GetContextValueInt8 获取上下文信息
func GetContextValueInt8(ctx context.Context, key string, defaultValue ...int8) int8 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(int8)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return int8(0)
		}
	}

	return rv
}

// GetContextValueInt16 获取上下文信息
func GetContextValueInt16(ctx context.Context, key string, defaultValue ...int16) int16 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(int16)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return int16(0)
		}
	}

	return rv
}

// GetContextValueInt32 获取上下文信息
func GetContextValueInt32(ctx context.Context, key string, defaultValue ...int32) int32 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(int32)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return int32(0)
		}
	}

	return rv
}

// GetContextValueInt64 获取上下文信息
func GetContextValueInt64(ctx context.Context, key string, defaultValue ...int64) int64 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(int64)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return int64(0)
		}
	}

	return rv
}

// GetContextValueUInt 获取上下文信息
func GetContextValueUInt(ctx context.Context, key string, defaultValue ...uint) uint {
	value := GetContextValue(ctx, key)
	rv, ok := value.(uint)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return uint(0)
		}
	}

	return rv
}

// GetContextValueUInt8 获取上下文信息
func GetContextValueUInt8(ctx context.Context, key string, defaultValue ...uint8) uint8 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(uint8)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return uint8(0)
		}
	}

	return rv
}

// GetContextValueUInt16 获取上下文信息
func GetContextValueUInt16(ctx context.Context, key string, defaultValue ...uint16) uint16 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(uint16)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return uint16(0)
		}
	}

	return rv
}

// GetContextValueUInt32 获取上下文信息
func GetContextValueUInt32(ctx context.Context, key string, defaultValue ...uint32) uint32 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(uint32)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return uint32(0)
		}
	}

	return rv
}

// GetContextValueUInt64 获取上下文信息
func GetContextValueUInt64(ctx context.Context, key string, defaultValue ...uint64) uint64 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(uint64)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return uint64(0)
		}
	}

	return rv
}

// GetContextValueBool 获取上下文信息
func GetContextValueBool(ctx context.Context, key string, defaultValue ...bool) bool {
	value := GetContextValue(ctx, key)
	rv, ok := value.(bool)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return false
		}
	}

	return rv
}

// GetContextValueBytes 获取上下文信息
func GetContextValueByte(ctx context.Context, key string, defaultValue ...byte) byte {
	value := GetContextValue(ctx, key)
	rv, ok := value.(byte)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return byte(0)
		}
	}

	return rv
}

// GetContextValueBytes 获取上下文信息
func GetContextValueBytes(ctx context.Context, key string, defaultValue ...[]byte) []byte {
	value := GetContextValue(ctx, key)
	rv, ok := value.([]byte)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return make([]byte, 0)
		}
	}

	return rv
}

// GetContextValueTime 获取上下文信息
func GetContextValueTime(ctx context.Context, key string, defaultValue ...time.Time) time.Time {
	value := GetContextValue(ctx, key)
	rv, ok := value.(time.Time)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return time.Time{}
		}
	}

	return rv
}

// GetContextValueFloat32 获取上下文信息
func GetContextValueFloat32(ctx context.Context, key string, defaultValue ...float32) float32 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(float32)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return float32(0)
		}
	}

	return rv
}

// GetContextValueFloat64 获取上下文信息
func GetContextValueFloat64(ctx context.Context, key string, defaultValue ...float64) float64 {
	value := GetContextValue(ctx, key)
	rv, ok := value.(float64)

	if len(defaultValue) > 0 {
		if value == nil || !ok {
			return defaultValue[0]
		}
	} else {
		if value == nil || !ok {
			return float64(0)
		}
	}

	return rv
}
