package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/app/modules/user/repository"
	"github.com/adamnasrudin03/go-template/pkg/driver"
)

type UserService interface {
	Register(ctx context.Context, input payload.RegisterReq) (res *models.User, err error)
	Login(ctx context.Context, input payload.LoginReq) (res *payload.LoginRes, err error)
	GetDetail(ctx context.Context, input payload.DetailReq) (*models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(
	userRepo repository.UserRepository,
	cache driver.RedisClient,
) UserService {
	return &userService{
		userRepository: userRepo,
	}
}
