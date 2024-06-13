package repository

import (
	"context"
	"errors"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"gorm.io/gorm"
)

func (r *AuthRepo) Login(ctx context.Context, input dto.LoginReq) (res *models.User, err error) {
	const opName = "AuthRepository-Login"
	err = r.DB.Select("id, username, email, password").
		Where("email = ? OR username = ?", input.Username, input.Username).
		WithContext(ctx).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Logger.Errorf("%v error get db: %v ", opName, err)
		return nil, helpers.ErrDB()
	}

	if !helpers.PasswordValid(res.Password, input.Password) {
		r.Logger.Errorf("%v invalid password ", opName)
		return nil, helpers.ErrInvalid("Kata Sandi", "Password")
	}

	return res, nil
}
