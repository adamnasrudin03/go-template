package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/log/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (c *logDel) GetList(ctx *gin.Context) {
	var (
		opName = "UserDelivery-GetList"
		userID = ctx.MustGet("id").(uint64)
		input  payload.ListLogReq
		err    error
	)

	err = ctx.ShouldBindQuery(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	input.UserID = userID
	res, err := c.Service.GetList(ctx, &input)
	if err != nil {
		c.Logger.Errorf("%v error: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}
