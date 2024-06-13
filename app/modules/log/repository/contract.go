package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LogRepository interface {
	CreateCache(ctx context.Context, key string, data interface{})
	GetCache(ctx context.Context, key string, res interface{})
	DelCache(ctx context.Context, key string) error
	GetList(ctx context.Context, params dto.ListLogReq) (res []models.Log, err error)
	Create(ctx context.Context, input models.Log) (err error)
	CreateLog(ctx context.Context, input models.Log) (err error)
}

type logRepo struct {
	DB     *gorm.DB
	Cache  driver.RedisClient
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewLogRepository(
	db *gorm.DB,
	cache driver.RedisClient,
	cfg *configs.Configs,
	logger *logrus.Logger,
) LogRepository {
	return &logRepo{
		DB:     db,
		Cache:  cache,
		Cfg:    cfg,
		Logger: logger,
	}
}
