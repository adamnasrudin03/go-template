package dto

type VerifyOtpRes struct {
	RequestID string `json:"request_id"`
	Otp       string `json:"-"`
}
