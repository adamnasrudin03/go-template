package service

import (
	"context"
	"fmt"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) Update(ctx context.Context, input payload.UpdateReq) (res *models.User, err error) {
	const opName = "UserService-Update"

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	_, err = srv.getDetail(ctx, payload.DetailReq{ID: input.ID, Columns: "id"})
	if err != nil {
		log.Printf("%v error check data: %+v \n", opName, err)
		return nil, helpers.ErrDB()
	}

	err = srv.checkIsNotDuplicate(ctx, payload.DetailReq{
		Email:    input.Email,
		Username: input.Username,
		NotID:    input.ID,
	})
	if err != nil {
		return nil, err
	}

	err = srv.userRepository.UpdateSpecificField(ctx, input.ConvertToUser())
	if err != nil {
		log.Printf("%v error update data: %+v \n", opName, err)
		return nil, helpers.ErrUpdatedDB()
	}

	res, err = srv.getDetail(ctx, payload.DetailReq{ID: input.ID})
	if err != nil {
		log.Printf("%v error get data: %+v \n", opName, err)
		return nil, helpers.ErrDB()
	}

	srv.userRepository.CreateCache(ctx, fmt.Sprintf("%v-%d", models.CacheUserDetail, res.ID), res)
	res.ConvertToResponse()

	return res, nil
}
