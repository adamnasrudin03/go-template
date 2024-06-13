package delivery

import (
	"github.com/adamnasrudin03/go-template/app/modules/auth/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthDelivery interface {
	RegisterUser(ctx *gin.Context)
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authDelivery struct {
	Service service.AuthService
	Logger  *logrus.Logger
}

func NewAuthDelivery(srv service.AuthService, logger *logrus.Logger) AuthDelivery {
	return &authDelivery{
		Service: srv,
		Logger:  logger,
	}
}
