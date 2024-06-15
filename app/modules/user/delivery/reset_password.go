package delivery

import (
	"net/http"
	"strconv"

	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) ResetPassword(ctx *gin.Context) {
	var (
		opName = "UserDelivery-ResetPassword"
		input  dto.ResetPasswordReq
		err    error
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		c.Logger.Errorf("%v error parse param: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrInvalid("ID Pengguna", "User ID"))
		return
	}
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	input.ID = ID
	input.UpdatedBy = ID
	err = c.Service.ResetPassword(ctx, &input)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, "Reset Password Successfully")
}
