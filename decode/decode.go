package decode

import (
	"net/url"

	"github.com/go-playground/form"
)

// Decoder 表单解析器，用于将 url.Values 内容解析到指定对象内
var Decoder = form.NewDecoder()

// DecodeForm 将指定 url.Values 值解析到指定对象 obj 内，obj 需要是对象指针
// 重复的 key ，值会被覆盖
func DecodeForm(obj interface{}, valuesArr ...url.Values) error {
	for _, values := range valuesArr {

		if values == nil {
			continue
		}

		// 重复对一个 obj 进行 decode
		// 前一次 decode 完成的 obj 不会被全量覆盖
		if err := Decoder.Decode(obj, values); err != nil {
			return err
		}
	}

	return nil
}