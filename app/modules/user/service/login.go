package service

import (
	"context"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) Login(ctx context.Context, input payload.LoginReq) (res *payload.LoginRes, err error) {
	const opName = "UserService-Login"
	defer helpers.PanicRecover(opName)

	user, err := srv.userRepository.Login(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v", opName, err)
		return res, err
	}
	isExist := user != nil && user.ID > 0
	if !isExist {
		return nil, helpers.ErrDataNotFound("Pengguna", "User")
	}

	res = &payload.LoginRes{}
	res.Token, err = helpers.GenerateToken(helpers.JWTClaims{
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
		srv.createLogActivity(context.Background(), logData)

	}(*user)
	return res, nil
}
