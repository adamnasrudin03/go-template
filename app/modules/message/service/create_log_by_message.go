package service

import (
	"context"
	"encoding/json"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/rabbitmq/amqp091-go"
)

func (srv *MessageSrv) createLog(ctx context.Context, d amqp091.Delivery, message string) {
	const opName = "MessageService-createLog"
	err := srv.CreateLogByMessage(ctx, message)
	if err != nil {
		srv.Logger.Warnf("%s; Consume queue %s, failed to process a message: %v", opName, d.RoutingKey, err)
		d.Nack(false, false)
	} else {
		srv.Logger.Infof("%s; Consume queue %s, successfully processed:  %s", opName, d.RoutingKey, d.Body)
		d.Ack(false)
	}
}

func (srv *MessageSrv) CreateLogByMessage(ctx context.Context, message string) (err error) {
	const opName = "MessageService-CreateLogByMessage"
	defer help.PanicRecover(opName)
	if message == "" {
		return response_mapper.ErrIsRequired("Pesan", "Message")
	}

	dto := models.Log{}
	err = json.Unmarshal([]byte(message), &dto)
	if err != nil {
		srv.Logger.Errorf("%v Unmarshal error: %v ", opName, err)
		return response_mapper.ErrUnmarshalJSON()
	}

	err = srv.RepoLog.Create(ctx, dto)
	if err != nil {
		srv.Logger.Errorf("%v error create db: %v ", opName, err)
		return response_mapper.ErrCreatedDB()
	}

	return nil
}
