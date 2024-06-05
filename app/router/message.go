package router

import (
	"github.com/adamnasrudin03/go-template/app/modules/message/delivery"

	"github.com/gin-gonic/gin"
)

func (r routes) MessageRouter(rg *gin.RouterGroup, messageDelivery delivery.MessageDelivery) {
	logs := rg.Group("/message")
	{
		logs.GET("/translate/id", messageDelivery.TranslateLangID)
		logs.GET("/consumer", messageDelivery.Consume)
	}
}
