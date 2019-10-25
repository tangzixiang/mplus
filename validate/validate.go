package validate

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"

	"github.com/pkg/errors"

	"github.com/tangzixiang/mplus/decode"
	"github.com/tangzixiang/mplus/errs"
	"github.com/tangzixiang/mplus/mhttp"
	"github.com/tangzixiang/mplus/mime"
	"github.com/tangzixiang/mplus/query"
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

// bindValidate 解析请求并将数据注入到指定对象，返回解析结果
func BindValidate(r *http.Request, obj interface{}, vr *ValidateResult) {
	bindValidate(r, obj, vr)
}

// Parse
func Parse(r *http.Request, vr *ValidateResult) {
	parse(r, vr)
}

// DecodeTo
func DecodeTo(r *http.Request, obj interface{}, vr *ValidateResult) {
	decodeTo(r, obj, vr)
}

// CheckValidateData 校验及获取一个用于绑定的 model 对象
func CheckValidateData(r *http.Request, validateData interface{}, vr *ValidateResult) (interface{}) {
	return checkValidateData(r, validateData, vr)
}

func bindValidate(r *http.Request, obj interface{}, vr *ValidateResult) {

	// 2. tag 规则校验
	if err := Validate.Struct(obj); err != nil {
		vr.Err = errs.ValidateErrorWrap(err, errs.ErrBodyValidate)
		return
	}

	// 3. 自定义校验
	if v, hasValidateMethod := obj.(RequestValidate); hasValidateMethod {
		if ok, errMsg := v.Validate(r); !ok {
			vr.Err = errs.ValidateErrorWrap(errors.New(errMsg), errs.ErrRequestValidate)
			return
		}
	}

	return
}

func parse(r *http.Request, vr *ValidateResult) {

	var err error

	// 查看请求 mime 类型
	vr.MediaType, err = mime.ParseRequestMediaType(r)
	if err != nil {
		vr.Err = errs.ValidateErrorWrap(err, errs.ErrMediaTypeParse)
		return
	}

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete /*delete 请求可以有主体 https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE */ : // 考虑做成动态的
		switch vr.MediaType {
		case mime.MIMEPOSTForm:
			if err := r.ParseForm(); err != nil { // 支持 POST PUT PATCH 含有主体
				vr.Err = errs.ValidateErrorWrap(err, errs.ErrBodyParse)
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

		case mime.MIMEMultipartPOSTForm:
			if err := r.ParseMultipartForm(mhttp.DefaultMemorySize()); err != nil {
				vr.Err = errs.ValidateErrorWrap(err, errs.ErrBodyParse)
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

		case mime.MIMEJSON:
			body := mhttp.DumpRequestPure(r)
			if len(body) != 0 {
				vr.BodyBytes = body
			} else if StrictJSONBodyCheck() { // 是否严格校验 json body
				vr.Err = errs.ValidateErrorWrap(errors.New("body empty"), errs.ErrBodyRead)
			}

			// catch error
			if vr.Err != nil {
				return
			}
		default:
			vr.Err = errs.ValidateErrorWrap(errors.New("mediaType not support"), errs.ErrMediaType)
		}
	default:
		// GET HEAD OPTION
		// parse url query
		parseQuery(r, vr)
	}
}

func decodeTo(r *http.Request, obj interface{}, vr *ValidateResult) {

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete /*delete 请求可以有主体 https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/DELETE */ : // 考虑做成动态的
		switch vr.MediaType {
		case mime.MIMEPOSTForm:
			fallthrough
		case mime.MIMEMultipartPOSTForm:
			if err := decode.DecodeForm(obj, vr.QueryValues, vr.BodyValues); err != nil {
				vr.Err = errs.ValidateErrorWrap(err, errs.ErrDecode)
				return
			}
		case mime.MIMEJSON:
			if len(vr.BodyBytes) > 0 {
				if err := json.Unmarshal(vr.BodyBytes, obj); err != nil {
					vr.Err = errs.ValidateErrorWrap(err, errs.ErrBodyUnmarshal)
				}
			}
			// 考虑到性能问题 json 请求默认不解析 query string 到对象内，后续需要可以通过 mplus.PP.GetQuery 获取
		}
	default:
		// GET HEAD OPTION
		if err := decode.DecodeForm(obj, vr.QueryValues); err != nil {
			vr.Err = errs.ValidateErrorWrap(err, errs.ErrDecode)
			return
		}
	}
}

// 解析 URL ,并将 URL 参数解析到指定对象
func parseQuery(r *http.Request, vr *ValidateResult) {
	parseQuery, err := query.ParseQuery(r)

	if err != nil {
		vr.Err = errs.ValidateErrorWrap(err, errs.ErrParseQuery)
		return
	}

	if parseQuery != nil && len(parseQuery) > 0 {
		vr.QueryValues = query.NewWith(vr.QueryValues).With(parseQuery).Values()
	}
}

func checkValidateData(r *http.Request, validateData interface{}, vr *ValidateResult) (interface{}) {

	// 如果是函数则获取函数返回的值,返回值只能是指针
	if voTypeFunc, ok := validateData.(ValidateFunc); ok {
		var err error
		if validateData, err = voTypeFunc(r); err != nil {
			vr.Err = errs.ValidateErrorWrap(err, errs.ErrModelSelect)
			return nil
		}

		if reflect.TypeOf(validateData).Kind() != reflect.Ptr {
			vr.Err = errs.ValidateErrorWrap(err, errs.ErrModelSelectType)
			return nil
		}
	}

	return reflect.New( // new vo
		reflect.TypeOf(validateData). // get type ptr
			Elem(), // get type
	).Interface()
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
