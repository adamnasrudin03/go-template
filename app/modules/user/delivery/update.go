package delivery

import (
	"net/http"
	"strconv"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) Update(ctx *gin.Context) {
	var (
		opName   = "UserDelivery-Update"
		userID   = ctx.MustGet("id").(uint64)
		userRole = ctx.MustGet("role").(string)
		input    payload.UpdateReq
		err      error
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

	if userRole != models.ROOT && userID != ID {
		helpers.RenderJSON(ctx.Writer, http.StatusForbidden, helpers.ErrCannotHaveAccessUpdateData())
		return
	}

	input.ID = ID
	input.UpdatedBy = userID
	res, err := c.Service.Update(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}
