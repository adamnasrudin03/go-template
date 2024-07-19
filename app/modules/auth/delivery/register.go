package delivery

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"

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
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	input.CreatedBy = AuthID
	res, err := c.Service.Register(ctx, input)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusCreated, res)
}
