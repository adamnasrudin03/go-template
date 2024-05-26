package repository

import (
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx *gin.Context, input models.User) (res models.User, err error)
	Login(ctx *gin.Context, input payload.LoginReq) (res models.User, er error)
	GetByEmail(ctx *gin.Context, email string) (res models.User, err error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		DB: db,
	}
}
