package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *userRepo) Register(ctx context.Context, input models.User) (res *models.User, err error) {
	const opName = "UserRepository-Register"
	err = r.DB.WithContext(ctx).Create(&input).Error
	if err != nil {
		r.Logger.Errorf("%v error register new user: %v ", opName, err)
		return nil, err
	}

	return &input, nil
}
