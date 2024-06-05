package delivery

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	logService "github.com/adamnasrudin03/go-template/app/modules/log/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MessageDelivery interface {
	Consume(ctx *gin.Context)
	TranslateLangID(ctx *gin.Context)
}

type msgDelivery struct {
	LogSrv logService.LogService
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewMessageDelivery(log logService.LogService, cfg *configs.Configs, logger *logrus.Logger) MessageDelivery {
	return &msgDelivery{
		LogSrv: log,
		Cfg:    cfg,
		Logger: logger,
	}
}
