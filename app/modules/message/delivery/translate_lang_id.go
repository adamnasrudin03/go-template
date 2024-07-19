package delivery

import (
	"net/http"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/modules/message/dto"
	"github.com/gin-gonic/gin"
)

func (c *msgDelivery) TranslateLangID(ctx *gin.Context) {
	var (
		opName = "UserDelivery-TranslateLangID"
		input  dto.Translate
		err    error
	)

	err = ctx.ShouldBindJSON(&input)
	c.Logger.Tracef("%v bind Body json: %v ", opName, err)
	if err != nil || input.OriginalText == "" {
		err = ctx.ShouldBindQuery(&input)
		c.Logger.Tracef("%v bind Query json: %v ", opName, err)
	}

	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	input.TargetLanguage = help.LangID
	text, err := help.Translate(input.OriginalText, help.Auto, input.TargetLanguage)
	if err != nil {
		c.Logger.Errorf("%v error translate: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, response_mapper.ErrFailedTranslateText())
		return
	}

	input.TranslatedText = text
	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, input)
}
