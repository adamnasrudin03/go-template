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
		c.Logger.Infof("%v bind Body json: %v ", opName, err)
		if input.OriginalText == "" {
			err = ctx.ShouldBindQuery(&input)
			c.Logger.Infof("%v bind Query json: %v ", opName, err)
		}
	}
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	input.TargetLanguage = helpers.LangID
	text, err := helpers.Translate(input.OriginalText, helpers.Auto, input.TargetLanguage)
	if err != nil {
		c.Logger.Errorf("%v error translate: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, helpers.ErrFailedTranslateText())
		return
	}

	input.TranslatedText = text
	helpers.RenderJSON(ctx.Writer, http.StatusOK, input)
}
