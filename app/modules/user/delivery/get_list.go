package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) GetList(ctx *gin.Context) {
	var (
		opName   = "UserDelivery-GetList"
		userRole = ctx.MustGet("role").(string)
		input    payload.ListUserReq
		err      error
	)

	err = ctx.ShouldBindQuery(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	input.UserRole = userRole
	res, err := c.Service.GetList(ctx, &input)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}
