package delivery

import (
	"net/http"
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (c *userDelivery) RegisterUser(ctx *gin.Context) {
	var (
		input payload.RegisterUserReq
	)
	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, nil))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, nil))
		return
	}

	input.Email = strings.TrimSpace(input.Email)
	req := payload.RegisterReq{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     models.USER,
	}
	res, err := c.Service.Register(ctx, req)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}
	helpers.RenderJSON(ctx.Writer, http.StatusCreated, res)
}
