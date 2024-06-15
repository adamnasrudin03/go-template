package delivery

import (
	"net/http"
	"strconv"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
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
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrInvalid("ID Pengguna", "User ID"))
		return
	}

	resp, err := c.Service.SendEmailResetPass(ctx, ID)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, resp)
}
