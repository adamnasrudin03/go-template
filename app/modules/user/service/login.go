package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) Login(ctx context.Context, input payload.LoginReq) (res *payload.LoginRes, err error) {
	const opName = "UserService-Login"

	user, err := srv.userRepository.Login(ctx, input)
	if err != nil {
		log.Printf("%v error: %v \n", opName, err)
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
		log.Printf("%v failed generate token: %v \n", opName, err)
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
		rabbit := driver.RabbitMQ{Body: logData.ToString(), QueueName: "insert_log"}
		rabbit.Publish()

	}(*user)
	return res, nil
}
