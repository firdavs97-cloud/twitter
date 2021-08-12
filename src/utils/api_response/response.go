package api_response

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ErrorValidateResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func NewErrorResponse(c *fiber.Ctx, statusCode int, errors interface{}) error {
	return c.Status(statusCode).JSON(errors) // TODO
}

func MessageResponse(c *fiber.Ctx, message interface{}) error {
	return c.Status(fiber.StatusOK).JSON(message) // TODO
}

func ValidateStruct(value interface{}) []*ErrorValidateResponse {
	var errors []*ErrorValidateResponse
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidateResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
