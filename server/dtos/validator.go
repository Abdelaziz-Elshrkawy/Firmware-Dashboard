package dtos

import "gopkg.in/go-playground/validator.v9"

var DtoVlidator *validator.Validate

func InitValidator() {
	DtoVlidator = validator.New()
}

func FormatValidationErrors(err error) []string {
	var errs []string

	if err == nil {
		return errs
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			// fieldErr := map[string]string{
			// 	"field":   e.Field(),
			// 	"tag":     e.Tag(),
			// 	"message": getErrorMessage(e),
			// }
			errs = append(errs, getErrorMessage(e))
		}
	}

	return errs
}

// Custom message mapping (optional)
func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "min":
		return e.Field() + " must be at least " + e.Param()
	case "max":
		return e.Field() + " cannot be more than " + e.Param()
	default:
		return e.Type().ChanDir().String()
	}
}
