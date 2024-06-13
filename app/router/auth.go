package router

import (
	"github.com/adamnasrudin03/go-template/app/middlewares"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/delivery"

	"github.com/gin-gonic/gin"
)

func (r routes) authRouter(rg *gin.RouterGroup, handler delivery.AuthDelivery) {
	auth := rg.Group("/auth")
	{
		auth.POST("/sign-up", handler.RegisterUser) // Only register role user
		auth.POST("/sign-in", handler.Login)
	}

	authRoot := rg.Group("/root/auth")
	{
		authRoot.Use(middlewares.Authentication())
		authRoot.POST("/sign-up", middlewares.AuthorizationMustBe([]string{models.ROOT}), handler.Register) // register role Admin or User
	}
}
