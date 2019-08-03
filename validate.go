package mplus

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"gopkg.in/go-playground/validator.v9"
)

// RequestValidate 实现该接口的不同VO(请求结构体)可以自校验
type RequestValidate interface {
	Validate(r *http.Request) (ok bool /*校验是否成功*/, errMsg string /*校验失败的原因*/)
}

// ValidateFunc 自定义当前请求需要用到的 VO 对象，用于
// 返回的 VO 不应该为 nil，若无法返回正确的 VO 应该在返回的 error 中进行说明
type ValidateFunc func(r *http.Request) (interface{}, error)

// ValidateResult 请求校验结果
type ValidateResult struct {
	Err         error
	MediaType   string
	BodyBytes   []byte
	BodyValues  url.Values
	QueryValues url.Values
}

// 校验器
var (
	Validate = validator.New()
)

// ParseValidate 解析请求并将数据注入到指定对象，返回解析结果
func ParseValidate(r *http.Request, obj interface{}) ValidateResult {
	var vr ValidateResult

	// 1. 解析
	if parse(r, obj, &vr); vr.Err != nil {
		return vr
	}

	// 2. tag 规则校验
	if err := Validate.Struct(obj); err != nil {
		vr.Err = ValidateErrorWrap(err, ErrBodyValidate)
		return vr
	}

	// 3. 自定义校验
	if v, hasValidateMethod := obj.(RequestValidate); hasValidateMethod {
		if ok, errMsg := v.Validate(r); !ok {
			vr.Err = ValidateErrorWrap(errors.New(errMsg), ErrRequestValidate)
			return vr
		}
	}

	return vr
}

func parse(r *http.Request, obj interface{}, vr *ValidateResult) {

	var err error

	// 查看请求 mime 类型
	vr.MediaType, err = ParseMediaType(r)
	if err != nil {
		vr.Err = ValidateErrorWrap(err, ErrMediaTypeParse)
		return
	}

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete /*delete 请求可以有主体 https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/DELETE */ : // 考虑做成动态的
		parsePost(r, obj, vr)
	default:
		// GET DELETE HEAD OPTION
	}

	// catch error
	if vr.Err != nil {
		return
	}

	if r.PostForm != nil {
		vr.BodyValues = r.PostForm
	} else {
		vr.BodyValues = url.Values{}
	}

	if r.Form != nil {
		vr.QueryValues = r.Form
	} else {
		vr.QueryValues = url.Values{}
	}

	// parse url query
	parseQueryDecode(r, obj, vr)
}

func parsePost(r *http.Request, obj interface{}, vr *ValidateResult) {
	switch vr.MediaType {
	case MIMEPOSTForm:
		if err := r.ParseForm(); err != nil {
			vr.Err = ValidateErrorWrap(err, ErrBodyParse)
		}
	case MIMEMultipartPOSTForm:
		if err := r.ParseMultipartForm(defaultMemory); err != nil {
			vr.Err = ValidateErrorWrap(err, ErrBodyParse)
		}
	case MIMEJSON:
		body := DumpRequestPure(r)
		if len(body) != 0 {
			vr.BodyBytes = body
			if err := json.Unmarshal(body, obj); err != nil {
				vr.Err = ValidateErrorWrap(err, ErrBodyUnmarshal)
			}
		} else {
			vr.Err = ValidateErrorWrap(errors.New("body empty"), ErrBodyRead)
		}
	default:
		vr.Err = ValidateErrorWrap(errors.New("mediaType not support"), ErrMediaType)
	}
}

// 解析 URL ,并将 URL 参数解析到指定对象
func parseQueryDecode(r *http.Request, obj interface{}, vr *ValidateResult) {
	parseQuery, err := ParseQuery(r)

	if err != nil {
		vr.Err = ValidateErrorWrap(err, ErrParseQuery)
		return
	}

	if parseQuery != nil && len(parseQuery) > 0 {
		vr.QueryValues = NewQueryWith(vr.QueryValues).With(parseQuery).v
	}

	if err := DecodeForm(obj, vr.QueryValues); err != nil {
		vr.Err = ValidateErrorWrap(err, ErrDecode)
		return
	}
}

const (
	sep       = ";"
	quo       = "'"
	failedMsg = " failed on tag "
)

// ValidatorStandErrMsg 构建请求错误提示信息
func ValidatorStandErrMsg(err error) string {

	vErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	buildRv := ""
	for i, e := range vErr {
		if i != 0 {
			buildRv += sep
		}
		buildRv += e.Namespace() + failedMsg + quo + e.Tag() + quo
	}

	return buildRv
}
