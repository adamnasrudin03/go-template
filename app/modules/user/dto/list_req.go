package dto

import (
	"strings"

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
		return helpers.ErrInvalid("order_by", "order_by")
	}

	m.BasedFilter.DefaultQuery()

	m.UserRole = helpers.ToUpper(m.UserRole)
	if m.UserRole != models.ROOT {
		m.NotIncRoleRoot = true
	}

	return nil
}
