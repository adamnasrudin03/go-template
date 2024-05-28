package delivery

import (
	"github.com/adamnasrudin03/go-template/app/modules/user/service"

	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	RegisterUser(ctx *gin.Context)
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userDelivery struct {
	Service service.UserService
}

func NewUserDelivery(srv service.UserService) UserDelivery {
	return &userDelivery{
		Service: srv,
	}
}
