package repository

import (
	"context"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *userRepo) Updates(ctx context.Context, input models.User) (res *models.User, err error) {
	const opName = "UserRepository-Updates"
	err = r.DB.WithContext(ctx).Where("id = ?", input.ID).Updates(&input).Error
	if err != nil {
		log.Printf("%v error: %v \n", opName, err)
		return nil, err
	}

	return &input, nil
}