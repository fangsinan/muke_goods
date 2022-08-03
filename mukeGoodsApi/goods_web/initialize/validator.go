package initialize

import (
	"fmt"
	"reflect"
	"strings"

	"webApi/goods_web/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func InitValidator(locale string) error {
	// 修改 --gin框架--  的validator 引擎属性  实现定制
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册获取json tag 的自定义方式
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zh := zh.New()
		en := en.New()
		uni := ut.New(en, zh, en)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)

		// validate := validator.New()  // 无需再重新new validator  使用gin配置好的validator
		// en_translations.RegisterDefaultTranslations(validate, trans)
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(validate, global.Trans)
			break
		case "zh":
			zh_translations.RegisterDefaultTranslations(validate, global.Trans)
			break
		default:
			en_translations.RegisterDefaultTranslations(validate, global.Trans)
		}
	}
	return nil
}
