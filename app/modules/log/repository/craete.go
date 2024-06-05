package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *logRepo) Create(ctx context.Context, input models.Log) (err error) {
	const opName = "LogRepository-Create"
	err = r.DB.WithContext(ctx).Create(&input).Error
	if err != nil {
		r.Logger.Errorf("%v error create db: %v ", opName, err)
		return err
	}

	return nil
}
