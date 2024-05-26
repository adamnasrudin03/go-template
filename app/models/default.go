package models

import "time"

type DefaultModel struct {
	CreatedBy uint64    `json:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedBy uint64    `json:"updated_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
