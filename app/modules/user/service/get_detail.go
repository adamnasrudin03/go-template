package service

import (
	"context"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *UserSrv) GetDetail(ctx context.Context, input dto.DetailReq) (*models.User, error) {
	const opName = "UserService-GetDetail"
	var (
		key = fmt.Sprintf("%v-%d", models.CacheUserDetail, input.ID)
		res = new(models.User)
		err error
	)

	defer func() {
		helpers.PanicRecover(opName)
		go func(dataLog dto.DetailReq) {
			now := time.Now()
			logData := models.Log{
				Name:        fmt.Sprintf("Read data user with id %d", dataLog.ID),
				Action:      models.Read,
				TableNameID: dataLog.ID,
				TableName:   "user",
				UserID:      dataLog.UserID,
				LogDateTime: now,
			}
			srv.RepoLog.CreateLogActivity(context.Background(), logData)
		}(input)
	}()

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	srv.RepoCache.GetCache(ctx, key, res)
	if res != nil && res.ID > 0 {
		return res, nil
	}

	res, err = srv.getDetail(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v ", opName, err)
		return nil, err
	}

	key = fmt.Sprintf("%v-%d", models.CacheUserDetail, res.ID)
	go srv.RepoCache.CreateCache(context.Background(), key, res, time.Minute*5)

	res.ConvertToResponse()
	return res, nil
}
