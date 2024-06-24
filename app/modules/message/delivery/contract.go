package delivery

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	logService "github.com/adamnasrudin03/go-template/app/modules/log/service"
	messageSrv "github.com/adamnasrudin03/go-template/app/modules/message/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MessageDelivery interface {
	Consume(ctx *gin.Context)
	TranslateLangID(ctx *gin.Context)
}

type msgDelivery struct {
	Service messageSrv.MessageService
	LogSrv  logService.LogService
	Cfg     *configs.Configs
	Logger  *logrus.Logger
}

func NewMessageDelivery(msgSrv messageSrv.MessageService, log logService.LogService, cfg *configs.Configs, logger *logrus.Logger) MessageDelivery {
	return &msgDelivery{
		Service: msgSrv,
		LogSrv:  log,
		Cfg:     cfg,
		Logger:  logger,
	}
}
