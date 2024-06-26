package delivery

import (
	"github.com/gin-gonic/gin"
)

func (c *msgDelivery) Consume(ctx *gin.Context) {
	if c.Cfg.App.UseRabbitMQ {
		c.Service.Consume(ctx)
	}
}
