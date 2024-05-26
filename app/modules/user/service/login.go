package service

import (
	"errors"
	"log"

	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (srv *userService) Login(ctx *gin.Context, input payload.LoginReq) (res payload.LoginRes, err error) {
	const opName = "UserService-Login"

	user, err := srv.userRepository.Login(ctx, input)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("%v error: %v \n", opName, err)
		return res, helpers.ErrDataNotFound("Pengguna", "User")
	}

	if err != nil {
		log.Printf("%v error: %v \n", opName, err)
		return res, err
	}

	res.Token, err = helpers.GenerateToken(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		log.Printf("%v failed generate token: %v \n", opName, err)
		return res, err
	}

	return res, nil
}
