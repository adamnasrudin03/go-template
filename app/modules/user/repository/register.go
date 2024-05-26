package repository

import (
	"log"

	"github.com/adamnasrudin03/go-template/app/models"

	"github.com/gin-gonic/gin"
)

func (r *userRepo) Register(ctx *gin.Context, input models.User) (res *models.User, err error) {
	const opName = "UserRepository-Register"
	err = r.DB.WithContext(ctx).Create(&input).Error
	if err != nil {
		log.Printf("%v error register new user: %v \n", opName, err)
		return nil, err
	}

	return &input, nil
}
