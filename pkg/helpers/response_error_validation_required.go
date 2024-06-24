package helpers

import "fmt"

func ErrIsRequired(id, en string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: fmt.Sprintf("%s harus diisi", id),
		EN: fmt.Sprintf("%s is required", en),
	}))
}

func ErrIsEmpty(id, en string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: fmt.Sprintf("%s tidak ada isinya", id),
		EN: fmt.Sprintf("%s is empty", en),
	}))
}
