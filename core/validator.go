package core

import (
	"net/url"
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
	// 注册自定义验证函数
	valid.RegisterValidation("custom_url", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		// 如果没有协议（如 http:// 或 https://），加上 https://
		if !strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://") {
			value = "https://" + value
		}
		// 尝试解析 URL
		_, err := url.ParseRequestURI(value)
		return err == nil
	})
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
