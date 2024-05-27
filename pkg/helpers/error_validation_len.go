package helpers

import "fmt"

func ErrTooShort(id, en string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: fmt.Sprintf("%s terlalu pendek", id),
		EN: fmt.Sprintf("%s is too short", en),
	}))
}

func ErrTooLong(id, en string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: fmt.Sprintf("%s terlalu panjang", id),
		EN: fmt.Sprintf("%s is too long", en),
	}))
}

func ErrTooMany(id, en string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: fmt.Sprintf("%s terlalu banyak", id),
		EN: fmt.Sprintf("%s is too many", en),
	}))
}

func ErrMinCharacters(id, en, min string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(
		MultiLanguages{
			ID: fmt.Sprintf("%s minimal %s karakter", id, min),
			EN: fmt.Sprintf("%s must be at least %s characters", en, min),
		},
	))
}

func ErrMaxCharacters(id, en, max string) *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: fmt.Sprintf("Isi %s maksimal %s karakter", id, max),
		EN: fmt.Sprintf("%s must be at most %s characters", en, max),
	}))
}
