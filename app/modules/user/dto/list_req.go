package dto

import (
	"strings"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
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
	m.Role = help.ToUpper(m.Role)
	m.SortBy = help.ToLower(m.SortBy)
	m.OrderBy = help.ToUpper(m.OrderBy)
	if len(m.OrderBy) > 0 && !models.IsValidOrderBy[m.OrderBy] {
		return response_mapper.ErrInvalid("order_by", "order_by")
	}

	m.BasedFilter.DefaultQuery()

	m.UserRole = help.ToUpper(m.UserRole)
	if m.UserRole != models.ROOT {
		m.NotIncRoleRoot = true
	}

	return nil
}
