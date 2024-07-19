package service

import (
	"context"
	"fmt"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/middlewares"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
)

func (srv *AuthSrv) Login(ctx context.Context, input dto.LoginReq) (res *dto.LoginRes, err error) {
	const opName = "AuthService-Login"
	defer help.PanicRecover(opName)

	user, err := srv.Repo.Login(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v", opName, err)
		return res, err
	}
	isExist := user != nil && user.ID > 0
	if !isExist {
		return nil, response_mapper.ErrDataNotFound("Pengguna", "User")
	}

	res = &dto.LoginRes{}
	res.Token, err = middlewares.GenerateToken(middlewares.JWTClaims{
		ID:       user.ID,
		Name:     user.Name,
		Role:     user.Role,
		Email:    user.Email,
		Username: user.Username,
	})
	if err != nil {
		srv.Logger.Errorf("%v failed generate token: %v", opName, err)
		return res, err
	}

	go func(dataLog models.User) {
		now := time.Now()
		logData := models.Log{
			Name:        fmt.Sprintf("Login user %s(%s)", dataLog.Name, dataLog.Email),
			Action:      models.Read,
			TableNameID: dataLog.ID,
			TableName:   "user",
			UserID:      dataLog.ID,
			LogDateTime: now,
		}
		srv.RepoLog.CreateLogActivity(context.Background(), logData)

	}(*user)
	return res, nil
}
