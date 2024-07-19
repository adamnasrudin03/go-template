package service

import (
	"context"
	"fmt"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	userDto "github.com/adamnasrudin03/go-template/app/modules/user/dto"
)

func (srv *AuthSrv) Register(ctx context.Context, input dto.RegisterReq) (res *models.User, err error) {
	const opName = "AuthService-Register"
	defer help.PanicRecover(opName)

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

	err = srv.RepoUser.CheckIsDuplicate(ctx, userDto.DetailReq{
		Email:    input.Email,
		Username: input.Username,
	})
	if err != nil {
		return nil, err
	}

	res, err = srv.Repo.Register(ctx, user)
	if err != nil || res == nil {
		srv.Logger.Errorf("%v error create data: %v", opName, err)
		return nil, response_mapper.ErrCreatedDB()
	}

	res.ConvertToResponse()

	if res.CreatedBy > 0 {
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
			srv.RepoLog.CreateLogActivity(context.Background(), logData)

		}(*res)
	}

	return res, nil
}
