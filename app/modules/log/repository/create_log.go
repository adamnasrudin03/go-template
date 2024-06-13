package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *logRepo) CreateLog(ctx context.Context, input models.Log) (err error) {
	const opName = "LogRepository-CreateLog"
	err = r.DB.WithContext(ctx).Create(&input).Error
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return err
	}

	return nil
}
