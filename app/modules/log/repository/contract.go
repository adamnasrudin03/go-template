package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/payload"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"gorm.io/gorm"
)

var config = configs.GetInstance()

type LogRepository interface {
	CreateCache(ctx context.Context, key string, data interface{})
	GetCache(ctx context.Context, key string, res interface{})
	DelCache(ctx context.Context, key string) error
	GetList(ctx context.Context, params payload.ListLogReq) (res []models.Log, err error)
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