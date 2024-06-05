package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (c *userDelivery) Register(ctx *gin.Context) {
	var (
		opName = "UserDelivery-Register"
		input  payload.RegisterReq
	)

	userID := ctx.MustGet("id").(uint64)
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	input.CreatedBy = userID
	res, err := c.Service.Register(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusCreated, res)
}
