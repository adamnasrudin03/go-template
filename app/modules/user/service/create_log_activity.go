package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) createLogActivity(ctx context.Context, input models.Log) (err error) {
	const opName = "UserService-createLogActivity"

	if srv.Cfg.App.UseRabbitMQ {
		srv.Logger.Info("Using insert log activity by rabbitMQ...")
		rabbit := driver.RabbitMQ{Body: input.ToString(), QueueName: models.QueueInsertLog}
		rabbit.Publish()
		return nil
	}

	err = srv.userRepository.InsertLog(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v", opName, err)
		return helpers.ErrCreatedDB()
	}

	return nil
}
