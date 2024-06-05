package delivery

import (
	"net/http"
	"strconv"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (c *userDelivery) GetDetail(ctx *gin.Context) {
	var (
		opName   = "UserDelivery-GetDetail"
		userID   = ctx.MustGet("id").(uint64)
		userRole = ctx.MustGet("role").(string)
		err      error
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		c.Logger.Errorf("%v error parse param: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrInvalid("ID Pengguna", "User ID"))
		return
	}

	if userRole != models.ROOT && userID != ID {
		helpers.RenderJSON(ctx.Writer, http.StatusForbidden, helpers.ErrCannotHaveAccessResources())
		return
	}

	res, err := c.Service.GetDetail(ctx, payload.DetailReq{
		ID:     ID,
		UserID: userID,
	})

	if err != nil {
		c.Logger.Errorf("%v error: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}
