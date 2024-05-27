package service

import (
	"context"
	"errors"
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
		Role:     input.Role,
		DefaultModel: models.DefaultModel{
			CreatedBy: input.CreatedBy,
			UpdatedBy: input.CreatedBy,
		},
	}
	checkUser, _ := srv.userRepository.GetDetail(ctx, payload.DetailReq{Email: user.Email})
	if checkUser != nil && checkUser.Email != "" {
		err = errors.New("email user has be registered")
		log.Printf("%v error check email: %v \n", opName, err)
		return nil, helpers.NewError(helpers.ErrConflict, helpers.NewResponseMultiLang(
			helpers.MultiLanguages{
				ID: "Alamat Surel Sudah Terdafar",
				EN: "Email Already Registered",
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
