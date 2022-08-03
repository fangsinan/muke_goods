package main

import (
	"fmt"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func main() {

	// en := en.New()
	zh := zh.New()
	uni = ut.New(zh)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("zh")

	validate = validator.New()
	// en_translations.RegisterDefaultTranslations(validate, trans)
	zh_translations.RegisterDefaultTranslations(validate, trans)

	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10"`
		Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "",
		Tagline:  "This tagline is way too long.",
		Tagline2: "2",
	}

	err := validate.Struct(user)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))
	}
}
