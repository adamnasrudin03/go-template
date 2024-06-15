package dto

import (
	"strings"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type ResetPasswordReq struct {
	ID              uint64 `json:"id"`
	RequestID       string `json:"request_id"`
	Otp             string `json:"otp"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
	UpdatedBy       uint64
}

func (m *ResetPasswordReq) Validate() error {
	if m.ID == 0 {
		return helpers.ErrIsRequired("ID Pengguna", "User ID")
	}
	if m.UpdatedBy == 0 {
		m.UpdatedBy = m.ID
	}

	m.RequestID = strings.TrimSpace(m.RequestID)
	if len(m.RequestID) == 0 {
		return helpers.ErrIsRequired("ID Permintaan", "Request ID")
	}
	if !helpers.IsValidUUID(m.RequestID) {
		return helpers.ErrInvalidFormat("ID Permintaan", "Request ID")
	}

	m.Otp = strings.TrimSpace(m.Otp)
	if len(m.Otp) == 0 {
		return helpers.ErrIsRequired("Otp", "Otp")
	}

	if m.NewPassword != m.ConfirmPassword {
		return helpers.ErrNewPasswordNotMatchWithConfirmPassword()
	}
	return nil
}
