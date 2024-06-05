package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/message/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *msgDelivery) TranslateLangID(ctx *gin.Context) {
	var (
		opName = "UserDelivery-TranslateLangID"
		input  payload.Translate
		err    error
	)

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}
	if input.Text == "" {
		err = ctx.ShouldBindQuery(&input)
		if err != nil {
			c.Logger.Errorf("%v error bind Query json: %v ", opName, err)
			helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
			return
		}
	}

	input.TargetLanguage = helpers.LangID
	text, err := helpers.Translate(input.Text, helpers.Auto, input.TargetLanguage)
	if err != nil {
		c.Logger.Errorf("%v error translate: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, helpers.ErrFailedTranslateText())
		return
	}

	input.Text = text
	helpers.RenderJSON(ctx.Writer, http.StatusOK, input)
}
