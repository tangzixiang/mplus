package mplus

import (
	"github.com/tangzixiang/mplus/validate"
)

type RequestValidate = validate.RequestValidate
type ValidateFunc = validate.ValidateFunc
type ValidateResult = validate.ValidateResult

var (
	Validate               = validate.Validate
	BindValidate           = validate.BindValidate
	DecodeTo               = validate.DecodeTo
	Parse                  = validate.Parse
	CheckValidateData      = validate.CheckValidateData
	ValidatorStandErrMsg   = validate.ValidatorStandErrMsg
	StrictJSONBodyCheck    = validate.StrictJSONBodyCheck
	SetStrictJSONBodyCheck = validate.SetStrictJSONBodyCheck
	SetStrictJSONBodyCheckLockChan = validate.SetStrictJSONBodyCheckLockChan
)