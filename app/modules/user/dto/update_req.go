package dto

import (
	"strings"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
)

type UpdateReq struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	UpdatedBy uint64 `json:"updated_by"`
}

func (m *UpdateReq) Validate() error {
	if m.ID == 0 {
		return response_mapper.ErrIsRequired("ID Pengguna", "User ID")
	}

	m.Email = strings.TrimSpace(m.Email)
	if m.Email != "" && !help.IsEmail(m.Email) {
		return response_mapper.ErrInvalidFormat("Surel", "Email")
	}

	m.Role = help.ToUpper(m.Role)
	if m.Role != "" && !models.IsUserRoleValid[m.Role] {
		return response_mapper.ErrInvalidFormat("Peran", "Role")
	}

	if m.UpdatedBy == 0 {
		m.UpdatedBy = m.ID
	}
	return nil
}

func (m *UpdateReq) ConvertToUser() models.User {
	return models.User{
		ID:       m.ID,
		Name:     m.Name,
		Role:     help.ToUpper(m.Role),
		Email:    m.Email,
		Username: m.Username,
		DefaultModel: models.DefaultModel{
			UpdatedBy: m.UpdatedBy,
		},
	}
}
