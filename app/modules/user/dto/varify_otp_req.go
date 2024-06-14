package dto

import (
	"strings"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type VerifyOtpReq struct {
	UserID    uint64
	RequestID string `json:"request_id"`
	Otp       string `json:"otp"`
}

func (m *VerifyOtpReq) Validate() (err error) {
	if m.UserID == 0 {
		return helpers.ErrIsRequired("ID Pengguna", "User ID")
	}

	m.RequestID = strings.TrimSpace(m.RequestID)
	if len(m.RequestID) == 0 {
		return helpers.ErrIsRequired("ID Permintaan", "Request ID")
	}

	m.Otp = strings.TrimSpace(m.Otp)
	if len(m.Otp) == 0 {
		return helpers.ErrIsRequired("Otp", "Otp")
	}

	return nil
}
