package payload

import (
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type ListReq struct {
	UserID uint64 `json:"user_id"`
	models.DefaultModel
}

func (m *ListReq) Validate() error {
	if m.UserID == 0 {
		return helpers.ErrIsRequired("ID Pengguna", "User ID")
	}

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
