package service

import (
	"context"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	cacheRepo "github.com/adamnasrudin03/go-template/app/modules/cache/repository"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	messageRepo "github.com/adamnasrudin03/go-template/app/modules/message/repository"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/app/modules/user/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetDetail(ctx context.Context, input dto.DetailReq) (*models.User, error)
	ChangePassword(ctx context.Context, input dto.ChangePasswordReq) error
	Update(ctx context.Context, input dto.UpdateReq) (res *models.User, err error)
	GetList(ctx context.Context, params *dto.ListUserReq) (*response_mapper.Pagination, error)
	SendEmailVerify(ctx context.Context, userID uint64) (*dto.VerifyOtpRes, error)
	VerifiedEmail(ctx context.Context, req *dto.VerifyOtpReq) (err error)
	SendEmailResetPass(ctx context.Context, userID uint64) (*dto.VerifyOtpRes, error)
	ResetPassword(ctx context.Context, input *dto.ResetPasswordReq) (err error)
}

type UserSrv struct {
	Repo        repository.UserRepository
	RepoCache   cacheRepo.CacheRepository
	RepoLog     logRepo.LogRepository
	RepoMessage messageRepo.MessageRepository
	Cfg         *configs.Configs
	Logger      *logrus.Logger
}

func NewUserService(params UserSrv) UserService {
	return &params
}
