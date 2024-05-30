package payload

import (
	"strings"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type ListLogReq struct {
	UserID uint64 `json:"user_id"`
	models.BasedFilter
}

func (m *ListLogReq) Validate() error {
	if m.UserID == 0 {
		return helpers.ErrIsRequired("ID Pengguna", "User ID")
	}

	m.SortBy = strings.TrimSpace(strings.ToLower(m.SortBy))
	m.OrderBy = strings.TrimSpace(strings.ToUpper(m.OrderBy))
	if len(m.OrderBy) > 0 && !models.IsValidOrderBy[m.OrderBy] {
		return helpers.ErrInvalid("sort_by", "sort_by")
	}

	m.BasedFilter.DefaultQuery()
	return nil
}

type LogRes struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Action    string    `json:"action"`
	UserID    uint64    `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
