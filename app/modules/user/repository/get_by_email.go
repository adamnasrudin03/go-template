package repository

import (
	"log"

	"github.com/adamnasrudin03/go-template/app/models"

	"github.com/gin-gonic/gin"
)

func (r *userRepo) GetByEmail(ctx *gin.Context, email string) (res models.User, err error) {
	const opName = "UserRepository-GetByEmail"
	if err = r.DB.Where("email = ?", email).Take(&res).Error; err != nil {
		log.Printf("%v error get db: %v \n", opName, err)
		return
	}
	return res, nil
}
