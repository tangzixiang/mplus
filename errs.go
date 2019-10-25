package mplus

import (
	"github.com/tangzixiang/mplus/errs"
)

type ValidateErrorType = errs.ValidateErrorType
type ValidateError = errs.ValidateError
type ValidateErrorFunc = errs.ValidateErrorFunc

const (
	ErrBodyRead        = errs.ErrBodyRead
	ErrBodyUnmarshal   = errs.ErrBodyUnmarshal
	ErrBodyParse       = errs.ErrBodyParse
	ErrMediaTypeParse  = errs.ErrMediaTypeParse
	ErrMediaType       = errs.ErrMediaType
	ErrDecode          = errs.ErrDecode
	ErrParseQuery      = errs.ErrParseQuery
	ErrBodyValidate    = errs.ErrBodyValidate
	ErrRequestValidate = errs.ErrRequestValidate
	ErrDefault         = errs.ErrDefault
	ErrModelSelect     = errs.ErrModelSelect
)

var (
	ValidateErrorTypeMsg               = errs.ValidateErrorTypeMsg
	ValidateErrorHub                   = errs.ValidateErrorHub
	ValidateErrorWrap                  = errs.ValidateErrorWrap
	GlobalValidateErrorHandler         = errs.GlobalValidateErrorHandler
	RegisterGlobalValidateErrorHandler = errs.RegisterGlobalValidateErrorHandler
	RegisterValidateErrorFunc          = errs.RegisterValidateErrorFunc
)
