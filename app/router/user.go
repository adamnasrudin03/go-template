package router

import (
	"github.com/adamnasrudin03/go-template/app/middlewares"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/delivery"

	"github.com/gin-gonic/gin"
)

func (r routes) userRouter(rg *gin.RouterGroup, userDelivery delivery.UserDelivery) {
	auth := rg.Group("/auth")
	{
		auth.POST("/sign-up", userDelivery.RegisterUser) // Only register role user
		auth.POST("/sign-in", userDelivery.Login)
	}

	users := rg.Group("/users")
	{
		users.Use(middlewares.Authentication())
		users.PATCH("/:id", middlewares.AuthorizationMustBe([]string{}), userDelivery.Update)
		users.GET("/detail", middlewares.AuthorizationMustBe([]string{}), userDelivery.GetDetail)
		users.PATCH("/change-password/:id", middlewares.AuthorizationMustBe([]string{}), userDelivery.ChangePassword)
	}
}

func (r routes) userRootRouter(rg *gin.RouterGroup, userDelivery delivery.UserDelivery) {
	auth := rg.Group("/root/auth")
	{
		auth.Use(middlewares.Authentication())
		auth.POST("/sign-up", middlewares.AuthorizationMustBe([]string{models.ROOT}), userDelivery.Register) // register role Admin or User
	}
}
