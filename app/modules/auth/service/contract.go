package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	"github.com/adamnasrudin03/go-template/app/modules/auth/repository"
	cacheRepo "github.com/adamnasrudin03/go-template/app/modules/cache/repository"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	userRepo "github.com/adamnasrudin03/go-template/app/modules/user/repository"
	"github.com/sirupsen/logrus"
)

type AuthService interface {
	Register(ctx context.Context, input dto.RegisterReq) (res *models.User, err error)
	Login(ctx context.Context, input dto.LoginReq) (res *dto.LoginRes, err error)
}

type AuthSrv struct {
	Repo      repository.AuthRepository
	RepoCache cacheRepo.CacheRepository
	RepoUser  userRepo.UserRepository
	RepoLog   logRepo.LogRepository
	Cfg       *configs.Configs
	Logger    *logrus.Logger
}

func NewAuthService(
	params AuthSrv,
) AuthService {
	return &params
}
