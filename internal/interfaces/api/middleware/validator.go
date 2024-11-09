package middleware

import (
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// バリデーターの初期化
		validate := validator.New()

		// トランスレーターの設定
		japanese := ja.New()
		uni := ut.New(japanese, japanese)
		trans, _ := uni.GetTranslator("ja")

		// 日本語エラーメッセージの登録
		ja_translations.RegisterDefaultTranslations(validate, trans)

		// カスタムメッセージの登録
		messages := validators.GetValidationMessages()
		for tag, message := range messages {
			registerCustomTranslation(validate, trans, tag, message)
		}

		// コンテキストにバリデーターとトランスレーターを保存
		c.Set("validator", validate)
		c.Set("translator", trans)

		c.Next()
	}
}

func registerCustomTranslation(validate *validator.Validate, trans ut.Translator, tag string, message string) {
	validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}
