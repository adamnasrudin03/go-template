package helpers

func ErrDB() *ResponseError {
	return NewError(ErrDatabase, NewResponseMultiLang(
		MultiLanguages{
			EN: "An error occurred while querying db",
			ID: "Terjadi kesalahan pada saat query db",
		}))
}
