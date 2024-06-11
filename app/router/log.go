package router

import (
	"github.com/adamnasrudin03/go-template/app/middlewares"
	"github.com/adamnasrudin03/go-template/app/modules/log/delivery"

	"github.com/gin-gonic/gin"
)

func (r routes) LogRouter(rg *gin.RouterGroup, logDelivery delivery.LogDelivery) {
	logs := rg.Group("/logs")
	{
		logs.Use(middlewares.Authentication())
		logs.GET("", middlewares.AuthorizationMustBe([]string{}), logDelivery.GetList)
		logs.GET("/download", middlewares.AuthorizationMustBe([]string{}), logDelivery.Download)
	}

}
