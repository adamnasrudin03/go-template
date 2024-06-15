package delivery

import (
	"github.com/adamnasrudin03/go-template/app/modules/user/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserDelivery interface {
	GetDetail(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetList(ctx *gin.Context)
	SendEmailVerify(ctx *gin.Context)
	VerifiedEmail(ctx *gin.Context)
	SendEmailResetPass(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

type userDelivery struct {
	Service service.UserService
	Logger  *logrus.Logger
}

func NewUserDelivery(srv service.UserService, logger *logrus.Logger) UserDelivery {
	return &userDelivery{
		Service: srv,
		Logger:  logger,
	}
}
