package delivery

import (
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/gin-gonic/gin"
)

func (c *msgDelivery) Consume(ctx *gin.Context) {
	if c.Cfg.App.UseRabbitMQ {
		for _, v := range models.QueueList {
			go c.consumeRabbitMQ(v)
		}
	}
}
