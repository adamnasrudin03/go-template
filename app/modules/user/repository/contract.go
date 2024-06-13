package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetDetail(ctx context.Context, input dto.DetailReq) (res *models.User, err error)
	Updates(ctx context.Context, input models.User) (res *models.User, err error)
	UpdateSpecificField(ctx context.Context, input models.User) (err error)
	GetList(ctx context.Context, params dto.ListUserReq) (res []models.User, err error)
	CheckIsDuplicate(ctx context.Context, input dto.DetailReq) (err error)
}

type userRepo struct {
	DB     *gorm.DB
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewUserRepository(
	db *gorm.DB,
	cfg *configs.Configs,
	logger *logrus.Logger,
) UserRepository {
	return &userRepo{
		DB:     db,
		Cfg:    cfg,
		Logger: logger,
	}
}
