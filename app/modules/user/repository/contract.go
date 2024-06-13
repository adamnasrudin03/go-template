package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateCache(ctx context.Context, key string, data interface{})
	GetCache(ctx context.Context, key string, res interface{})
	DelCache(ctx context.Context, key string) error
	Register(ctx context.Context, input models.User) (res *models.User, err error)
	Login(ctx context.Context, input dto.LoginReq) (res *models.User, er error)
	GetDetail(ctx context.Context, input dto.DetailReq) (res *models.User, err error)
	Updates(ctx context.Context, input models.User) (res *models.User, err error)
	UpdateSpecificField(ctx context.Context, input models.User) (err error)
	InsertLog(ctx context.Context, input models.Log) (err error)
	GetList(ctx context.Context, params dto.ListUserReq) (res []models.User, err error)
	CheckIsDuplicate(ctx context.Context, input dto.DetailReq) (err error)
}

type userRepo struct {
	DB     *gorm.DB
	Cache  driver.RedisClient
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewUserRepository(
	db *gorm.DB,
	cache driver.RedisClient,
	cfg *configs.Configs,
	logger *logrus.Logger,
) UserRepository {
	return &userRepo{
		DB:     db,
		Cache:  cache,
		Cfg:    cfg,
		Logger: logger,
	}
}
