package helpers

func ErrPasswordNotMatch() error {
	return NewError(ErrValidation, NewResponseMultiLang(
		MultiLanguages{
			ID: "Kata sandi tidak sesuai",
			EN: "Password not match",
		},
	))
}

func ErrNewPasswordNotMatchWithConfirmPassword() error {
	return NewError(ErrValidation, NewResponseMultiLang(
		MultiLanguages{
			ID: "Kata sandi baru tidak sesuai dengan kata sandi konfirmasi",
			EN: "New password not match with confirmation password",
		},
	))
}

func ErrHashPasswordFailed() error {
	return NewError(ErrFromUseCase, NewResponseMultiLang(
		MultiLanguages{
			ID: "Gagal hash kata sandi",
			EN: "Failed to hash password",
		},
	))
}

func ErrCannotHaveAccessUpdateData() *ResponseError {
	return NewError(ErrForbidden, NewResponseMultiLang(MultiLanguages{
		ID: "Tidak memiliki akses untuk mengubah data",
		EN: "Does not have access to change data",
	}))
}

func ErrCannotHaveAccessResources() *ResponseError {
	return NewError(ErrForbidden, NewResponseMultiLang(MultiLanguages{
		ID: "Anda tidak diizinkan untuk mengakses sumber daya ini",
		EN: "You are not allowed to access this resources",
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
