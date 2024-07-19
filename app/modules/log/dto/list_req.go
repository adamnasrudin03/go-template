package dto

import (
	"strings"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
)

type ListLogReq struct {
	UserID     uint64 `json:"user_id"`
	UsePreload bool
	models.BasedFilter
}

func (m *ListLogReq) Validate() error {
	if m.UserID == 0 {
		return response_mapper.ErrIsRequired("ID Pengguna", "User ID")
	}

	m.SortBy = strings.TrimSpace(strings.ToLower(m.SortBy))
	m.OrderBy = strings.TrimSpace(strings.ToUpper(m.OrderBy))
	if len(m.OrderBy) > 0 && !models.IsValidOrderBy[m.OrderBy] {
		return response_mapper.ErrInvalid("order_by", "order_by")
	}

	m.BasedFilter.DefaultQuery()
	m.UsePreload = true
	return nil
}
