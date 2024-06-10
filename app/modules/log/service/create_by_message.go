package service

import (
	"context"
	"encoding/json"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *logSrv) CreateByMessage(ctx context.Context, message string) (err error) {
	const opName = "LogService-CreateByMessage"
	defer helpers.PanicRecover(opName)
	if message == "" {
		return helpers.ErrIsRequired("Pesan", "Message")
	}

	payload := models.Log{}
	err = json.Unmarshal([]byte(message), &payload)
	if err != nil {
		srv.Logger.Errorf("%v Unmarshal error: %v ", opName, err)
		return helpers.ErrUnmarshalJSON()
	}

	err = srv.Repo.Create(ctx, payload)
	if err != nil {
		srv.Logger.Errorf("%v error create db: %v ", opName, err)
		return helpers.ErrCreatedDB()
	}

	return nil
}
