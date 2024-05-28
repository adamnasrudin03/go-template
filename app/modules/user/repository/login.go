package repository

import (
	"context"
	"errors"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"gorm.io/gorm"
)

func (r *userRepo) Login(ctx context.Context, input payload.LoginReq) (res *models.User, err error) {
	const opName = "UserRepository-Login"
	err = r.DB.Where("email = ? OR username = ?", input.Username, input.Username).WithContext(ctx).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		log.Printf("%v error get db: %v \n", opName, err)
		return nil, helpers.ErrDB()
	}

	if !helpers.PasswordValid(res.Password, input.Password) {
		log.Printf("%v invalid password \n", opName)
		return nil, helpers.ErrInvalid("Kata Sandi", "Password")
	}

	return res, nil
}
