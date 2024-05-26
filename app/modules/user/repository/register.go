package repository

import (
	"log"

	"github.com/adamnasrudin03/go-template/app/models"

	"github.com/gin-gonic/gin"
)

func (r *userRepo) Register(ctx *gin.Context, input models.User) (res models.User, err error) {
	const opName = "UserRepository-Register"
	if err := r.DB.Create(&input).Error; err != nil {
		log.Printf("%v error register new user: %v \n", opName, err)
		return input, err
	}

	return input, nil
}
