package dto

import (
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
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
		return helpers.ErrIsRequired("ID Pengguna", "User ID")
	}

	m.Email = strings.TrimSpace(m.Email)
	if m.Email != "" && !helpers.IsValidEmail(m.Email) {
		return helpers.ErrInvalidFormat("Surel", "Email")
	}

	m.Role = helpers.ToUpper(m.Role)
	if m.Role != "" && !models.IsUserRoleValid[m.Role] {
		return helpers.ErrInvalidFormat("Peran", "Role")
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
		Role:     helpers.ToUpper(m.Role),
		Email:    m.Email,
		Username: m.Username,
		DefaultModel: models.DefaultModel{
			UpdatedBy: m.UpdatedBy,
		},
	}
}
