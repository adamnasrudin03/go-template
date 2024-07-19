package delivery

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"

	"github.com/gin-gonic/gin"
)

func (c *logDel) Download(ctx *gin.Context) {
	var (
		opName = "UserDelivery-Download"
		userID = ctx.MustGet("id").(uint64)
		input  dto.ListLogReq
		err    error
	)

	err = ctx.ShouldBindQuery(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	input.UserID = userID
	err = c.Service.Download(ctx, &input)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

}
