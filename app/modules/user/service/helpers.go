package service

import (
	"context"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) getDetail(ctx context.Context, input payload.DetailReq) (*models.User, error) {
	const opName = "UserService-getDetail"
	var (
		res = new(models.User)
		err error
	)

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	res, err = srv.userRepository.GetDetail(ctx, input)
	if err != nil {
		log.Printf("%v error: %v \n", opName, err)
		return nil, helpers.ErrDB()
	}

	isExist := res != nil && res.ID > 0
	if !isExist {
		return nil, helpers.ErrDataNotFound("Pengguna", "User")
	}

	return res, nil
}
