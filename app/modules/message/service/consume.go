package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (srv *MessageSrv) Consume(ctx context.Context) {
	if srv.Cfg.App.UseRabbitMQ {
		for _, v := range models.QueueList {
			go srv.consumeRabbitMQ(v)
		}
	}
}
