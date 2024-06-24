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
	return NewError(ErrForbidden, NewResponseMultiLang(MultiLanguages{
		ID: "Tidak ada akses untuk data ini",
		EN: "You don't have access to this data",
	}))
}

func ErrUnmarshalJSON() *ResponseError {
	return NewError(ErrUnknown, NewResponseMultiLang(MultiLanguages{
		ID: "Gagal membatalkan marshal JSON",
		EN: "Failed to unmarshal JSON",
	}))
}

func ErrFailedTranslateText() *ResponseError {
	return NewError(ErrUnknown, NewResponseMultiLang(MultiLanguages{
		ID: "Gagal menerjemahkan teks",
		EN: "Failed to translate text",
	}))
}

func ErrRouteNotFound() *ResponseError {
	return NewError(ErrNoFound, NewResponseMultiLang(MultiLanguages{
		ID: "Rute tidak ditemukan",
		EN: "Route not found",
	}))
}
func ErrGetRequest() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(
		MultiLanguages{
			ID: "Gagal membaca request data",
			EN: "Failed to parse data",
		}))
}

func ErrCannotUpdateData() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: "Tidak diperbolehkan mengubah data",
		EN: "Changing data is not allowed",
	}))
}
