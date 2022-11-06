package translator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

var UT = ut.New(en.New(), ja.New(), zh.New())

func FromFiberCtx(ctx *fiber.Ctx) ut.Translator {
	tags, _, err := language.ParseAcceptLanguage(ctx.Get("Accept-Language"))
	if err != nil {
		return UT.GetFallback()
	}

	for _, tag := range tags {
		base, confidence := tag.Base()
		if confidence >= language.High {
			if t, ok := UT.GetTranslator(base.String()); ok {
				return t
			}
		}
	}

	return UT.GetFallback()
}
