package helper

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/go-playground/validator/v10"
)

type ResponseSuccess struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFailed struct {
	Success bool        `json:"success"`
	Errors  interface{} `json:"error"`
	Message string      `json:"message"`
}

func ApiSuccessResponse(message string, data any) ResponseSuccess {
	res := ResponseSuccess{
		Success: true,
		Message: message,
		Data:    data,
	}

	return res
}

func ApiFailedResponse(err any) ResponseFailed {
	res := ResponseFailed{
		Success: false,
		Errors:  err,
		Message: "Error has been occured!",
	}

	error := fmt.Sprintf("%s", err)

	sentry.CaptureException(errors.New(error))

	return res
}

func ErrorValidationResponse(err error) []string {
	var errors []string
	for _, error := range err.(validator.ValidationErrors) {
		errorValidationMsg := fmt.Sprintf("Error on field %s, condition: %s %s", error.Field(), error.ActualTag(), error.Param())

		errors = append(errors, errorValidationMsg)
	}

	return errors
}
