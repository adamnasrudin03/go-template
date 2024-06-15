package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) VerifiedEmail(ctx *gin.Context) {
	var (
		opName = "UserDelivery-VerifiedEmail"
		userID = ctx.MustGet("id").(uint64)
		input  dto.VerifyOtpReq
		err    error
	)

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	input.UserID = userID
	err = c.Service.VerifiedEmail(ctx, &input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, "Success Verified Email")
}
