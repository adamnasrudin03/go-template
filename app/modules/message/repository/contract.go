package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/modules/message/dto"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MessageRepository interface {
	SendEmail(ctx context.Context, params dto.SendEmailReq) (err error)
}

type messageRepo struct {
	DB     *gorm.DB
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewMessageRepository(
	db *gorm.DB,
	cfg *configs.Configs,
	logger *logrus.Logger,
) MessageRepository {
	return &messageRepo{
		DB:     db,
		Cfg:    cfg,
		Logger: logger,
	}
}
