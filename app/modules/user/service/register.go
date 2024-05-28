package service

import (
	"context"
	"log"
	"strings"

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

	checkUser, _ := srv.userRepository.GetDetail(ctx, payload.DetailReq{Columns: "id", Email: user.Email})
	if checkUser != nil && checkUser.ID > 0 {
		log.Printf("%v Email has be registered \n", opName)
		return nil, helpers.NewError(helpers.ErrConflict, helpers.NewResponseMultiLang(
			helpers.MultiLanguages{
				ID: "Surel Sudah Terdafar",
				EN: "Email Already Registered",
			},
		))
	}

	checkUser, _ = srv.userRepository.GetDetail(ctx, payload.DetailReq{Columns: "id", Username: user.Username})
	if checkUser != nil && checkUser.ID > 0 {
		log.Printf("%v Username has be registered \n", opName)
		return nil, helpers.NewError(helpers.ErrConflict, helpers.NewResponseMultiLang(
			helpers.MultiLanguages{
				ID: "Username Sudah Terdafar",
				EN: "Username Already Registered",
			},
		))
	}

	res, err = srv.userRepository.Register(ctx, user)
	if err != nil || res == nil {
		log.Printf("%v error create data: %+v \n", opName, err)
		return nil, helpers.ErrDB()
	}

	res.Password = ""
	res.Salt = ""
	res.Role = strings.ReplaceAll(strings.ToLower(res.Role), "_", " ")

	return res, nil
}
