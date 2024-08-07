package dto

import (
	"strings"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
)

type DetailReq struct {
	ID       uint64 `json:"id"`
	NotID    uint64 `json:"not_id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Columns  string `json:"columns"`
	UserID   uint64 `json:"user_id"` // user id for checking by token
}

func (m *DetailReq) Validate() error {
	m.Email = strings.TrimSpace(m.Email)
	m.Name = strings.TrimSpace(m.Name)
	m.Role = strings.TrimSpace(m.Role)
	m.Username = strings.TrimSpace(m.Username)

	isNotRequired := m.ID == 0 && m.NotID == 0 && m.Email == "" && m.Name == "" && m.Role == "" && m.Username == ""
	if isNotRequired {
		return response_mapper.NewError(response_mapper.ErrValidation, response_mapper.NewResponseMultiLang(
			response_mapper.MultiLanguages{
				ID: "Harap masukkan minimal satu parameter yang diperlukan",
				EN: "Please provide at least one required parameter",
			},
		))
	}

	return nil
}
