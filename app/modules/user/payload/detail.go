package payload

import (
	"strings"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type DetailReq struct {
	ID       uint64 `json:"id"`
	NotID    uint64 `json:"not_id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Columns  string `json:"columns"`
}

func (m *DetailReq) Validate() error {
	m.Email = strings.TrimSpace(m.Email)
	m.Name = strings.TrimSpace(m.Name)
	m.Role = strings.TrimSpace(m.Role)
	m.Username = strings.TrimSpace(m.Username)

	isNotRequired := m.ID == 0 && m.NotID == 0 && m.Email == "" && m.Name == "" && m.Role == "" && m.Username == ""
	if isNotRequired {
		return helpers.NewError(helpers.ErrValidation, helpers.NewResponseMultiLang(
			helpers.MultiLanguages{
				ID: "Harap masukkan minimal satu parameter yang diperlukan",
				EN: "Please provide at least one required parameter",
			},
		))
	}

	return nil
}
