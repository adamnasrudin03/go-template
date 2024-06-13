package service

import (
	"context"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) Register(ctx context.Context, input dto.RegisterReq) (res *models.User, err error) {
	const opName = "UserService-Register"
	defer helpers.PanicRecover(opName)

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
		Username: input.Username,
		Role:     input.Role,
		DefaultModel: models.DefaultModel{
			CreatedBy: input.CreatedBy,
			UpdatedBy: input.CreatedBy,
		},
	}

	err = srv.userRepository.CheckIsDuplicate(ctx, dto.DetailReq{
		Email:    input.Email,
		Username: input.Username,
	})
	if err != nil {
		return nil, err
	}

	res, err = srv.userRepository.Register(ctx, user)
	if err != nil || res == nil {
		srv.Logger.Errorf("%v error create data: %v", opName, err)
		return nil, helpers.ErrCreatedDB()
	}

	res.ConvertToResponse()

	go func(dataLog models.User) {
		now := time.Now()
		logData := models.Log{
			Name:        fmt.Sprintf("Registered user %s(%s) with %s role", dataLog.Name, dataLog.Email, dataLog.Role),
			Action:      models.Created,
			TableNameID: dataLog.ID,
			TableName:   "user",
			UserID:      dataLog.CreatedBy,
			LogDateTime: now,
		}
		srv.createLogActivity(context.Background(), logData)

	}(*res)

	return res, nil
}
