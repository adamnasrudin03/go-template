package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/app/modules/user/repository"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetDetail(ctx context.Context, input dto.DetailReq) (*models.User, error)
	ChangePassword(ctx context.Context, input dto.ChangePasswordReq) error
	Update(ctx context.Context, input dto.UpdateReq) (res *models.User, err error)
	GetList(ctx context.Context, params *dto.ListUserReq) (*helpers.Pagination, error)
}

type userService struct {
	userRepository repository.UserRepository
	Cfg            *configs.Configs
	Logger         *logrus.Logger
}

func NewUserService(
	userRepo repository.UserRepository,
	cfg *configs.Configs,
	logger *logrus.Logger,
) UserService {
	return &userService{
		userRepository: userRepo,
		Cfg:            cfg,
		Logger:         logger,
	}
}
