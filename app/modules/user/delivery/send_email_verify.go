package delivery

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
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
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, resp)
}
