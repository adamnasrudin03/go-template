package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) GetDetail(ctx context.Context, input payload.DetailReq) (*models.User, error) {
	const opName = "UserService-GetDetail"
	var (
		key = fmt.Sprintf("%v-%d", opName, input.ID)
		res = new(models.User)
		err error
	)

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	srv.userRepository.GetCache(ctx, key, &res)
	if res != nil && res.ID > 0 {
		return res, nil
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

	key = fmt.Sprintf("%v-%d", opName, res.ID)
	srv.userRepository.CreateCache(ctx, key, res)

	res.Password = ""
	res.Salt = ""
	res.Role = strings.ReplaceAll(strings.ToLower(res.Role), "_", " ")
	return res, nil
}
