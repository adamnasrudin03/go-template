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
