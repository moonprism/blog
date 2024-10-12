package core

import (
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	zhtrans "github.com/go-playground/validator/v10/translations/zh"
)

type Validator struct {
	valid *validator.Validate
	trans ut.Translator
}

func NewValidator() *Validator {
	en := en.New()
	zh := zh.New()

	uni := ut.New(en, zh)
	// 报错信息的中文翻译器
	trans, _ := uni.GetTranslator("zh")

	valid := validator.New()
	zhtrans.RegisterDefaultTranslations(valid, trans)

	return &Validator{
		valid,
		trans,
	}
}

func (v *Validator) Struct(data any) map[string]string {
	err := v.valid.Struct(data)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors).Translate(v.trans)
	result := map[string]string{}
	for k, v := range errs {
		key := k[strings.Index(k, ".")+1:]
		result[key] = strings.TrimPrefix(v, key)
	}
	return result
}
