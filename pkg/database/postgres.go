package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/adamnasrudin03/go-template/pkg/seeders"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	db     *gorm.DB
	err    error
	cfg    = configs.GetInstance()
	logger = driver.Logger(cfg)
)

// SetupDbConnection is creating a new connection to our database
func SetupDbConnection() *gorm.DB {
	logLevel := gormLogger.Silent
	if cfg.App.Env == "dev" {
		logLevel = gormLogger.Info
	}

	dbLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      true,        // Disable color
		},
	)
	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.DbName,
		cfg.DB.Port)

	db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		logger.Panicf("Failed to create a connection to database , %v", err)
		return nil
	}

	if cfg.DB.DbIsMigrate {
		//auto migration entity db
		db.AutoMigrate(
			&models.User{},
			&models.Log{},
		)
	}

	go func(db *gorm.DB) {
		seeders.InitUser(db)
	}(db)

	logger.Info("Connection Database Success!")
	return db
}

// CloseDbConnection method is closing a connection between your app and your db
func CloseDbConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		logger.Panicf("Failed to close connection from database, %v", err)
		return
	}

	dbSQL.Close()
}

func GetDB() *gorm.DB {
	return db
}
