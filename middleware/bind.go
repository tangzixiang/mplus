package middleware

import (
	"net/http"
	"reflect"

	"github.com/pkg/errors"
	"github.com/tangzixiang/mplus/context"
	"github.com/tangzixiang/mplus/errs"
	"github.com/tangzixiang/mplus/mhttp"
	"github.com/tangzixiang/mplus/validate"
)

// Bind 通用型校验请求数据有效性的 request middleware,该方法会校验请求对象并置于请求上下文中
//
// 第一步 解析请求，读取请求数据
//
// 第二步 如果 validateData 是 ValidateFunc 类型，则调用该函数返回需要绑定数据的对象指针，执行第三步
//
// 第三步 如果 validateData 是对象指针则根据 tag 进行解析校验
//
// 第四步 校验失败则终止请求链，校验成功则将绑定数据的对象放入 request context
//
//  Bind((*VO)(nil)) or Bind(func(r *http.Request) (interface{}, error){return (*VO)(nil)})
//
func Bind(validateData interface{}) http.Handler {

	// 入参只允许指针及函数类型
	checkBindType(validateData)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var vo interface{}
		var vr validate.ValidateResult

		// 1. 解析数据
		if validate.Parse(r, &vr); vr.Err != nil {
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		if len(vr.BodyBytes) != 0 { // 放这里是因为有可能序列化失败但是 body 读取成功，可以供后续使用
			context.SetContextValue(r.Context(), context.BodyData, vr.BodyBytes)
		}

		// 得到一个全新的 VO 的对象
		vo = validate.CheckValidateData(r, validateData, &vr)
		if vr.Err != nil {
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		// 绑定数据
		if validate.DecodeTo(r, vo, &vr); vr.Err != nil {
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		// 数据校验
		if validate.BindValidate(r, vo, &vr); vr.Err != nil {
			dealValidateResultErr(w, r, vr.Err)
			return
		}

		context.SetContextValue(r.Context(), context.ReqData, vo)
	})
}

func checkBindType(validateData interface{}) {
	if _, ok := validateData.(validate.ValidateFunc); ok {
		return
	}

	if reflect.TypeOf(validateData).Kind() != reflect.Ptr {
		panic(errors.New("bind data must be object ptr or ValidateFunc"))
	}
}

func dealValidateResultErr(w http.ResponseWriter, r *http.Request, err error) {
	cErr, ok := errors.Cause(err).(errs.ValidateError)

	if !ok {
		return
	}

	// 如果存在全局解析异常处理器则优先派遣至全局
	if errs.GlobalValidateErrorHandler != nil{
		errs.GlobalValidateErrorHandler(w, r, cErr)
		return
	}

	if errHandler, exists := errs.ValidateErrorHub[cErr.Type()]; exists {
		errHandler(w, r, cErr)
		return
	}

	mhttp.BadRequest(w, r)
	return
}
