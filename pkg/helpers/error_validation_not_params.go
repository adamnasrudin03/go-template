package helpers

func ErrGetRequest() *ResponseError {
	return NewError(ErrDatabase, NewResponseMultiLang(
		MultiLanguages{
			ID: "Gagal membaca request data",
			EN: "Failed to parse data",
		}))
}

func ErrCannotUpdateData() *ResponseError {
	errMsg := MultiLanguages{
		ID: "Tidak diperbolehkan mengubah data",
		EN: "Cannot update data",
	}
	return NewError(ErrValidation, NewResponseMultiLang(errMsg))
}
