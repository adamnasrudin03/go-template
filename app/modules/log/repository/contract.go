package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LogRepository interface {
	GetList(ctx context.Context, params dto.ListLogReq) (res []models.Log, err error)
	Create(ctx context.Context, input models.Log) (err error)
	CreateLogActivity(ctx context.Context, input models.Log) (err error)
}

type logRepo struct {
	DB     *gorm.DB
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewLogRepository(
	db *gorm.DB,
	cfg *configs.Configs,
	logger *logrus.Logger,
) LogRepository {
	return &logRepo{
		DB:     db,
		Cfg:    cfg,
		Logger: logger,
	}
}
