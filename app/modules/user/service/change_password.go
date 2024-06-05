package service

import (
	"context"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) ChangePassword(ctx context.Context, input payload.ChangePasswordReq) error {
	const opName = "UserService-ChangePassword"
	var (
		key  = fmt.Sprintf("%v-%d", models.CacheUserDetail, input.ID)
		user = new(models.User)
		err  error
	)

	err = input.Validate()
	if err != nil {
		return err
	}

	srv.userRepository.GetCache(ctx, key, &user)
	useCache := user != nil && user.ID > 0
	if !useCache {
		user, err = srv.getDetail(ctx, payload.DetailReq{ID: input.ID})
		if err != nil {
			srv.Logger.Errorf("%v error: %v ", opName, err)
			return err
		}
	}

	if !helpers.PasswordValid(user.Password, input.Password) {
		srv.Logger.Errorf("%v invalid password current ", opName)
		return helpers.ErrPasswordNotMatch()
	}

	newPass, err := helpers.HashPassword(input.NewPassword)
	if err != nil {
		srv.Logger.Errorf("%v error hash password: %v ", opName, err)
		return helpers.ErrHashPasswordFailed()
	}
	user.UpdatedBy = input.UpdatedBy
	user.UpdatedAt = time.Now()
	user.Password = newPass

	user, err = srv.userRepository.Updates(ctx, *user)
	if err != nil {
		srv.Logger.Errorf("%v error: %v ", opName, err)
		return helpers.ErrUpdatedDB()
	}

	go func(dataLog models.User) {
		newCtx := context.Background()
		now := time.Now()
		srv.userRepository.CreateCache(newCtx, key, dataLog)
		logData := models.Log{
			Name:        fmt.Sprintf("Change Password User %s(%s)", dataLog.Name, dataLog.Email),
			Action:      models.Updated,
			TableNameID: dataLog.ID,
			TableName:   "user",
			UserID:      dataLog.UpdatedBy,
			LogDateTime: now,
		}
		rabbit := driver.RabbitMQ{Body: logData.ToString(), QueueName: "insert_log"}
		rabbit.Publish()
	}(*user)

	return nil
}
