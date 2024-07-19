package delivery

import (
	"net/http"
	"strconv"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) Update(ctx *gin.Context) {
	var (
		opName   = "UserDelivery-Update"
		userID   = ctx.MustGet("id").(uint64)
		userRole = ctx.MustGet("role").(string)
		input    dto.UpdateReq
		err      error
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		c.Logger.Errorf("%v error parse param: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrInvalid("ID Pengguna", "User ID"))
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	if userRole != models.ROOT && userID != ID {
		response_mapper.RenderJSON(ctx.Writer, http.StatusForbidden, response_mapper.ErrCannotHaveAccessUpdateData())
		return
	}

	input.ID = ID
	input.UpdatedBy = userID
	res, err := c.Service.Update(ctx, input)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, res)
}
