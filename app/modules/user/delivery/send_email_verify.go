package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) SendEmailVerify(ctx *gin.Context) {
	var (
		opName = "UserDelivery-SendEmailVerify"
		userID = ctx.MustGet("id").(uint64)
		err    error
	)

	resp, err := c.Service.SendEmailVerify(ctx, userID)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, resp)
}
