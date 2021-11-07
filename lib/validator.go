package lib

import (
	"callme/models"
	"github.com/go-playground/validator/v10"
)

func ValidateRegister(user models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.StructPartial(user, "Password", "Email", "Username")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Status = 403
			element.Description = "validation failed in condition: " + err.Tag()
			if err.Param() != "" {
				element.Description = "validation failed condition: " + err.Tag() + " = " + err.Param()
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
