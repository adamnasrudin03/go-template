package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
)

type ResponseError struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Err     error          `json:"-"`
	Message MultiLanguages `json:"message"`
}

type MultiLanguages struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

func (e *MultiLanguages) Error() string {
	if e.EN != "" {
		return e.EN
	} else if e.ID != "" {
		return e.ID
	}
	return "something went wrong"
}

func NewResponseMultiLang(languages MultiLanguages) *MultiLanguages {
	return &languages
}

func NewError(code TypeError, err error) *ResponseError {

	var respErr MultiLanguages
	if errValue, isMatch := err.(*MultiLanguages); isMatch {
		if errValue != nil {
			respErr = *errValue
		} else {
			respErr = MultiLanguages{
				ID: err.Error(),
				EN: err.Error(),
			}
		}
	} else {
		respErr = MultiLanguages{
			ID: err.Error(),
			EN: err.Error(),
		}
	}
	return &ResponseError{
		Status:  StatusMapping(int(code)),
		Code:    int(code),
		Err:     err,
		Message: respErr,
	}
}

func (e *ResponseError) Error() string {
	return e.Err.Error()
}

// FormatValidationError function is used to format validation errors that occur
// during user input validation. It changes the default format of errors from
// github.com/go-playground/validator/v10.
//
// This function iterates over the validation errors and constructs a formatted
// error message. The error message is then translated into the Indonesian language
// using the Translate function. The resulting error message is then returned as
// a ResponseError struct.
//
// Parameters:
// - err: The validation errors that occurred.
//
// Returns:
// - error: The formatted validation error.
func FormatValidationError(err error) error {
	// Initialize variables to store the translated and untranslated error messages.
	var (
		msgIdn  string
		msgEnUs string
	)

	// Iterate over the validation errors.
	for _, e := range err.(validator.ValidationErrors) {
		// Add a comma and space to the existing error message if it's not empty.
		if msgEnUs != "" {
			msgEnUs = fmt.Sprintf("%v, ", strings.TrimSpace(msgEnUs))
		}

		// Construct the error message based on the validation error.
		if e.Tag() == "email" {
			msgEnUs = msgEnUs + fmt.Sprintf("%v must be type %v", e.Field(), e.Tag())
		} else {
			msgEnUs = msgEnUs + fmt.Sprintf("%v is %v %v", e.Field(), e.Tag(), e.Param())
		}

		// If the parameter is not empty and the type is a string, add the word "character".
		if e.Param() != "" && e.Type().Name() == "string" {
			msgEnUs = msgEnUs + " character"
		}
	}

	// Add a period to the end of the error message and remove any leading or trailing spaces.
	msgEnUs = strings.TrimSpace(msgEnUs) + "."

	// Set the target language to Indonesian.
	langTo := language.Indonesian.String()

	// Translate the error message to Indonesian using the Translate function.
	msgIdn, errTranslate := Translate(msgEnUs, Auto, langTo)

	// If there is an error during translation, set the translated message to the untranslated message
	// and log the error.
	if errTranslate != nil {
		msgIdn = msgEnUs
		log.Printf("Translate Text %v to %v error: %v \n", Auto, langTo, errTranslate)
	}

	// Return the formatted validation error as a ResponseError struct.
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: msgIdn,
		EN: msgEnUs,
	}))
}
