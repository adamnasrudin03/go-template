package router

import (
	"github.com/adamnasrudin03/go-template/app/middlewares"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/delivery"

	"github.com/gin-gonic/gin"
)

func (r routes) userRouter(rg *gin.RouterGroup, userDelivery delivery.UserDelivery) {
	users := rg.Group("/auth")
	{
		users.POST("/sign-up", userDelivery.RegisterUser)
		users.POST("/sign-in", userDelivery.Login)
	}
}

func (r routes) userRouterAuth(rg *gin.RouterGroup, userDelivery delivery.UserDelivery) {
	users := rg.Group("/auth-admin")
	{
		users.Use(middlewares.Authentication())
		users.GET("/detail", middlewares.AuthorizationMustBe([]string{models.ALL}), userDelivery.GetDetail)
		users.POST("/sign-up", middlewares.AuthorizationMustBe([]string{models.ROOT}), userDelivery.Register)
	}
}
