package service

import (
	"context"
	"fmt"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
)

func (srv *UserSrv) Update(ctx context.Context, input dto.UpdateReq) (res *models.User, err error) {
	const opName = "UserService-Update"
	defer help.PanicRecover(opName)

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	_, err = srv.getDetail(ctx, dto.DetailReq{ID: input.ID, Columns: "id"})
	if err != nil {
		srv.Logger.Errorf("%v error check data: %v", opName, err)
		return nil, err
	}

	err = srv.Repo.CheckIsDuplicate(ctx, dto.DetailReq{
		Email:    input.Email,
		Username: input.Username,
		NotID:    input.ID,
	})
	if err != nil {
		return nil, err
	}

	err = srv.Repo.UpdateSpecificField(ctx, input.ConvertToUser())
	if err != nil {
		srv.Logger.Errorf("%v error update data: %v", opName, err)
		return nil, response_mapper.ErrUpdatedDB()
	}

	res, err = srv.getDetail(ctx, dto.DetailReq{ID: input.ID})
	if err != nil {
		srv.Logger.Errorf("%v error get data: %v", opName, err)
		return nil, response_mapper.ErrDB()
	}

	res.ConvertToResponse()

	go func(dataLog models.User) {
		now := time.Now()
		logData := models.Log{
			Name:        fmt.Sprintf("Updated data user %s(%s)", dataLog.Name, dataLog.Email),
			Action:      models.Updated,
			TableNameID: dataLog.ID,
			TableName:   "user",
			UserID:      dataLog.UpdatedBy,
			LogDateTime: now,
		}
		srv.RepoLog.CreateLogActivity(context.Background(), logData)

	}(*res)

	return res, nil
}
