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

func (srv *UserSrv) ChangePassword(ctx context.Context, input dto.ChangePasswordReq) error {
	const opName = "UserService-ChangePassword"
	var (
		key  = models.GenerateKeyCacheUserDetail(input.ID)
		user = new(models.User)
		err  error
	)
	defer help.PanicRecover(opName)

	err = input.Validate()
	if err != nil {
		return err
	}

	if ok := srv.RepoCache.GetCache(ctx, key, user); !ok {
		user, err = srv.getDetail(ctx, dto.DetailReq{ID: input.ID})
		if err != nil {
			srv.Logger.Errorf("%v error: %v ", opName, err)
			return err
		}
	}

	if !help.PasswordIsValid(user.Password, input.Password) {
		srv.Logger.Errorf("%v invalid password current ", opName)
		return response_mapper.ErrPasswordNotMatch()
	}

	newPass, err := help.HashPassword(input.NewPassword)
	if err != nil {
		srv.Logger.Errorf("%v error hash password: %v ", opName, err)
		return response_mapper.ErrHashPasswordFailed()
	}
	user.UpdatedBy = input.UpdatedBy
	user.Password = newPass

	user, err = srv.Repo.Updates(ctx, *user)
	if err != nil {
		srv.Logger.Errorf("%v error: %v ", opName, err)
		return response_mapper.ErrUpdatedDB()
	}

	go func(dataLog models.User) {
		newCtx := context.Background()
		now := time.Now()
		srv.RepoCache.CreateCache(newCtx, key, dataLog, time.Minute*5)
		logData := models.Log{
			Name:        fmt.Sprintf("Change Password User %s(%s)", dataLog.Name, dataLog.Email),
			Action:      models.Updated,
			TableNameID: dataLog.ID,
			TableName:   "user",
			UserID:      dataLog.UpdatedBy,
			LogDateTime: now,
		}
		srv.RepoLog.CreateLogActivity(newCtx, logData)
	}(*user)

	return nil
}
