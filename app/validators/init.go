package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	Uni *ut.UniversalTranslator
)

func init() {
	// 	注册验证中文提示
	zh := zh.New()
	Uni = ut.New(zh, zh)
	ValidateTrans, _ := Uni.GetTranslator("zh")
	binding.Validator = new(DefaultValidator)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = zhTranslations.RegisterDefaultTranslations(v, ValidateTrans)
	}
}
