package repository

import (
	"errors"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (r *userRepo) Login(ctx *gin.Context, input payload.LoginReq) (res models.User, err error) {
	const opName = "UserRepository-Login"
	if err = r.DB.Where("email = ?", input.Email).Take(&res).Error; err != nil {
		log.Printf("%v error get db: %v \n", opName, err)
		return
	}

	if !helpers.PasswordValid(res.Password, input.Password) {
		err = errors.New("invalid password")
		log.Printf("%v error cek pass: %v \n", opName, err)
		return
	}

	return res, nil
}
