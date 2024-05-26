package helpers

import "fmt"

func ErrIsRequired(id, en string) *ResponseError {
	errMsg := MultiLanguages{
		ID: fmt.Sprintf("%s harus diisi", id),
		EN: fmt.Sprintf("%s is required", en),
	}

	return NewError(ErrValidation, NewResponseMultiLang(errMsg))
}

func ErrIsEmpty(id, en string) *ResponseError {
	errMsg := MultiLanguages{
		ID: fmt.Sprintf("%s tidak ada isinya", id),
		EN: fmt.Sprintf("%s is empty", en),
	}

	return NewError(ErrValidation, NewResponseMultiLang(errMsg))
}
