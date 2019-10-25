package mplus

import (
	"github.com/tangzixiang/mplus/context"
)

// 内部使用请求上下文键值
const (
	ReqData  = context.ReqData
	BodyData = context.BodyData
)

var (
	GetContextValue        = context.GetContextValue
	SetContextValue        = context.SetContextValue
	NewContext             = context.NewContext
	CopyContext            = context.CopyContext
	GetContextValueString  = context.GetContextValueString
	GetContextValueInt     = context.GetContextValueInt
	GetContextValueInt8    = context.GetContextValueInt8
	GetContextValueInt16   = context.GetContextValueInt16
	GetContextValueInt32   = context.GetContextValueInt32
	GetContextValueInt64   = context.GetContextValueInt64
	GetContextValueUInt    = context.GetContextValueUInt
	GetContextValueUInt8   = context.GetContextValueUInt8
	GetContextValueUInt16  = context.GetContextValueUInt16
	GetContextValueUInt32  = context.GetContextValueUInt32
	GetContextValueUInt64  = context.GetContextValueUInt64
	GetContextValueBool    = context.GetContextValueBool
	GetContextValueByte    = context.GetContextValueByte
	GetContextValueBytes   = context.GetContextValueBytes
	GetContextValueTime    = context.GetContextValueTime
	GetContextValueFloat32 = context.GetContextValueFloat32
	GetContextValueFloat64 = context.GetContextValueFloat64
)
