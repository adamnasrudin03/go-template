package service

import (
	"context"
	"fmt"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
)

func (srv *userService) GetDetail(ctx context.Context, input payload.DetailReq) (*models.User, error) {
	const opName = "UserService-GetDetail"
	var (
		key = fmt.Sprintf("%v-%d", models.CacheUserDetail, input.ID)
		res = new(models.User)
		err error
	)

	defer func() {
		res.ConvertToResponse()
		go func(dataLog payload.DetailReq) {
			newCtx := context.Background()
			srv.userRepository.InsertLog(newCtx, models.Log{
				Name:        fmt.Sprintf("Read data user with id %d", dataLog.ID),
				Action:      models.Read,
				TableNameID: dataLog.ID,
				TableName:   "user",
				UserID:      dataLog.UserID,
			})
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
		log.Printf("%v error: %v \n", opName, err)
		return nil, err
	}

	key = fmt.Sprintf("%v-%d", models.CacheUserDetail, res.ID)
	go srv.userRepository.CreateCache(context.Background(), key, res)

	return res, nil
}
