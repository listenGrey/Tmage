package controller

import (
	"Tmage/models"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func InitTranslator(locale string) (err error) {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// 为SignUpParam注册自定义校验方法
		validate.RegisterStructValidation(SignUpParamStructLevelValidation, models.RegisterFrom{})
		enT := en.New()
		zhT := zh.New()
		// 第一个是备用的语言，后面是应该支持的
		uni := ut.New(enT, zhT, enT)

		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		switch locale {
		case "en":
			err = en_translations.RegisterDefaultTranslations(validate, trans)
		case "zh":
			err = zh_translations.RegisterDefaultTranslations(validate, trans)
		default:
			err = en_translations.RegisterDefaultTranslations(validate, trans)
		}
		return
	}
	return
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.RegisterFrom)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "confirm_password", "ConfirmPassword", "eqfield", "password")
	}
}
