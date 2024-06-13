package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (c *authDelivery) Register(ctx *gin.Context) {
	var (
		opName = "AuthDelivery-Register"
		input  dto.RegisterReq
	)

	AuthID := ctx.MustGet("id").(uint64)
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	input.CreatedBy = AuthID
	res, err := c.Service.Register(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusCreated, res)
}
