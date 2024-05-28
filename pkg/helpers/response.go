package helpers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
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
		statusCode = StatusErrorMapping(statusCode)
		status = mappings[statusCode]
	}

	status = strings.TrimSpace(status)
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
func FormatValidationError(err error) error {
	var (
		msgIdn  string
		msgEnUs string
	)

	for _, e := range err.(validator.ValidationErrors) {
		if msgEnUs != "" {
			msgEnUs = fmt.Sprintf("%v, ", strings.TrimSpace(msgEnUs))
		}

		if e.Tag() == "email" {
			msgEnUs = msgEnUs + fmt.Sprintf("%v must be type %v", e.Field(), e.Tag())
		} else {
			msgEnUs = msgEnUs + fmt.Sprintf("%v is %v %v", e.Field(), e.Tag(), e.Param())
		}

		if e.Param() != "" && e.Type().Name() == "string" {
			msgEnUs = msgEnUs + " character"
		}

	}

	msgEnUs = strings.TrimSpace(msgEnUs) + "."
	langTo := language.Indonesian.String()
	msgIdn, errTranslate := Translate(msgEnUs, Auto, langTo)
	if errTranslate != nil {
		msgIdn = msgEnUs
		log.Printf("Translate Text %v to %v error: %v \n", Auto, langTo, errTranslate)
	}

	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: msgIdn,
		EN: msgEnUs,
	}))
}
