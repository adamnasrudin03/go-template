package delivery

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) ChangePassword(ctx *gin.Context) {
	var (
		opName   = "UserDelivery-ChangePassword"
		userID   = ctx.MustGet("id").(uint64)
		userRole = ctx.MustGet("role").(string)
		input    payload.ChangePasswordReq
		err      error
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Printf("%v error parse param: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrInvalid("ID Pengguna", "User ID"))
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		log.Printf("%v error bind json: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	if userRole != models.ROOT && userID != ID {
		helpers.RenderJSON(ctx.Writer, http.StatusForbidden, helpers.ErrCannotHaveAccessUpdateData())
		return
	}

	input.ID = ID
	input.UpdatedBy = userID
	err = c.Service.ChangePassword(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, "Password has been changed")
}
