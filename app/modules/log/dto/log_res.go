package dto

import (
	"time"
)

type LogRes struct {
	ID        uint64    `json:"id"`
	DateTime  time.Time `json:"date_time"`
	Name      string    `json:"name"`
	Action    string    `json:"action"`
	UserID    uint64    `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
