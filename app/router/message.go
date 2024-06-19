package router

import (
	"github.com/adamnasrudin03/go-template/app/middlewares"
	"github.com/adamnasrudin03/go-template/app/modules/message/delivery"

	"github.com/gin-gonic/gin"
)

func (r routes) MessageRouter(rg *gin.RouterGroup, messageDelivery delivery.MessageDelivery) {
	logs := rg.Group("/message")
	{
		logs.GET("/translate/id", middlewares.SetAuthBasic(), messageDelivery.TranslateLangID)
		logs.GET("/consumer", middlewares.SetAuthBasic(), messageDelivery.Consume)
	}
}
