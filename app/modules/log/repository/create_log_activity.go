package repository

import (
	"context"

	help "github.com/adamnasrudin03/go-helpers"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/driver"
)

func (r *logRepo) CreateLogActivity(ctx context.Context, input models.Log) (err error) {
	const opName = "LogRepository-CreateLogActivity"
	defer help.PanicRecover(opName)

	if r.Cfg.App.UseRabbitMQ {
		r.Logger.Info("Using insert log activity by rabbitMQ...")
		rabbit := driver.RabbitMQ{Body: input.ToString(), QueueName: models.QueueInsertLog}
		rabbit.Publish()
		return nil
	}

	err = r.Create(ctx, input)
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return err
	}

	return nil
}
