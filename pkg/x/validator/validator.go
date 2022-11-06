package validator

import (
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	jaTranslations "github.com/go-playground/validator/v10/translations/ja"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/rs/zerolog/log"

	"github.com/GalvinGao/gofiber-template/pkg/x/translator"
)

var validate *validator.Validate = validator.New()

func init() {
	var err error
	entr, _ := translator.UT.GetTranslator("en")
	err = enTranslations.RegisterDefaultTranslations(validate, entr)
	if err != nil {
		log.Warn().Err(err).Str("locale", "en").Msg("could not register translation")
	}

	zhtr, _ := translator.UT.GetTranslator("zh")
	err = zhTranslations.RegisterDefaultTranslations(validate, zhtr)
	if err != nil {
		log.Warn().Err(err).Str("locale", "zh").Msg("could not register translation")
	}

	jatr, _ := translator.UT.GetTranslator("ja")
	err = jaTranslations.RegisterDefaultTranslations(validate, jatr)
	if err != nil {
		log.Warn().Err(err).Str("locale", "ja").Msg("could not register translation")
	}
}

func Struct(s interface{}) error {
	return validate.Struct(s)
}
