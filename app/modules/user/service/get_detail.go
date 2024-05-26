package service

import (
	"log"
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (srv *userService) GetDetail(ctx *gin.Context, input models.User) (res *models.User, err error) {
	const opName = "UserService-Register"

	res, err = srv.userRepository.GetDetail(ctx, input)
	if err != nil {
		log.Printf("%v error: %v \n", opName, err)
		return nil, helpers.ErrDB()
	}

	isExist := res != nil && res.ID > 0
	if !isExist {
		return nil, helpers.ErrDataNotFound("Pengguna", "User")
	}

	res.Password = ""
	res.Role = strings.ReplaceAll(strings.ToLower(res.Role), "_", " ")

	return res, nil
}
