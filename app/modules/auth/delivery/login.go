package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (c *authDelivery) Login(ctx *gin.Context) {
	var (
		opName = "AuthDelivery-Login"
		input  dto.LoginReq
	)

	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	err = validate.Struct(input)
	if err != nil {
		c.Logger.Errorf("%v error validate struct: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.FormatValidationError(err))
		return
	}

	res, err := c.Service.Login(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}
