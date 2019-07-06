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

	ReaderFrom(src io.Reader) (n int64, err error)
}

var (
	_ http.ResponseWriter = &responseWrite{}
	_ ResponseWriter      = &responseWrite{}
)

func (w *responseWrite) SetStatus(status int) {
	w.WriteHeader(status)
}

func (w *responseWrite) Status() int {
	return w.status
}

// WriteHeader 设置响应头
func (w *responseWrite) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
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

// SetHTTPRespStatus 设置响应状态, w 应该是 responseWrite 实例
func SetHTTPRespStatus(w http.ResponseWriter, status int) http.ResponseWriter {
	w.WriteHeader(status)
	return w
}

// ReaderFrom 将指定流写入响应内 , w 应该是 responseWrite 实例
func ReaderFrom(w http.ResponseWriter, src io.Reader) (n int64, err error) {
	return w.(ResponseWriter).ReaderFrom(src)
}
