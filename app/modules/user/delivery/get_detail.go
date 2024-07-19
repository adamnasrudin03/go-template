package delivery

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/gin-gonic/gin"
)

func (c *userDelivery) GetDetail(ctx *gin.Context) {
	var (
		opName   = "UserDelivery-GetDetail"
		userID   = ctx.MustGet("id").(uint64)
		userRole = ctx.MustGet("role").(string)
		idParam  = strings.TrimSpace(ctx.Param("id"))
		err      error
		ID       uint64
	)

	ID = userID
	if idParam != "" {
		tmp, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.Logger.Errorf("%v error parse param: %v ", opName, err)
			response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrInvalid("ID Pengguna", "User ID"))
			return
		}
		ID = tmp
	}

	if userRole != models.ROOT && userID != ID {
		response_mapper.RenderJSON(ctx.Writer, http.StatusForbidden, response_mapper.ErrCannotHaveAccessResources())
		return
	}

	res, err := c.Service.GetDetail(ctx, dto.DetailReq{
		ID:     ID,
		UserID: userID,
	})

	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, res)
}
