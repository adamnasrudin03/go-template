package delivery

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (c *authDelivery) RegisterUser(ctx *gin.Context) {
	var (
		opName = "AuthDelivery-RegisterUser"
		input  dto.RegisterUserReq
	)
	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	err = validate.Struct(input)
	if err != nil {
		c.Logger.Errorf("%v error validate struct: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.FormatValidationError(err))
		return
	}

	input.Role = models.USER
	req := input.ConvertToRegisterReq()
	res, err := c.Service.Register(ctx, req)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}
	response_mapper.RenderJSON(ctx.Writer, http.StatusCreated, res)
}
