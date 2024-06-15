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

func ErrOtpExpired() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: "Kode OTP sudah kedaluwarsa",
		EN: "OTP code has expired",
	}))
}

func ErrGenerateOtp() *ResponseError {
	return NewError(ErrUnknown, NewResponseMultiLang(MultiLanguages{
		ID: "Gagal membuat kode OTP",
		EN: "Failed to generate OTP code",
	}))
}

func ErrOtpInvalid() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: "Kode OTP tidak valid",
		EN: "Invalid OTP code",
	}))
}

func ErrEmailIsVerified() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: "Email sudah terverifikasi",
		EN: "Email is already verified",
	}))
}

func ErrEmailNotVerified() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: "Email belum terverifikasi",
		EN: "Email has not been verified",
	}))
}
