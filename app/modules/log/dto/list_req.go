package dto

import (
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type ListLogReq struct {
	UserID     uint64 `json:"user_id"`
	UsePreload bool
	models.BasedFilter
}

func (m *ListLogReq) Validate() error {
	if m.UserID == 0 {
		return helpers.ErrIsRequired("ID Pengguna", "User ID")
	}

	m.SortBy = strings.TrimSpace(strings.ToLower(m.SortBy))
	m.OrderBy = strings.TrimSpace(strings.ToUpper(m.OrderBy))
	if len(m.OrderBy) > 0 && !models.IsValidOrderBy[m.OrderBy] {
		return helpers.ErrInvalid("order_by", "order_by")
	}

	m.BasedFilter.DefaultQuery()
	m.UsePreload = true
	return nil
}
