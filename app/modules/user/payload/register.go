package payload

import (
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type RegisterReq struct {
	Name      string `json:"name" validate:"required"`
	Role      string `json:"role"  validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=4"`
	Password  string `json:"password" validate:"required,min=4"`
	CreatedBy uint64 `json:"created_by"`
}

func (m *RegisterReq) Validate() error {
	if len(strings.TrimSpace(m.Name)) == 0 {
		return helpers.ErrIsRequired("Nama", "Name")
	}

	m.Role = strings.TrimSpace(m.Role)
	if len(m.Role) == 0 {
		return helpers.ErrIsRequired("Peran", "Role")
	}
	if !models.IsUserRoleValid[m.Role] {
		return helpers.ErrInvalidFormat("Peran", "Role")
	}

	m.Email = strings.TrimSpace(m.Email)
	if len(m.Email) == 0 {
		return helpers.ErrIsRequired("Surel", "Email")
	}
	if !helpers.IsValidEmail(m.Email) {
		return helpers.ErrInvalidFormat("Surel", "Email")
	}

	m.Password = strings.TrimSpace(m.Password)
	if len(m.Password) < 4 {
		return helpers.ErrMinCharacters("Kata Sandi", "Password", "4")
	}

	m.Username = strings.TrimSpace(m.Username)
	if len(m.Username) < 4 {
		return helpers.ErrMinCharacters("Nama pengguna", "Username", "4")
	}

	return nil
}
