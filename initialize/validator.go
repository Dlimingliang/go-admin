package initialize

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtTranslations "github.com/go-playground/validator/v10/translations/zh"

	"github.com/Dlimingliang/go-admin/global"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册自定义validator及其翻译器

		//注册validator默认翻译器
		initValidatorTrans(v)
	}

}

func initValidatorTrans(v *validator.Validate) {
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New()
	uni := ut.New(zhT)
	global.ValidatorTrans, _ = uni.GetTranslator("zh")
	if err := zhtTranslations.RegisterDefaultTranslations(v, global.ValidatorTrans); err != nil {
		global.GaLog.Error("注册默认validator翻译器失败")
	}
}
