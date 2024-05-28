package service

import (
	"context"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) Register(ctx context.Context, input payload.RegisterReq) (res *models.User, err error) {
	const opName = "UserService-Register"

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

	err = srv.checkIsNotDuplicate(ctx, payload.DetailReq{
		Email:    input.Email,
		Username: input.Username,
	})
	if err != nil {
		return nil, err
	}

	res, err = srv.userRepository.Register(ctx, user)
	if err != nil || res == nil {
		log.Printf("%v error create data: %+v \n", opName, err)
		return nil, helpers.ErrCreatedDB()
	}

	res.ConvertToResponse()

	return res, nil
}
