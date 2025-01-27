package utils

import (
	"github.com/go-playground/validator/v10"
)

// parseValidationError memproses error validasi dan menghasilkan pesan error custom.
func ParseValidationError(err error) string {
	var errorMsg string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			field := vErr.Field() // Nama field yang gagal
			tag := vErr.Tag()     // Tag validasi yang gagal
			switch tag {
			case "required":
				errorMsg = field + " is required"
			default:
				errorMsg = field + " is invalid"
			}
		}
	} else {
		errorMsg = err.Error() // Error lain (non-validasi)
	}
	return errorMsg
}
