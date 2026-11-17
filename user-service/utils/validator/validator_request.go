package validator

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Validator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidator() *Validator {
	en := en.New()
	uni := ut.New(en, en)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal().Msg("Translator not found")
	}

	validate := validator.New()

	return &Validator{
		Validator:  validate,
		Translator: trans,
	}
}

func (v *Validator) Validate(data interface{}) error {
	err := v.Validator.Struct(data)
	if err != nil {
		object, _ := err.(validator.ValidationErrors)
		for _, fieldError := range object {
			log.Info().Msgf("Validation Error: %s:%s", fieldError.Field(), fieldError.Translate(v.Translator))
			return errors.New(fieldError.Translate(v.Translator))
		}
	}
	return nil
}
