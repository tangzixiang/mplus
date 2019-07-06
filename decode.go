package mplus

import (
	"net/url"

	"github.com/go-playground/form"
)

// Decoder 表单解析器，用于将 url.Values 内容解析到指定对象内
var Decoder = form.NewDecoder()

// DecodeForm 将指定 url.Values 值解析到指定对象内
func DecodeForm(obj interface{}, valuesArr ...url.Values) error {
	for _, values := range valuesArr {

		if values == nil {
			continue
		}

		if err := Decoder.Decode(obj, values); err != nil {
			return err
		}
	}

	return nil
}
