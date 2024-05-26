package delivery

import (
	"net/http"
	"strings"

	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (c *userDelivery) Register(ctx *gin.Context) {
	var (
		input payload.RegisterReq
	)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
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
	input.CreatedBy = uint64(userID)
	res, err := c.Service.Register(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusCreated, res)
}
