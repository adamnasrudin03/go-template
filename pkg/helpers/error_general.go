package helpers

import "fmt"

func ErrReadContext() *ResponseError {
	return NewError(ErrForbidden, NewResponseMultiLang(
		MultiLanguages{
			ID: "Gagal membaca data konteks",
			EN: "Failed to read context data",
		}))
}

func ErrNotFound() *ResponseError {
	return NewError(ErrNoFound, NewResponseMultiLang(
		MultiLanguages{
			ID: "Data tidak ditemukan",
			EN: "Data not found",
		}))
}

func ErrDataNotFound(id, en string) *ResponseError {
	return NewError(ErrNoFound, NewResponseMultiLang(
		MultiLanguages{
			ID: fmt.Sprintf("Data %s tidak ditemukan", id),
			EN: fmt.Sprintf("Data %s not found", en),
		}))
}

func ErrNotAccess() *ResponseError {
	errMsg := MultiLanguages{
		ID: "Tidak ada akses untuk data ini",
		EN: "You don't have access to this data",
	}
	return NewError(ErrForbidden, NewResponseMultiLang(errMsg))
}

func ErrUnmarshalJSON() *ResponseError {
	return NewError(ErrUnknown, NewResponseMultiLang(MultiLanguages{
		ID: "Gagal membatalkan marshal JSON",
		EN: "Failed to unmarshal JSON",
	}))
}
