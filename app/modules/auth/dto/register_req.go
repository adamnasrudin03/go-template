package dto

import (
	"strings"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
)

type RegisterReq struct {
	Name      string `json:"name" validate:"required"`
	Role      string `json:"role"  validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=4"`
	Password  string `json:"password" validate:"required,min=4"`
	CreatedBy uint64 `json:"created_by"`
	UpdatedBy uint64 `json:"updated_by"`
}

func (m *RegisterReq) Validate() error {
	if len(strings.TrimSpace(m.Name)) == 0 {
		return response_mapper.ErrIsRequired("Nama", "Name")
	}

	m.Role = strings.TrimSpace(m.Role)
	if len(m.Role) == 0 {
		return response_mapper.ErrIsRequired("Peran", "Role")
	}
	if !models.IsUserRoleValid[m.Role] {
		return response_mapper.ErrInvalidFormat("Peran", "Role")
	}

	m.Email = strings.TrimSpace(m.Email)
	if len(m.Email) == 0 {
		return response_mapper.ErrIsRequired("Surel", "Email")
	}
	if !help.IsEmail(m.Email) {
		return response_mapper.ErrInvalidFormat("Surel", "Email")
	}

	m.Password = strings.TrimSpace(m.Password)
	if len(m.Password) < 4 {
		return response_mapper.ErrMinCharacters("Kata Sandi", "Password", "4")
	}

	m.Username = strings.TrimSpace(m.Username)
	if len(m.Username) < 4 {
		return response_mapper.ErrMinCharacters("Nama pengguna", "username", "4")
	}

	return nil
}
