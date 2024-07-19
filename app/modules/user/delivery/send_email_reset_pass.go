package delivery

import (
	"net/http"
	"strconv"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) SendEmailResetPass(ctx *gin.Context) {
	var (
		opName = "UserDelivery-SendEmailResetPass"
		err    error
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		c.Logger.Errorf("%v error parse param: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrInvalid("ID Pengguna", "User ID"))
		return
	}

	resp, err := c.Service.SendEmailResetPass(ctx, ID)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, resp)
}
