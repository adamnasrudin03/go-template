package delivery

import (
	"github.com/adamnasrudin03/go-template/app/modules/log/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogDelivery interface {
	GetList(ctx *gin.Context)
}

type logDel struct {
	Service service.LogService
	Logger  *logrus.Logger
}

func NewLogDelivery(srv service.LogService, logger *logrus.Logger) LogDelivery {
	return &logDel{
		Service: srv,
		Logger:  logger,
	}
}
