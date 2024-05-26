package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name" `
	Role     string `gorm:"not null;default:'USER'" json:"role"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" `
	Password string `gorm:"not null" json:"password,omitempty"`
	Salt     string `gorm:"not null" json:"salt,omitempty"`
	DefaultModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		u.Salt = fmt.Sprintf(`%v-%v`, time.Now().Unix(), u.Email)
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}

	u.Password = hashedPass
	if !IsUserRoleValid[u.Role] {
		err = errors.New("role is invalid, must be 'ADMIN' or 'USER'")
		return
	}

	return
}
