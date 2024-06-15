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
		users.GET("/send-email-verify", middlewares.AuthorizationMustBe([]string{}), userDelivery.SendEmailVerify)
		users.POST("/verified-email", middlewares.AuthorizationMustBe([]string{}), userDelivery.VerifiedEmail)
	}

	usersNoAuth := rg.Group("/users")
	{
		usersNoAuth.GET("/request-reset-password/:id", userDelivery.SendEmailResetPass)
		usersNoAuth.PATCH("/validate-reset-password/:id", userDelivery.ResetPassword)
	}
}
