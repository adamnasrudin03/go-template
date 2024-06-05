package helpers

func ErrFailedTranslateText() *ResponseError {
	return NewError(ErrUnknown, NewResponseMultiLang(MultiLanguages{
		ID: "Gagal menerjemahkan teks",
		EN: "Failed to translate text",
	}))
}
