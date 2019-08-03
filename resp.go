package mplus

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
	ReaderFrom(src io.Reader) (n int64, err error)
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

func (w *responseWrite) WriteHead(code int) {
	w.SetStatus(code)
	w.ResponseWriter.WriteHeader(code)
}

// ReaderFrom 将指定流写入响应内
func (w *responseWrite) ReaderFrom(src io.Reader) (n int64, err error) {
	return w.ResponseWriter.(io.ReaderFrom).ReadFrom(src)
}

// NewResponseWrite 包装 http.ResponseWriter 获得额外的方法及数据
func NewResponseWrite(w http.ResponseWriter) ResponseWriter {
	return &responseWrite{ResponseWriter: w, status: defaultStatus}
}

// GetHTTPRespStatus 获取响应状态
func GetHTTPRespStatus(w http.ResponseWriter) int {
	if respW, ok := w.(ResponseWriter); ok {
		return respW.Status()
	}

	return http.StatusOK
}

// SetHTTPRespStatus 设置响应状态, w 应该是 responseWrite 实例, coverSupper 决定是否覆盖到底层的 http.ResponseWriter,默认 true
func SetHTTPRespStatus(w http.ResponseWriter, status int, coverSupper ... bool) http.ResponseWriter {

	cover := true
	if len(coverSupper) > 0 && !coverSupper[0] {
		cover = false
	}

	rw, ok := w.(ResponseWriter)

	if ! ok {
		w.WriteHeader(status)
		return w
	}

	if cover {
		rw.WriteHead(status)
	} else {
		rw.SetStatus(status)
	}
	return w
}

// ReaderFrom 将指定流写入响应内 , w 应该是 responseWrite 实例
func ReaderFrom(w http.ResponseWriter, src io.Reader) (n int64, err error) {
	return w.(ResponseWriter).ReaderFrom(src)
}
