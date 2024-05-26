package helpers

import "fmt"

func ErrMustBeMoreThanZero(id, en string) *ResponseError {
	errMsg := MultiLanguages{
		ID: fmt.Sprintf("%s harus lebih dari 0", id),
		EN: fmt.Sprintf("%s must be more than 0", en),
	}
	return NewError(ErrValidation, NewResponseMultiLang(errMsg))
}

func ErrCannotBeMoreThan(id, en, max string) *ResponseError {
	errMsg := MultiLanguages{
		ID: fmt.Sprintf("%s tidak boleh lebih dari %s", id, max),
		EN: fmt.Sprintf("%s cannot be more than %s", en, max),
	}
	return NewError(ErrValidation, NewResponseMultiLang(errMsg))
}

func ErrIsDuplicate(id, en string) *ResponseError {
	errMsg := MultiLanguages{
		ID: fmt.Sprintf("%s sudah ada", id),
		EN: fmt.Sprintf("%s already exists", en),
	}
	return NewError(ErrValidation, NewResponseMultiLang(errMsg))
}
