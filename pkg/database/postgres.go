package database

import (
	"fmt"
	"log"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/seeders"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

// SetupDbConnection is creating a new connection to our database
func SetupDbConnection() *gorm.DB {
	configs := configs.GetInstance()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.DB.Host,
		configs.DB.Username,
		configs.DB.Password,
		configs.DB.DbName,
		configs.DB.Port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	if configs.DB.DebugMode {
		if configs.App.Env == "dev" {
			db = db.Debug()
		}
	}

	if configs.DB.DbIsMigrate {
		//auto migration entity db
		db.AutoMigrate(
			&models.User{},
			&models.Log{},
		)
	}

	go func(db *gorm.DB) {
		seeders.InitUser(db)
	}(db)

	log.Println("Connection Database Success!")
	return db
}

// CloseDbConnection method is closing a connection between your app and your db
func CloseDbConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}

func GetDB() *gorm.DB {
	return db
}
