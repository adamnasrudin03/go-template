package delivery

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) GetList(ctx *gin.Context) {
	var (
		opName   = "UserDelivery-GetList"
		userRole = ctx.MustGet("role").(string)
		input    dto.ListUserReq
		err      error
	)

	err = ctx.ShouldBindQuery(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	input.UserRole = userRole
	res, err := c.Service.GetList(ctx, &input)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, res)
}
