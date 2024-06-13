package router

import (
	"github.com/adamnasrudin03/go-template/app/middlewares"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/delivery"

	"github.com/gin-gonic/gin"
)

func (r routes) userRouter(rg *gin.RouterGroup, userDelivery delivery.UserDelivery) {

	users := rg.Group("/users")
	{
		users.Use(middlewares.Authentication())
		users.PATCH("/:id", middlewares.AuthorizationMustBe([]string{}), userDelivery.Update)
		users.GET("/:id", middlewares.AuthorizationMustBe([]string{}), userDelivery.GetDetail)
		users.GET("", middlewares.AuthorizationMustBe([]string{models.ROOT, models.ADMIN}), userDelivery.GetList)
		users.PATCH("/change-password/:id", middlewares.AuthorizationMustBe([]string{}), userDelivery.ChangePassword)
	}
}
