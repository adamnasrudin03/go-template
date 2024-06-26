package delivery

import (
	"github.com/adamnasrudin03/go-template/app/configs"
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
	Cfg     *configs.Configs
	Logger  *logrus.Logger
}

func NewMessageDelivery(msgSrv messageSrv.MessageService, cfg *configs.Configs, logger *logrus.Logger) MessageDelivery {
	return &msgDelivery{
		Service: msgSrv,
		Cfg:     cfg,
		Logger:  logger,
	}
}
