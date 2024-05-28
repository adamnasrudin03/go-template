package delivery

import (
	"log"
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (c *userDelivery) Login(ctx *gin.Context) {
	var (
		opName = "UserDelivery-Login"
		input  payload.LoginReq
	)

	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		log.Printf("%v error bind json: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	err = validate.Struct(input)
	if err != nil {
		log.Printf("%v error validate struct: %v \n", opName, err)
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
