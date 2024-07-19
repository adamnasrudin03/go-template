package dto

import (
	"strings"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
)

type VerifyOtpReq struct {
	UserID    uint64
	RequestID string `json:"request_id"`
	Otp       string `json:"otp"`
}

func (m *VerifyOtpReq) Validate() (err error) {
	if m.UserID == 0 {
		return response_mapper.ErrIsRequired("ID Pengguna", "User ID")
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

	return nil
}
