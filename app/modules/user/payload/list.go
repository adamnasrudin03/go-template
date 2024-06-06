package payload

import (
	"strings"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type ListUserReq struct {
	Search         string `json:"search" form:"search"`
	Role           string `json:"role" form:"role"`
	UserRole       string
	NotIncRoleRoot bool
	models.BasedFilter
}

func (m *ListUserReq) Validate() error {
	m.Search = strings.TrimSpace(m.Search)
	m.Role = helpers.ToUpper(m.Role)
	m.SortBy = helpers.ToLower(m.SortBy)
	m.OrderBy = helpers.ToUpper(m.OrderBy)
	if len(m.OrderBy) > 0 && !models.IsValidOrderBy[m.OrderBy] {
		return helpers.ErrInvalid("sort_by", "sort_by")
	}

	m.BasedFilter.DefaultQuery()

	m.UserRole = helpers.ToUpper(m.UserRole)
	if m.UserRole != models.ROOT {
		m.NotIncRoleRoot = true
	}

	return nil
}

type UserRes struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
