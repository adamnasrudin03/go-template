package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	ID              uint64     `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name" gorm:"not null"`
	Role            string     `json:"role" gorm:"not null;default:'USER'"`
	Username        string     `json:"username" gorm:"not null;uniqueIndex"`
	Email           string     `json:"email" gorm:"not null;uniqueIndex"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"  gorm:"null; default:null"`
	Password        string     `json:"password,omitempty" gorm:"not null"`
	Salt            string     `json:"salt,omitempty" gorm:"not null"`
	DefaultModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		u.Salt = fmt.Sprintf(`%v-%v`, time.Now().Unix(), u.Email)
	}

	hashedPass, err := help.HashPassword(u.Password)
	if err != nil {
		log.Printf("failed hash password: %v ", err)
		return response_mapper.ErrHashPasswordFailed()
	}

	u.Password = hashedPass
	if u.Role != ROOT && !IsUserRoleValid[u.Role] {
		err = errors.New("role is invalid, must be 'ADMIN' or 'USER'")
		return
	}

	return
}

func (u *User) ConvertToResponse() {
	u.Password = ""
	u.Salt = ""
	u.Role = strings.ReplaceAll(help.ToLower(u.Role), "_", " ")
}
