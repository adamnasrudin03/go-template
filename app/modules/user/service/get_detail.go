package service

import (
	"context"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/driver"
)

func (srv *userService) GetDetail(ctx context.Context, input payload.DetailReq) (*models.User, error) {
	const opName = "UserService-GetDetail"
	var (
		key = fmt.Sprintf("%v-%d", models.CacheUserDetail, input.ID)
		res = new(models.User)
		err error
	)

	defer func() {
		go func(dataLog payload.DetailReq) {
			now := time.Now()
			logData := models.Log{
				Name:        fmt.Sprintf("Read data user with id %d", dataLog.ID),
				Action:      models.Read,
				TableNameID: dataLog.ID,
				TableName:   "user",
				UserID:      dataLog.UserID,
				LogDateTime: now,
			}
			rabbit := driver.RabbitMQ{Body: logData.ToString(), QueueName: models.QueueInsertLog}
			rabbit.Publish()

		}(input)

	}()

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	srv.userRepository.GetCache(ctx, key, &res)
	if res != nil && res.ID > 0 {
		return res, nil
	}

	res, err = srv.getDetail(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v ", opName, err)
		return nil, err
	}

	key = fmt.Sprintf("%v-%d", models.CacheUserDetail, res.ID)
	go srv.userRepository.CreateCache(context.Background(), key, res)

	res.ConvertToResponse()
	return res, nil
}
