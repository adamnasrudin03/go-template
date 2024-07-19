package service

import (
	"context"
	"encoding/json"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (srv *logSrv) CreateByMessage(ctx context.Context, message string) (err error) {
	const opName = "LogService-CreateByMessage"
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

	err = srv.Repo.Create(ctx, dto)
	if err != nil {
		srv.Logger.Errorf("%v error create db: %v ", opName, err)
		return response_mapper.ErrCreatedDB()
	}

	return nil
}
