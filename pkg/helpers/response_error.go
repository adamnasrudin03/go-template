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
