package service

import (
	"context"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *UserSrv) ChangePassword(ctx context.Context, input dto.ChangePasswordReq) error {
	const opName = "UserService-ChangePassword"
	var (
		key  = models.GenerateKeyCacheUserDetail(input.ID)
		user = new(models.User)
		err  error
	)
	defer helpers.PanicRecover(opName)

	err = input.Validate()
	if err != nil {
		return err
	}

	srv.RepoCache.GetCache(ctx, key, user)
	useCache := user != nil && user.ID > 0
	if !useCache {
		user, err = srv.getDetail(ctx, dto.DetailReq{ID: input.ID})
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
	user.Password = newPass

	user, err = srv.Repo.Updates(ctx, *user)
	if err != nil {
		srv.Logger.Errorf("%v error: %v ", opName, err)
		return helpers.ErrUpdatedDB()
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
