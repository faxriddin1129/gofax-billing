package requests

import (
	"github.com/go-playground/validator/v10"
	"log"
	"microservice/internal/constants"
)

var validate = validator.New()

func ValidateProvider(fl validator.FieldLevel) bool {
	provider := fl.Field().String()
	for _, p := range constants.ValidProviders {
		if p == provider {
			return true
		}
	}
	return false
}

func ValidateCurrency(fl validator.FieldLevel) bool {
	currency := fl.Field().String()
	for _, p := range constants.ValidCurrencies {
		if p == currency {
			return true
		}
	}
	return false
}

func init() {

	errPro := validate.RegisterValidation("provider", ValidateProvider)
	if errPro != nil {
		log.Fatalf("Validator registration failed: %v", errPro)
	}

	errCur := validate.RegisterValidation("currency", ValidateCurrency)
	if errCur != nil {
		log.Fatalf("Validator registration failed: %v", errCur)
	}
}
