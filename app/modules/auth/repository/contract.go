package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(ctx context.Context, input models.User) (res *models.User, err error)
	Login(ctx context.Context, input dto.LoginReq) (res *models.User, er error)
}

type AuthRepo struct {
	DB     *gorm.DB
	Cache  driver.RedisClient
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewAuthRepository(
	params AuthRepo,
) AuthRepository {
	return &params
}
