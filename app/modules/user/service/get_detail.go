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
		key = models.GenerateKeyCacheUserDetail(input.ID)
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

	if ok := srv.RepoCache.GetCache(ctx, key, res); !ok {
		res, err = srv.getDetail(ctx, input)
		if err != nil {
			srv.Logger.Errorf("%v error: %v ", opName, err)
			return nil, err
		}
	}

	res.ConvertToResponse()
	return res, nil
}
