package utils

import (
	"encoding/json"
	"fmt"
	"ginBlog/src/gin-blog/package/exception"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	//实例化需要转换的语言
	zh := zh.New()
	uni := ut.New(zh, zh)
	Trans, _ = uni.GetTranslator("zh")
	Validate = validator.New()
	//注册转换的语言为默认语言
	translations.RegisterDefaultTranslations(Validate, Trans)

}

//CommonError 错误格式
type CommonError []string

func NewValidatorError(err error) {
	res := CommonError{}
	fmt.Println(err)
	switch err.(type) {
	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			tranStr := e.Translate(Trans)
			//将结构体字段转换map中的key为小写
			f := strings.ToLower(tranStr)
			res = append(res, f)
		}
		str, _ := json.Marshal(res)
		panic(exception.BodyError(string(str), nil))
		return

	default:
		panic(exception.BodyError("请求参数异常", nil))
	}

}
