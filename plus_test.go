package mplus

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	"github.com/tangzixiang/mplus/message"
)

func TestPP_ReqBody(t *testing.T) {

	content := NewQuery().AddPairs("name", "tom", "age", "18").Encode()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(content))
	resp := httptest.NewRecorder()
	resp.Body = &bytes.Buffer{}
	resp.Body.WriteString(content)

	pp := PlusPlus(resp, req)

	assert.Equal(t, content, pp.ReqBody())

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	assert.Nil(t, err)
	assert.Equal(t, content, string(bodyBytes))
}

func TestPP_ReqBodyMap(t *testing.T) {

	contentBytes := []byte(`{"name":"tom","age":18}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewReader(contentBytes))
	resp := httptest.NewRecorder()

	m := map[string]interface{}{}
	assert.Nil(t, json.Unmarshal(contentBytes, &m))

	pp := PlusPlus(resp, req)
	bodyM, err := pp.ReqBodyMap()
	assert.Nil(t, err)

	for key, value := range m {
		assert.Equal(t, value, bodyM[key])
	}

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	assert.Nil(t, err)

	assert.Equal(t, contentBytes, bodyBytes)
}

type TestUnmarshaler struct {
	m map[string]interface{}
}

func (t *TestUnmarshaler) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &t.m)
}

func TestPP_ReqBodyToUnmarshaler(t *testing.T) {

	contentBytes := []byte(`{"name":"tom","age":18}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewReader(contentBytes))
	resp := httptest.NewRecorder()

	m := map[string]interface{}{}
	assert.Nil(t, json.Unmarshal(contentBytes, &m))

	pp := PlusPlus(resp, req)
	mer := &TestUnmarshaler{m: map[string]interface{}{}}
	assert.Nil(t, pp.ReqBodyToUnmarshaler(mer))

	for key, value := range m {
		assert.Equal(t, value, mer.m[key])
	}

	// 	read again
	bodyBytes, err := ioutil.ReadAll(pp.Req().Body)
	assert.Nil(t, err)

	assert.Equal(t, contentBytes, bodyBytes)

}

func TestPlusPlus(t *testing.T) {


}

func TestPP_Req(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Req(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Req() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Resp(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   http.ResponseWriter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Resp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Resp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_RequestID(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.RequestID(); got != tt.want {
				t.Errorf("PP.RequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_VO(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.VO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.VO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Abort(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Abort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Abort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_IsAbort(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.IsAbort(); got != tt.want {
				t.Errorf("PP.IsAbort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Error(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		code int
		msg  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Error(tt.args.code, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_EmptyError(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.EmptyError(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.EmptyError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_AbortEmptyError(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.AbortEmptyError(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.AbortEmptyError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ErrorMsg(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		message message.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ErrorMsg(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ErrorMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_AbortErrorMsg(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		message message.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.AbortErrorMsg(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.AbortErrorMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Plain(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		code int
		msg  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Plain(tt.args.code, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Plain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_EmptyPlain(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.EmptyPlain(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.EmptyPlain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Redirect(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		url  string
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Redirect(tt.args.url, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Redirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_JSON(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		data   interface{}
		status int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.JSON(tt.args.data, tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_JSONOK(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.JSONOK(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.JSONOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_OK(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.OK(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.OK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Created(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Created(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Created() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Accepted(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Accepted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Accepted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NonAuthoritativeInfo(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NonAuthoritativeInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NonAuthoritativeInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NoContent(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NoContent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NoContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ResetContent(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ResetContent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ResetContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_PartialContent(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.PartialContent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.PartialContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_MultiStatus(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.MultiStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.MultiStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_AlreadyReported(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.AlreadyReported(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.AlreadyReported() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_IMUsed(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.IMUsed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.IMUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_MultipleChoices(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.MultipleChoices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.MultipleChoices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_MovedPermanently(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.MovedPermanently(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.MovedPermanently() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Found(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Found(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Found() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SeeOther(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SeeOther(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.SeeOther() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NotModified(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NotModified(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NotModified() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_UseProxy(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.UseProxy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.UseProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_TemporaryRedirect(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.TemporaryRedirect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.TemporaryRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_PermanentRedirect(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.PermanentRedirect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.PermanentRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_BadRequest(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.BadRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.BadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Unauthorized(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Unauthorized(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Unauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_PaymentRequired(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.PaymentRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.PaymentRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Forbidden(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Forbidden(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Forbidden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NotFound(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NotFound(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NotAllowed(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.MethodNotAllowed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NotAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NotAcceptable(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NotAcceptable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NotAcceptable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ProxyAuthRequired(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ProxyAuthRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ProxyAuthRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_RequestTimeout(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.RequestTimeout(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.RequestTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Conflict(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Conflict(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Conflict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Gone(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Gone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Gone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_LengthRequired(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.LengthRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.LengthRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_PreconditionFailed(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.PreconditionFailed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.PreconditionFailed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_RequestEntityTooLarge(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.RequestEntityTooLarge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.RequestEntityTooLarge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_RequestURITooLong(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.RequestURITooLong(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.RequestURITooLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_UnsupportedMediaType(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.UnsupportedMediaType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.UnsupportedMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_RequestedRangeNotSatisfiable(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.RequestedRangeNotSatisfiable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.RequestedRangeNotSatisfiable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ExpectationFailed(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ExpectationFailed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ExpectationFailed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Teapot(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Teapot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Teapot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_MisdirectedRequest(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.MisdirectedRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.MisdirectedRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_UnprocessableEntity(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.UnprocessableEntity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.UnprocessableEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Locked(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Locked(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Locked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_FailedDependency(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.FailedDependency(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.FailedDependency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_TooEarly(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.TooEarly(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.TooEarly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_UpgradeRequired(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.UpgradeRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.UpgradeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_PreconditionRequired(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.PreconditionRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.PreconditionRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_TooManyRequests(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.TooManyRequests(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.TooManyRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_RequestHeaderFieldsTooLarge(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.RequestHeaderFieldsTooLarge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.RequestHeaderFieldsTooLarge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_UnavailableForLegalReasons(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.UnavailableForLegalReasons(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.UnavailableForLegalReasons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_InternalServerError(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.InternalServerError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.InternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NotImplemented(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NotImplemented(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NotImplemented() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_BadGateway(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.BadGateway(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.BadGateway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ServiceUnavailable(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ServiceUnavailable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ServiceUnavailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GatewayTimeout(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GatewayTimeout(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GatewayTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_HTTPVersionNotSupported(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.HTTPVersionNotSupported(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.HTTPVersionNotSupported() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_VariantAlsoNegotiates(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.VariantAlsoNegotiates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.VariantAlsoNegotiates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_InsufficientStorage(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.InsufficientStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.InsufficientStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_LoopDetected(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.LoopDetected(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.LoopDetected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NotExtended(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NotExtended(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NotExtended() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_NetworkAuthenticationRequired(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.NetworkAuthenticationRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.NetworkAuthenticationRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Get(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Set(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Set(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetR(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.SetR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetStringR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetStringR(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetStringR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetIntR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetIntR(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetIntR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetInt8R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetInt8R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetInt8R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetInt16R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value int16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetInt16R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetInt16R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetInt32R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetInt32R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetInt32R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetInt64R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetInt64R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetInt64R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetUintR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetUintR(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetUintR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetUint8R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetUint8R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetUint8R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetUint16R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetUint16R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetUint16R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetUint32R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetUint32R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetUint32R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetUint64R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetUint64R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetUint64R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetBoolR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetBoolR(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetBoolR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetByteR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetByteR(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetByteR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetBytesR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetBytesR(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.SetBytesR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetTimeR(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetTimeR(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.SetTimeR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetFloat32R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetFloat32R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetFloat32R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SetFloat64R(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SetFloat64R(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("PP.SetFloat64R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetString(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetString(tt.args.key); got != tt.want {
				t.Errorf("PP.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt(tt.args.key); got != tt.want {
				t.Errorf("PP.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt8(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt8(tt.args.key); got != tt.want {
				t.Errorf("PP.GetInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt16(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt16(tt.args.key); got != tt.want {
				t.Errorf("PP.GetInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt32(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt32(tt.args.key); got != tt.want {
				t.Errorf("PP.GetInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt64(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt64(tt.args.key); got != tt.want {
				t.Errorf("PP.GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt(tt.args.key); got != tt.want {
				t.Errorf("PP.GetUInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt8(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt8(tt.args.key); got != tt.want {
				t.Errorf("PP.GetUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt16(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt16(tt.args.key); got != tt.want {
				t.Errorf("PP.GetUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt32(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt32(tt.args.key); got != tt.want {
				t.Errorf("PP.GetUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt64(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt64(tt.args.key); got != tt.want {
				t.Errorf("PP.GetUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetBool(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetBool(tt.args.key); got != tt.want {
				t.Errorf("PP.GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetByte(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetByte(tt.args.key); got != tt.want {
				t.Errorf("PP.GetByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetBytes(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetBytes(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GetBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetTime(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetTime(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetFloat32(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetFloat32(tt.args.key); got != tt.want {
				t.Errorf("PP.GetFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetFloat64(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetFloat64(tt.args.key); got != tt.want {
				t.Errorf("PP.GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetDf(tt.args.key, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GetDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetStringDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetStringDf(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetStringDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetIntDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetIntDf(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetIntDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt8Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt8Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetInt8Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt16Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue int16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt16Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetInt16Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt32Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt32Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetInt32Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetInt64Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetInt64Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetInt64Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUIntDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUIntDf(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetUIntDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt8Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt8Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetUInt8Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt16Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt16Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetUInt16Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt32Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt32Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetUInt32Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetUInt64Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetUInt64Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetUInt64Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetBoolDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetBoolDf(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetBoolDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetByteDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetByteDf(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetByteDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetBytesDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetBytesDf(tt.args.key, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GetBytesDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetTimeDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetTimeDf(tt.args.key, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GetTimeDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetFloat32Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetFloat32Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetFloat32Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetFloat64Df(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetFloat64Df(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetFloat64Df() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_CallbackByCode(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		errorCode int
		respData  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.CallbackByCode(tt.args.errorCode, tt.args.respData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.CallbackByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Queries(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   url.Values
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Queries(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Queries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetQuery(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantValue  string
		wantExists bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			gotValue, gotExists := p.GetQuery(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("PP.GetQuery() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotExists != tt.wantExists {
				t.Errorf("PP.GetQuery() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestPP_Query(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if gotValue := p.Query(tt.args.key); gotValue != tt.wantValue {
				t.Errorf("PP.Query() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestPP_GetQueryDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if gotValue := p.GetQueryDf(tt.args.key, tt.args.defaultValue); gotValue != tt.wantValue {
				t.Errorf("PP.GetQueryDf() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestPP_GetHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetHeader(tt.args.key); got != tt.want {
				t.Errorf("PP.GetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetHeaderDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetHeaderDf(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetHeaderDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetHeaderValues(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetHeaderValues(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GetHeaderValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_SplitHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key      string
		splitSep string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.SplitHeader(tt.args.key, tt.args.splitSep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.SplitHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_WriteHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.WriteHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.WriteHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_WriteHeaders(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		headers map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.WriteHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.WriteHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_AppendHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key    string
		values []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.AppendHeader(tt.args.key, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.AppendHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_AppendHeaders(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		headers map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.AppendHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.AppendHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetRespHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetRespHeader(tt.args.key); got != tt.want {
				t.Errorf("PP.GetRespHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetRespHeaderDf(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetRespHeaderDf(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("PP.GetRespHeaderDf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetRespHeaderValues(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetRespHeaderValues(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.GetRespHeaderValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_WriteRespHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.WriteRespHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.WriteRespHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_WriteRespHeaders(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		headers map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.WriteRespHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.WriteRespHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_AppendRespHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		key    string
		values []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.AppendRespHeader(tt.args.key, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.AppendRespHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_AppendRespHeaders(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		headers map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.AppendRespHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.AppendRespHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ReqHeader(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ReqHeader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ReqHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ReqURL(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ReqURL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ReqURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ReqHost(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ReqHost(); got != tt.want {
				t.Errorf("PP.ReqHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Handler(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Handler(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_ServeHTTP(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.ServeHTTP(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.ServeHTTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_FormFile(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *multipart.FileHeader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			got, err := p.FormFile(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("PP.FormFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.FormFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Status(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Status(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetStatus(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetStatus(); got != tt.want {
				t.Errorf("PP.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_CopyReq(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.CopyReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PP.CopyReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_Method(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.Method(); got != tt.want {
				t.Errorf("PP.Method() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPP_GetClientIP(t *testing.T) {
	type fields struct {
		w http.ResponseWriter
		r *http.Request
		h http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlusPlus(tt.fields.w, tt.fields.r)
			if got := p.GetClientIP(); got != tt.want {
				t.Errorf("PP.GetClientIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
