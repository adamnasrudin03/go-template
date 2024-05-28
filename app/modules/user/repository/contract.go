package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"gorm.io/gorm"
)

var config = configs.GetInstance()

type UserRepository interface {
	CreateCache(ctx context.Context, key string, data interface{})
	GetCache(ctx context.Context, key string, res interface{})
	DelCache(ctx context.Context, key string) error
	Register(ctx context.Context, input models.User) (res *models.User, err error)
	Login(ctx context.Context, input payload.LoginReq) (res *models.User, er error)
	GetDetail(ctx context.Context, input payload.DetailReq) (res *models.User, err error)
	Updates(ctx context.Context, input models.User) (res *models.User, err error)
	UpdateSpecificField(ctx context.Context, input models.User) (err error)
}

type userRepo struct {
	DB    *gorm.DB
	Cache driver.RedisClient
}

func NewUserRepository(
	db *gorm.DB,
	cache driver.RedisClient) UserRepository {
	return &userRepo{
		DB:    db,
		Cache: cache,
	}
}
