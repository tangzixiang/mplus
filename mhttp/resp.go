package mhttp

import (
	"io"
	"net/http"
)

const defaultStatus = http.StatusOK

type responseWrite struct {
	http.ResponseWriter

	status int
}

// ResponseWriter 请求响应对象
type ResponseWriter interface {
	http.ResponseWriter

	SetStatus(status int)
	Status() int

	WriteHead(int)
}

var (
	_ http.ResponseWriter = &responseWrite{}
	_ ResponseWriter      = &responseWrite{}
)

func (w *responseWrite) SetStatus(status int) {
	w.status = status
}

func (w *responseWrite) Status() int {
	return w.status
}

func (w *responseWrite) WriteHead(statusCode int) {
	w.SetStatus(statusCode)
	w.ResponseWriter.WriteHeader(statusCode)
}

// ReaderFrom 将指定流写入响应内
func (w *responseWrite) ReaderFrom(src io.Reader) (n int64, err error) {
	return w.ResponseWriter.(io.ReaderFrom).ReadFrom(src)
}

// NewResponseWrite 包装 http.ResponseWriter 获得额外的方法及数据
func NewResponseWrite(w http.ResponseWriter) ResponseWriter {
	return &responseWrite{ResponseWriter: w, status: defaultStatus}
}

// GetHTTPRespStatus 获取响应状态，
// w 应为 responseWrite 实例，responseWrite 实例只能获取通过 NewResponseWrite 获取
func GetHTTPRespStatus(w http.ResponseWriter) int {
	if respW, ok := w.(ResponseWriter); ok {
		return respW.Status()
	}

	return http.StatusOK
}

// SetHTTPRespStatus 设置响应状态,
// w 应为 responseWrite 实例，responseWrite 实例只能获取通过 NewResponseWrite 获取
// coverSupper 用于决定是否覆盖到底层的 http.ResponseWriter,默认 true
// 若 coverSupper 设置为 false 则仅更新 responseWrite 实例的 status，不更新底层 http.ResponseWriter
// 若无特殊使用需求，coverSupper 采用默认值即可
func SetHTTPRespStatus(w http.ResponseWriter, statusCode int, coverSupper ...bool) http.ResponseWriter {
	cover := true
	if len(coverSupper) > 0 && !coverSupper[0] {
		cover = false
	}

	rw, ok := w.(ResponseWriter)

	if !ok {// 不支持非 ResponseWriter
		return w
	}

	if cover {
		rw.WriteHead(statusCode)
	} else {
		rw.SetStatus(statusCode)
	}
	return w
}

// UnWrapResponseWriter 解包 ResponseWriter 获取内部的 http.ResponseWriter
// 当前方法与 NewResponseWrite 相对应
func UnWrapResponseWriter(resp http.ResponseWriter) http.ResponseWriter {
	_resp, ok := resp.(*responseWrite)
	if !ok {
		return nil
	}
	return _resp.ResponseWriter
}
