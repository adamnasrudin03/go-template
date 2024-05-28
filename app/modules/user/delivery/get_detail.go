package delivery

import (
	"log"
	"net/http"

	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (c *userDelivery) GetDetail(ctx *gin.Context) {
	const opName = "UserDelivery-GetDetail"
	userID := ctx.MustGet("id").(uint64)

	res, err := c.Service.GetDetail(ctx, payload.DetailReq{
		ID: userID,
	})

	if err != nil {
		log.Printf("%v error: %v \n", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}
