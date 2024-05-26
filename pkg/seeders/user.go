package seeders

import (
	"github.com/adamnasrudin03/go-template/app/models"

	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) {
	tx := db.Begin()
	var users = []models.User{}
	tx.Select("id").Where("role = ? ", models.ROOT).Find(&users)
	if len(users) == 0 {
		user := models.User{
			Name:     "Super Admin",
			Password: "password123",
			Email:    "superadmin@gmail.com",
			Role:     models.ROOT,
		}
		tx.Create(&user)
	}

	tx.Commit()
}
