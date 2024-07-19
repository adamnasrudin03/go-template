package dto

import (
	"strings"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
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
		return response_mapper.ErrIsRequired("ID Pengguna", "User ID")
	}
	if m.UpdatedBy == 0 {
		m.UpdatedBy = m.ID
	}

	m.RequestID = strings.TrimSpace(m.RequestID)
	if len(m.RequestID) == 0 {
		return response_mapper.ErrIsRequired("ID Permintaan", "Request ID")
	}
	if !help.IsUUID(m.RequestID) {
		return response_mapper.ErrInvalidFormat("ID Permintaan", "Request ID")
	}

	m.Otp = strings.TrimSpace(m.Otp)
	if len(m.Otp) == 0 {
		return response_mapper.ErrIsRequired("Otp", "Otp")
	}

	if m.NewPassword != m.ConfirmPassword {
		return response_mapper.ErrNewPasswordNotMatchWithConfirmPassword()
	}
	return nil
}
