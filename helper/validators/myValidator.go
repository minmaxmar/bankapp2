package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	expiryDateRegex = regexp.MustCompile(`^(0[1-9]|1[0-2])\/([0-9]{2})$`) // MM/YY format
)

func ValidateExpiryDate(fl validator.FieldLevel) bool {
	expiryDate := fl.Field().String()
	return expiryDateRegex.MatchString(expiryDate)
}
