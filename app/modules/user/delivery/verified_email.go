package delivery

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
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
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	input.UserID = userID
	err = c.Service.VerifiedEmail(ctx, &input)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, "Success Verified Email")
}
