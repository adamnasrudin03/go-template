package service

import (
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/app/modules/user/repository"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(ctx *gin.Context, input payload.RegisterReq) (res models.User, err error)
	Login(ctx *gin.Context, input payload.LoginReq) (res payload.LoginRes, err error)
}

type userService struct {
	userRepository repository.UserRepository
	cache          driver.RedisClient
}

func NewUserService(
	userRepo repository.UserRepository,
	cache driver.RedisClient,
) UserService {
	return &userService{
		userRepository: userRepo,
		cache:          cache,
	}
}
