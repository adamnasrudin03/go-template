package repository

import (
	"context"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *userRepo) Updates(ctx context.Context, input models.User) (res *models.User, err error) {
	const opName = "UserRepository-Updates"
	input.UpdatedAt = time.Now()
	err = r.DB.WithContext(ctx).Where("id = ?", input.ID).Updates(&input).Error
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return nil, err
	}

	return &input, nil
}
