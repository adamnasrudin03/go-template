package service

import (
	"log"

	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (srv *userService) Login(ctx *gin.Context, input payload.LoginReq) (res *payload.LoginRes, err error) {
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
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	})
	if err != nil {
		log.Printf("%v failed generate token: %v \n", opName, err)
		return res, err
	}

	return res, nil
}
