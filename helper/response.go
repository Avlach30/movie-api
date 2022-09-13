package helper

import "github.com/go-playground/validator/v10"

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

	return res
}

func ErrorValidationResponse(err error) []string {
	var errors []string
	for _, error := range err.(validator.ValidationErrors) {
		errors = append(errors, error.Error())
	}

	return errors
}
