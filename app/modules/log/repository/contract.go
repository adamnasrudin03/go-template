package repository

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"gorm.io/gorm"
)

var config = configs.GetInstance()

type LogRepository interface {
}

type logRepo struct {
	DB    *gorm.DB
	Cache driver.RedisClient
}

func NewLogRepository(
	db *gorm.DB,
	cache driver.RedisClient) LogRepository {
	return &logRepo{
		DB:    db,
		Cache: cache,
	}
}
