package helpers

import "fmt"

func ErrInvalid(id, en string) *ResponseError {
	// Construct the error message using the provided arguments.
	errMsg := MultiLanguages{
		ID: fmt.Sprintf("%s tidak valid", id),
		EN: fmt.Sprintf("Invalid %s", en),
	}

	// Create and return the error.
	return NewError(ErrValidation, NewResponseMultiLang(errMsg))
}

func ErrInvalidFormat(id, en string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: fmt.Sprintf("Format %s tidak valid", id),
		EN: fmt.Sprintf("Invalid %s format", en),
	}))
}
