package delivery

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (c *msgDelivery) Consume(ctx *gin.Context) {
	if c.Cfg.App.UseRabbitMQ {
		for _, v := range models.QueueList {
			go c.consumeRabbitMQ(v)
		}
	}
	helpers.RenderJSON(ctx.Writer, http.StatusOK, "Request Success")
}
