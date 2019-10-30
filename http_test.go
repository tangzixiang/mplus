package mplus

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/tangzixiang/mplus/message"
)

func TestAbort(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "", nil)
	assert.Equal(t, err, nil)

	req = req.WithContext(NewContext(req.Context()))
	assert.Equal(t, IsAbort(req), false)
	assert.Equal(t, IsAbort(Abort(req)), true)
}

func TestError(t *testing.T) {
	respR := httptest.NewRecorder()

	respW := NewResponseWrite(respR)
	assert.Equal(t, respW.Status(), http.StatusOK)

	Error(respW, MessageStatusInternalServerError.Set(http.StatusText(http.StatusInternalServerError)))
	assert.Equal(t, respW.Status(), http.StatusInternalServerError)

	resp := respR.Result()
	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, strings.TrimSpace(string(bodyBytes)), http.StatusText(http.StatusInternalServerError))
}

func TestDumpRequest(t *testing.T) {

	content := NewQuery().AddPairs("name", "tom", "age", "18").Encode()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(content))

	assert.Equal(t, content, DumpRequest(req))

	bodyBytes, err := ioutil.ReadAll(req.Body)

	assert.Nil(t, err)
	assert.Equal(t, content, string(bodyBytes))
}

func TestCopyRequest(t *testing.T) {
	content := NewQuery().AddPairs("name", "tom", "age", "18").Encode()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(content))

	assert.Nil(t, req.ParseForm())

	req2 := CopyRequest(req)

	assert.Equal(t, req.Proto, req2.Proto)
	assert.Equal(t, req.ProtoMajor, req2.ProtoMajor)
	assert.Equal(t, req.ProtoMinor, req2.ProtoMinor)
	assert.Equal(t, req.Method, req2.Method)
	assert.Equal(t, req.ContentLength, req2.ContentLength)
	assert.Equal(t, req.TransferEncoding, req2.TransferEncoding)
	assert.Equal(t, req.Host, req2.Host)
	assert.Equal(t, req.RemoteAddr, req2.RemoteAddr)
	assert.Equal(t, req.RequestURI, req2.RequestURI)
	assert.Equal(t, req.URL, req2.URL)

	for key, values := range req.Header {
		assert.ElementsMatch(t, values, req2.Header[key])
	}
	for key, values := range req.Form {
		assert.ElementsMatch(t, values, req2.Form[key])
	}
	for key, values := range req.PostForm {
		assert.ElementsMatch(t, values, req2.PostForm[key])
	}
}

func TestCopyRequest_file(t *testing.T) {
	var buff bytes.Buffer
	multipartWriter := multipart.NewWriter(&buff)

	fw, err := multipartWriter.CreateFormFile("pic", "pic.png")
	assert.Nil(t, err)

	fw.Write([]byte(`pic`))

	assert.Nil(t, multipartWriter.WriteField("name", "tom"))

	multipartWriter.Close()

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", &buff)

	// set multipart/form-data; boundary=xxxxx
	SetRequestHeader(req, HeaderContentType, multipartWriter.FormDataContentType())

	assert.Nil(t, req.ParseMultipartForm(DefaultMemorySize()))

	req2 := CopyRequest(req)

	assert.Equal(t, req.Proto, req2.Proto)
	assert.Equal(t, req.ProtoMajor, req2.ProtoMajor)
	assert.Equal(t, req.ProtoMinor, req2.ProtoMinor)
	assert.Equal(t, req.Method, req2.Method)
	assert.Equal(t, req.ContentLength, req2.ContentLength)
	assert.Equal(t, req.TransferEncoding, req2.TransferEncoding)
	assert.Equal(t, req.Host, req2.Host)
	assert.Equal(t, req.RemoteAddr, req2.RemoteAddr)
	assert.Equal(t, req.RequestURI, req2.RequestURI)
	assert.Equal(t, req.URL, req2.URL)

	for key, values := range req.Header {
		assert.ElementsMatch(t, values, req2.Header[key])
	}
	for key, values := range req.Form {
		assert.ElementsMatch(t, values, req2.Form[key])
	}
	for key, values := range req.PostForm {
		assert.ElementsMatch(t, values, req2.PostForm[key])
	}

	assert.NotEqual(t, 0, len(req2.MultipartForm.Value))
	assert.NotEqual(t, 0, len(req2.MultipartForm.File))
	assert.Equal(t, req.FormValue("name"), req2.FormValue("name"))

	mf, _, err := req.FormFile("pic")
	assert.Nil(t, err)

	reqPicData, err := ioutil.ReadAll(mf)
	assert.Nil(t, err)

	mf.Close()

	mf, _, err = req2.FormFile("pic")
	assert.Nil(t, err)

	req2PicData, err := ioutil.ReadAll(mf)
	assert.Nil(t, err)

	mf.Close()

	assert.ElementsMatch(t, reqPicData, req2PicData)
}

func TestCopyRequest_context(t *testing.T) {

	jsonstr := `{"name":"Tom","age":50,"gender":"male","email":"10086@fox.com"}`
	buffer := bytes.NewBuffer([]byte(jsonstr))

	req01 := httptest.NewRequest(http.MethodPost, "http://localhost", buffer)
	SetRequestHeader(req01, HeaderContentType, MIMEJSON)

	req01 = req01.WithContext(NewContext(req01.Context()))
	req02 := CopyRequest(req01)

	assert.Equal(t, GetContextValue(req01.Context(), "1"), GetContextValue(req02.Context(), "1"))
	assert.Equal(t, GetHeader(req01, HeaderContentType), GetHeader(req02, HeaderContentType))

	SetContextValue(req01.Context(), "1", "1")
	SetContextValue(req02.Context(), "1", "2")
	assert.NotEqual(t, GetContextValue(req01.Context(), "1"), GetContextValue(req02.Context(), "1"))

	req01.Header.Del(HeaderContentType)
	assert.NotEqual(t, GetHeader(req01, HeaderContentType), GetHeader(req02, HeaderContentType))
}

func Test_responseWrite_SetStatus(t *testing.T) {

	recorder := httptest.NewRecorder()
	respW := NewResponseWrite(recorder)

	// default status is 200
	assert.Equal(t, http.StatusOK, respW.Status())

	respW.SetStatus(http.StatusBadRequest)
	assert.Equal(t, http.StatusBadRequest, respW.Status())

	// SetStatus not set the StatusCode for http.ResponseWriter
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func Test_responseWrite_Status(t *testing.T) {

	recorder := httptest.NewRecorder()
	respW := NewResponseWrite(recorder)

	// default status is 200
	assert.Equal(t, http.StatusOK, respW.Status())

	respW.SetStatus(http.StatusBadRequest)
	assert.Equal(t, http.StatusBadRequest, respW.Status())

	// SetStatus not set the StatusCode for http.ResponseWriter
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func Test_responseWrite_WriteHead(t *testing.T) {
	recorder := httptest.NewRecorder()
	respW := NewResponseWrite(recorder)

	// default status is 200
	assert.Equal(t, http.StatusOK, respW.Status())

	//  WriteHeader set http.ResponseWriter status, not set responseWrite
	respW.WriteHeader(http.StatusBadRequest)

	assert.Equal(t, http.StatusOK, respW.Status())
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestNewResponseWrite(t *testing.T) {

	recorder := httptest.NewRecorder()
	respW := NewResponseWrite(recorder)

	assert.Implements(t, (*ResponseWriter)(nil), respW)
	assert.Implements(t, (*http.ResponseWriter)(nil), respW)
}

func TestGetHTTPRespStatus(t *testing.T) {
	recorder := httptest.NewRecorder()
	respW := NewResponseWrite(recorder)

	SetHTTPRespStatus(respW, http.StatusBadRequest)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
	assert.Equal(t, http.StatusBadRequest, respW.Status())
	assert.Equal(t, http.StatusBadRequest, GetHTTPRespStatus(respW))
}

func TestUnWrapResponseWriter(t *testing.T) {
	recorder := httptest.NewRecorder()
	_recorder, ok := UnWrapResponseWriter(NewResponseWrite(recorder)).(*httptest.ResponseRecorder)

	assert.True(t, ok)
	assert.Equal(t, recorder, _recorder)
}

func TestSetHTTPRespStatus(t *testing.T) {
	recorder := httptest.NewRecorder()
	respW := NewResponseWrite(recorder)

	SetHTTPRespStatus(respW, http.StatusBadRequest, false)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, http.StatusBadRequest, respW.Status())
	assert.Equal(t, http.StatusBadRequest, GetHTTPRespStatus(respW))

	recorder = httptest.NewRecorder()
	respW = NewResponseWrite(recorder)
	SetHTTPRespStatus(respW, http.StatusBadRequest)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
	assert.Equal(t, http.StatusBadRequest, respW.Status())
	assert.Equal(t, http.StatusBadRequest, GetHTTPRespStatus(respW))

}

func TestStatusCodeMethod(t *testing.T) {

	getRequest := func() *http.Request { return httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil) }

	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		h          http.HandlerFunc
		statusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test OK", args: args{w: httptest.NewRecorder(), r: getRequest(), h: OK, statusCode: http.StatusOK}},
		{name: "test Created", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Created, statusCode: http.StatusCreated}},
		{name: "test Accepted", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Accepted, statusCode: http.StatusAccepted}},
		{name: "test NonAuthoritativeInfo", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NonAuthoritativeInfo, statusCode: http.StatusNonAuthoritativeInfo}},
		{name: "test NoContent", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NoContent, statusCode: http.StatusNoContent}},
		{name: "test ResetContent", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ResetContent, statusCode: http.StatusResetContent}},
		{name: "test PartialContent", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PartialContent, statusCode: http.StatusPartialContent}},
		{name: "test MultiStatus", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MultiStatus, statusCode: http.StatusMultiStatus}},
		{name: "test AlreadyReported", args: args{w: httptest.NewRecorder(), r: getRequest(), h: AlreadyReported, statusCode: http.StatusAlreadyReported}},
		{name: "test IMUsed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: IMUsed, statusCode: http.StatusIMUsed}},
		{name: "test MultipleChoices", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MultipleChoices, statusCode: http.StatusMultipleChoices}},
		{name: "test MovedPermanently", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MovedPermanently, statusCode: http.StatusMovedPermanently}},
		{name: "test Found", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Found, statusCode: http.StatusFound}},
		{name: "test SeeOther", args: args{w: httptest.NewRecorder(), r: getRequest(), h: SeeOther, statusCode: http.StatusSeeOther}},
		{name: "test NotModified", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotModified, statusCode: http.StatusNotModified}},
		{name: "test UseProxy", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UseProxy, statusCode: http.StatusUseProxy}},
		{name: "test TemporaryRedirect", args: args{w: httptest.NewRecorder(), r: getRequest(), h: TemporaryRedirect, statusCode: http.StatusTemporaryRedirect}},
		{name: "test PermanentRedirect", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PermanentRedirect, statusCode: http.StatusPermanentRedirect}},
		{name: "test BadRequest", args: args{w: httptest.NewRecorder(), r: getRequest(), h: BadRequest, statusCode: http.StatusBadRequest}},
		{name: "test Unauthorized", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Unauthorized, statusCode: http.StatusUnauthorized}},
		{name: "test PaymentRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PaymentRequired, statusCode: http.StatusPaymentRequired}},
		{name: "test Forbidden", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Forbidden, statusCode: http.StatusForbidden}},
		{name: "test NotFound", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotFound, statusCode: http.StatusNotFound}},
		{name: "test MethodNotAllowed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MethodNotAllowed, statusCode: http.StatusMethodNotAllowed}},
		{name: "test NotAcceptable", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotAcceptable, statusCode: http.StatusNotAcceptable}},
		{name: "test ProxyAuthRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ProxyAuthRequired, statusCode: http.StatusProxyAuthRequired}},
		{name: "test RequestTimeout", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestTimeout, statusCode: http.StatusRequestTimeout}},
		{name: "test Conflict", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Conflict, statusCode: http.StatusConflict}},
		{name: "test Gone", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Gone, statusCode: http.StatusGone}},
		{name: "test LengthRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: LengthRequired, statusCode: http.StatusLengthRequired}},
		{name: "test PreconditionFailed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PreconditionFailed, statusCode: http.StatusPreconditionFailed}},
		{name: "test RequestEntityTooLarge", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestEntityTooLarge, statusCode: http.StatusRequestEntityTooLarge}},
		{name: "test RequestURITooLong", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestURITooLong, statusCode: http.StatusRequestURITooLong}},
		{name: "test UnsupportedMediaType", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UnsupportedMediaType, statusCode: http.StatusUnsupportedMediaType}},
		{name: "test RequestedRangeNotSatisfiable", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestedRangeNotSatisfiable, statusCode: http.StatusRequestedRangeNotSatisfiable}},
		{name: "test ExpectationFailed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ExpectationFailed, statusCode: http.StatusExpectationFailed}},
		{name: "test Teapot", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Teapot, statusCode: http.StatusTeapot}},
		{name: "test MisdirectedRequest", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MisdirectedRequest, statusCode: http.StatusMisdirectedRequest}},
		{name: "test UnprocessableEntity", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UnprocessableEntity, statusCode: http.StatusUnprocessableEntity}},
		{name: "test Locked", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Locked, statusCode: http.StatusLocked}},
		{name: "test FailedDependency", args: args{w: httptest.NewRecorder(), r: getRequest(), h: FailedDependency, statusCode: http.StatusFailedDependency}},
		{name: "test TooEarly", args: args{w: httptest.NewRecorder(), r: getRequest(), h: TooEarly, statusCode: http.StatusTooEarly}},
		{name: "test UpgradeRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UpgradeRequired, statusCode: http.StatusUpgradeRequired}},
		{name: "test PreconditionRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PreconditionRequired, statusCode: http.StatusPreconditionRequired}},
		{name: "test TooManyRequests", args: args{w: httptest.NewRecorder(), r: getRequest(), h: TooManyRequests, statusCode: http.StatusTooManyRequests}},
		{name: "test RequestHeaderFieldsTooLarge", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestHeaderFieldsTooLarge, statusCode: http.StatusRequestHeaderFieldsTooLarge}},
		{name: "test UnavailableForLegalReasons", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UnavailableForLegalReasons, statusCode: http.StatusUnavailableForLegalReasons}},
		{name: "test InternalServerError", args: args{w: httptest.NewRecorder(), r: getRequest(), h: InternalServerError, statusCode: http.StatusInternalServerError}},
		{name: "test NotImplemented", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotImplemented, statusCode: http.StatusNotImplemented}},
		{name: "test BadGateway", args: args{w: httptest.NewRecorder(), r: getRequest(), h: BadGateway, statusCode: http.StatusBadGateway}},
		{name: "test ServiceUnavailable", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ServiceUnavailable, statusCode: http.StatusServiceUnavailable}},
		{name: "test GatewayTimeout", args: args{w: httptest.NewRecorder(), r: getRequest(), h: GatewayTimeout, statusCode: http.StatusGatewayTimeout}},
		{name: "test HTTPVersionNotSupported", args: args{w: httptest.NewRecorder(), r: getRequest(), h: HTTPVersionNotSupported, statusCode: http.StatusHTTPVersionNotSupported}},
		{name: "test VariantAlsoNegotiates", args: args{w: httptest.NewRecorder(), r: getRequest(), h: VariantAlsoNegotiates, statusCode: http.StatusVariantAlsoNegotiates}},
		{name: "test InsufficientStorage", args: args{w: httptest.NewRecorder(), r: getRequest(), h: InsufficientStorage, statusCode: http.StatusInsufficientStorage}},
		{name: "test LoopDetected", args: args{w: httptest.NewRecorder(), r: getRequest(), h: LoopDetected, statusCode: http.StatusLoopDetected}},
		{name: "test NotExtended", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotExtended, statusCode: http.StatusNotExtended}},
		{name: "test NetworkAuthenticationRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NetworkAuthenticationRequired, statusCode: http.StatusNetworkAuthenticationRequired}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PreMiddleware(func(w http.ResponseWriter, r *http.Request) {
				assertResp(t, tt.args.h, w, r, tt.args.statusCode)

			}).ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func TestRegisterHttpStatusMethod(t *testing.T) {

	getRequest := func() *http.Request { return httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil) }

	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		h          http.HandlerFunc
		statusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test OK", args: args{w: httptest.NewRecorder(), r: getRequest(), h: OK, statusCode: http.StatusOK}},
		{name: "test Created", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Created, statusCode: http.StatusCreated}},
		{name: "test Accepted", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Accepted, statusCode: http.StatusAccepted}},
		{name: "test NonAuthoritativeInfo", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NonAuthoritativeInfo, statusCode: http.StatusNonAuthoritativeInfo}},
		{name: "test NoContent", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NoContent, statusCode: http.StatusNoContent}},
		{name: "test ResetContent", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ResetContent, statusCode: http.StatusResetContent}},
		{name: "test PartialContent", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PartialContent, statusCode: http.StatusPartialContent}},
		{name: "test MultiStatus", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MultiStatus, statusCode: http.StatusMultiStatus}},
		{name: "test AlreadyReported", args: args{w: httptest.NewRecorder(), r: getRequest(), h: AlreadyReported, statusCode: http.StatusAlreadyReported}},
		{name: "test IMUsed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: IMUsed, statusCode: http.StatusIMUsed}},
		{name: "test MultipleChoices", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MultipleChoices, statusCode: http.StatusMultipleChoices}},
		{name: "test MovedPermanently", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MovedPermanently, statusCode: http.StatusMovedPermanently}},
		{name: "test Found", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Found, statusCode: http.StatusFound}},
		{name: "test SeeOther", args: args{w: httptest.NewRecorder(), r: getRequest(), h: SeeOther, statusCode: http.StatusSeeOther}},
		{name: "test NotModified", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotModified, statusCode: http.StatusNotModified}},
		{name: "test UseProxy", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UseProxy, statusCode: http.StatusUseProxy}},
		{name: "test TemporaryRedirect", args: args{w: httptest.NewRecorder(), r: getRequest(), h: TemporaryRedirect, statusCode: http.StatusTemporaryRedirect}},
		{name: "test PermanentRedirect", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PermanentRedirect, statusCode: http.StatusPermanentRedirect}},
		{name: "test BadRequest", args: args{w: httptest.NewRecorder(), r: getRequest(), h: BadRequest, statusCode: http.StatusBadRequest}},
		{name: "test Unauthorized", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Unauthorized, statusCode: http.StatusUnauthorized}},
		{name: "test PaymentRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PaymentRequired, statusCode: http.StatusPaymentRequired}},
		{name: "test Forbidden", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Forbidden, statusCode: http.StatusForbidden}},
		{name: "test NotFound", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotFound, statusCode: http.StatusNotFound}},
		{name: "test MethodNotAllowed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MethodNotAllowed, statusCode: http.StatusMethodNotAllowed}},
		{name: "test NotAcceptable", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotAcceptable, statusCode: http.StatusNotAcceptable}},
		{name: "test ProxyAuthRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ProxyAuthRequired, statusCode: http.StatusProxyAuthRequired}},
		{name: "test RequestTimeout", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestTimeout, statusCode: http.StatusRequestTimeout}},
		{name: "test Conflict", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Conflict, statusCode: http.StatusConflict}},
		{name: "test Gone", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Gone, statusCode: http.StatusGone}},
		{name: "test LengthRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: LengthRequired, statusCode: http.StatusLengthRequired}},
		{name: "test PreconditionFailed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PreconditionFailed, statusCode: http.StatusPreconditionFailed}},
		{name: "test RequestEntityTooLarge", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestEntityTooLarge, statusCode: http.StatusRequestEntityTooLarge}},
		{name: "test RequestURITooLong", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestURITooLong, statusCode: http.StatusRequestURITooLong}},
		{name: "test UnsupportedMediaType", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UnsupportedMediaType, statusCode: http.StatusUnsupportedMediaType}},
		{name: "test RequestedRangeNotSatisfiable", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestedRangeNotSatisfiable, statusCode: http.StatusRequestedRangeNotSatisfiable}},
		{name: "test ExpectationFailed", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ExpectationFailed, statusCode: http.StatusExpectationFailed}},
		{name: "test Teapot", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Teapot, statusCode: http.StatusTeapot}},
		{name: "test MisdirectedRequest", args: args{w: httptest.NewRecorder(), r: getRequest(), h: MisdirectedRequest, statusCode: http.StatusMisdirectedRequest}},
		{name: "test UnprocessableEntity", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UnprocessableEntity, statusCode: http.StatusUnprocessableEntity}},
		{name: "test Locked", args: args{w: httptest.NewRecorder(), r: getRequest(), h: Locked, statusCode: http.StatusLocked}},
		{name: "test FailedDependency", args: args{w: httptest.NewRecorder(), r: getRequest(), h: FailedDependency, statusCode: http.StatusFailedDependency}},
		{name: "test TooEarly", args: args{w: httptest.NewRecorder(), r: getRequest(), h: TooEarly, statusCode: http.StatusTooEarly}},
		{name: "test UpgradeRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UpgradeRequired, statusCode: http.StatusUpgradeRequired}},
		{name: "test PreconditionRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: PreconditionRequired, statusCode: http.StatusPreconditionRequired}},
		{name: "test TooManyRequests", args: args{w: httptest.NewRecorder(), r: getRequest(), h: TooManyRequests, statusCode: http.StatusTooManyRequests}},
		{name: "test RequestHeaderFieldsTooLarge", args: args{w: httptest.NewRecorder(), r: getRequest(), h: RequestHeaderFieldsTooLarge, statusCode: http.StatusRequestHeaderFieldsTooLarge}},
		{name: "test UnavailableForLegalReasons", args: args{w: httptest.NewRecorder(), r: getRequest(), h: UnavailableForLegalReasons, statusCode: http.StatusUnavailableForLegalReasons}},
		{name: "test InternalServerError", args: args{w: httptest.NewRecorder(), r: getRequest(), h: InternalServerError, statusCode: http.StatusInternalServerError}},
		{name: "test NotImplemented", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotImplemented, statusCode: http.StatusNotImplemented}},
		{name: "test BadGateway", args: args{w: httptest.NewRecorder(), r: getRequest(), h: BadGateway, statusCode: http.StatusBadGateway}},
		{name: "test ServiceUnavailable", args: args{w: httptest.NewRecorder(), r: getRequest(), h: ServiceUnavailable, statusCode: http.StatusServiceUnavailable}},
		{name: "test GatewayTimeout", args: args{w: httptest.NewRecorder(), r: getRequest(), h: GatewayTimeout, statusCode: http.StatusGatewayTimeout}},
		{name: "test HTTPVersionNotSupported", args: args{w: httptest.NewRecorder(), r: getRequest(), h: HTTPVersionNotSupported, statusCode: http.StatusHTTPVersionNotSupported}},
		{name: "test VariantAlsoNegotiates", args: args{w: httptest.NewRecorder(), r: getRequest(), h: VariantAlsoNegotiates, statusCode: http.StatusVariantAlsoNegotiates}},
		{name: "test InsufficientStorage", args: args{w: httptest.NewRecorder(), r: getRequest(), h: InsufficientStorage, statusCode: http.StatusInsufficientStorage}},
		{name: "test LoopDetected", args: args{w: httptest.NewRecorder(), r: getRequest(), h: LoopDetected, statusCode: http.StatusLoopDetected}},
		{name: "test NotExtended", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NotExtended, statusCode: http.StatusNotExtended}},
		{name: "test NetworkAuthenticationRequired", args: args{w: httptest.NewRecorder(), r: getRequest(), h: NetworkAuthenticationRequired, statusCode: http.StatusNetworkAuthenticationRequired}},
	}

	for _, tt := range tests {
		RegisterHttpStatusMethod(tt.args.statusCode, func(w http.ResponseWriter, r *http.Request, m message.Message, statusCode int) {
			Abort(r)
			JSON(w, r, Data{}, statusCode)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PreMiddleware(func(w http.ResponseWriter, r *http.Request) {
				assertJSONResp(t, tt.args.h, w, r, tt.args.statusCode)
			}).ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func assertResp(t *testing.T, handler http.HandlerFunc, w http.ResponseWriter, r *http.Request, statusCode int) {
	handler(w, r)

	requestMediaType, err := ParseRequestMediaType(r)
	assert.Nil(t, err)
	assert.Equal(t, MIMEPOSTForm, requestMediaType)

	respMediaType, err := ParseResponseMediaType(w)
	assert.Nil(t, err)
	assert.Equal(t, MIMEPlain, respMediaType)
	assert.Equal(t, statusCode, GetHTTPRespStatus(w))
}

func assertJSONResp(t *testing.T, handler http.HandlerFunc, w http.ResponseWriter, r *http.Request, statusCode int) {
	handler(w, r)

	requestMediaType, err := ParseRequestMediaType(r)
	assert.Nil(t, err)
	assert.Equal(t, MIMEPOSTForm, requestMediaType)

	respMediaType, err := ParseResponseMediaType(w)
	assert.Nil(t, err)
	assert.Equal(t, MIMEJSON, respMediaType)
	assert.Equal(t, statusCode, GetHTTPRespStatus(w))

	assert.True(t, IsAbort(r))

	bodyBytes, err := ioutil.ReadAll(UnWrapResponseWriter(w).(*httptest.ResponseRecorder).Body)
	assert.Nil(t, err)

	assert.JSONEq(t, string(bodyBytes), `{}`)
}

func TestSetDefaultMemorySize(t *testing.T) {

	size := int64(64 * 1024 * 1024)

	assert.NotEqual(t, size, DefaultMemorySize())

	SetDefaultMemorySize(size)
	assert.Equal(t, size, DefaultMemorySize())
}

func TestNotAbort(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	assert.False(t, IsAbort(r))
	Abort(r)
	assert.True(t, IsAbort(r))
	NotAbort(r)
	assert.False(t, IsAbort(r))
}

func TestIsAbort(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	assert.False(t, IsAbort(r))
	Abort(r)
	assert.True(t, IsAbort(r))
}

func TestErrorEmpty(t *testing.T) {
	respR := httptest.NewRecorder()

	respW := NewResponseWrite(respR)
	assert.Equal(t, respW.Status(), http.StatusOK)

	ErrorEmpty(respW, MessageStatusInternalServerError.Set(http.StatusText(http.StatusInternalServerError)))
	assert.Equal(t, respW.Status(), http.StatusInternalServerError)

	resp := respR.Result()
	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, strings.TrimSpace(string(bodyBytes)), "")
}

func TestPlain(t *testing.T) {
	respR := httptest.NewRecorder()

	respW := NewResponseWrite(respR)
	assert.Equal(t, respW.Status(), http.StatusOK)

	Plain(respW, MessageStatusBadRequest.Set(http.StatusText(http.StatusBadRequest)))
	assert.Equal(t, respW.Status(), http.StatusBadRequest)

	resp := respR.Result()
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, strings.TrimSpace(string(bodyBytes)), http.StatusText(http.StatusBadRequest))
}

func TestPlainEmpty(t *testing.T) {
	respR := httptest.NewRecorder()

	respW := NewResponseWrite(respR)
	assert.Equal(t, respW.Status(), http.StatusOK)

	PlainEmpty(respW, MessageStatusBadRequest.Set(http.StatusText(http.StatusBadRequest)))
	assert.Equal(t, respW.Status(), http.StatusBadRequest)

	resp := respR.Result()
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, strings.TrimSpace(string(bodyBytes)), "")
}

func TestAbortEmptyError(t *testing.T) {
	respR := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	respW := NewResponseWrite(respR)
	assert.Equal(t, respW.Status(), http.StatusOK)

	AbortEmptyError(respW, r, MessageStatusInternalServerError.Set(http.StatusText(http.StatusInternalServerError)))
	assert.Equal(t, respW.Status(), http.StatusInternalServerError)

	resp := respR.Result()
	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, strings.TrimSpace(string(bodyBytes)), "")
	assert.True(t, IsAbort(r))
}

func TestAbortPlain(t *testing.T) {
	respR := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	respW := NewResponseWrite(respR)
	assert.Equal(t, respW.Status(), http.StatusOK)

	AbortEmptyPlain(respW, r, MessageStatusBadRequest.Set(http.StatusText(http.StatusBadRequest)))
	assert.Equal(t, respW.Status(), http.StatusBadRequest)

	resp := respR.Result()
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, strings.TrimSpace(string(bodyBytes)), "")
	assert.True(t, IsAbort(r))
}

func TestJSONOK(t *testing.T) {
	respR := httptest.NewRecorder()
	w := NewResponseWrite(respR)
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	JSONOK(w, r, Data{})

	requestMediaType, err := ParseRequestMediaType(r)
	assert.Nil(t, err)
	assert.Equal(t, MIMEPOSTForm, requestMediaType)

	respMediaType, err := ParseResponseMediaType(w)
	assert.Nil(t, err)
	assert.Equal(t, MIMEJSON, respMediaType)
	assert.Equal(t, http.StatusOK, GetHTTPRespStatus(w))

	assert.True(t, IsAbort(r))

	bodyBytes, err := ioutil.ReadAll(UnWrapResponseWriter(w).(*httptest.ResponseRecorder).Body)
	assert.Nil(t, err)

	assert.JSONEq(t, string(bodyBytes), `{}`)
}

func TestJSON(t *testing.T) {
	respR := httptest.NewRecorder()
	w := NewResponseWrite(respR)
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	JSON(w, r, Data{}, http.StatusBadRequest)

	requestMediaType, err := ParseRequestMediaType(r)
	assert.Nil(t, err)
	assert.Equal(t, MIMEPOSTForm, requestMediaType)

	respMediaType, err := ParseResponseMediaType(w)
	assert.Nil(t, err)
	assert.Equal(t, MIMEJSON, respMediaType)
	assert.Equal(t, http.StatusBadRequest, GetHTTPRespStatus(w))

	assert.True(t, IsAbort(r))

	bodyBytes, err := ioutil.ReadAll(UnWrapResponseWriter(w).(*httptest.ResponseRecorder).Body)
	assert.Nil(t, err)

	assert.JSONEq(t, string(bodyBytes), `{}`)
}

func TestRedirect(t *testing.T) {
	respR := httptest.NewRecorder()
	w := NewResponseWrite(respR)
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
	r = r.WithContext(NewContext(r.Context()))

	Redirect(w, r, "http://whatever.com/", http.StatusMovedPermanently)

	assert.True(t, IsAbort(r))

	respMediaType, err := ParseResponseMediaType(w)
	assert.Nil(t, err)
	assert.Equal(t, MIMEHTML, respMediaType)
	assert.Equal(t, http.StatusMovedPermanently, GetHTTPRespStatus(w))

	resp := UnWrapResponseWriter(w).(*httptest.ResponseRecorder).Result()

	assert.Equal(t, http.StatusMovedPermanently, resp.StatusCode)
	assert.Equal(t, "http://whatever.com/", GetResponseHeader(w, HeaderLocation))
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)
	assert.Equal(t, `<a href="http://whatever.com/">Moved Permanently</a>.`, strings.TrimSpace(string(bodyBytes)))
}

func TestDumpRequestPure(t *testing.T) {

	content := NewQuery().AddPairs("name", "tom", "age", "18").Encode()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(content))

	assert.ElementsMatch(t, []byte(content), DumpRequestPure(req))

	bodyBytes, err := ioutil.ReadAll(req.Body)

	assert.Nil(t, err)
	assert.Equal(t, content, string(bodyBytes))
}
