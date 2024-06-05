package models

import (
	"log"
	"time"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"gorm.io/gorm"
)

const (
	Updated = `Updated`
	Deleted = `Deleted`
	Created = `Created`
	Read    = `Read`
)

// Log represents the model for an log
type Log struct {
	ID          uint64       `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"not null"`
	Action      string       `json:"action" gorm:"not null"`
	TableNameID uint64       `json:"table_name_id"`
	TableName   string       `json:"table_name"`
	UserID      uint64       `json:"user_id"`
	LogDateTime time.Time    `json:"log_date_time"`
	User        UserRelation `json:"user" gorm:"ForeignKey:UserID"`
	DefaultModel
}

type UserRelation struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

func (UserRelation) TableName() string {
	return "users"
}

func (m *Log) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if m.LogDateTime.IsZero() {
		m.LogDateTime = now
	}
	if m.DefaultModel.CreatedAt.IsZero() {
		m.DefaultModel.CreatedAt = now
	}
	if m.DefaultModel.UpdatedAt.IsZero() {
		m.DefaultModel.UpdatedAt = now
	}

	return
}

func (m *Log) ToString() string {
	temp, err := helpers.SafeJsonMarshal(m)
	if err != nil {
		log.Printf("models.Log.ToString() error: %v ", err)
		return ""
	}

	return string(temp)
}
