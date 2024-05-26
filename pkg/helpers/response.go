package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// StatusMapping maps HTTP status code to a descriptive string.
// The returned string can be used as the 'status' field in ResponseDefault.
func StatusMapping(statusCode int) string {
	mappings := map[int]string{
		http.StatusOK:                  "Success",
		http.StatusCreated:             "Created",
		http.StatusBadRequest:          "Bad Request",
		http.StatusConflict:            "Conflict",
		http.StatusUnauthorized:        "Unauthorized",
		http.StatusForbidden:           "Forbidden",
		http.StatusNotFound:            "Not Found",
		http.StatusInternalServerError: "Internal Server Error",
	}
	status := mappings[statusCode]
	if status == "" {
		status = mappings[StatusErrorMapping(statusCode)]
	}

	if status == "" {
		status = http.StatusText(statusCode)
	}
	return status
}

// APIResponse is for generating template responses
func APIResponse(message string, statusCode int, data interface{}) ResponseDefault {
	return ResponseDefault{
		Status:  StatusMapping(statusCode),
		Message: message,
		Data:    data,
	}
}

// FormatValidationError func which holds errors during user input validation
func FormatValidationError(err error) string {
	var errors string

	for _, e := range err.(validator.ValidationErrors) {
		if errors != "" {
			errors = fmt.Sprintf("%v, ", strings.TrimSpace(errors))
		}

		if e.Tag() == "email" {
			errors = errors + fmt.Sprintf("%v must be type %v", e.Field(), e.Tag())
		} else {
			errors = errors + fmt.Sprintf("%v is %v %v", e.Field(), e.Tag(), e.Param())
		}

		if e.Param() != "" && e.Type().Name() == "string" {
			errors = errors + " character"
		}

	}

	return strings.TrimSpace(errors) + "."
}
