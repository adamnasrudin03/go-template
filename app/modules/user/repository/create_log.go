package repository

import (
	"context"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *userRepo) InsertLog(ctx context.Context, input models.Log) (err error) {
	const opName = "UserRepository-InsertLog"
	err = r.DB.WithContext(ctx).Create(&input).Error
	if err != nil {
		log.Printf("%v error: %v \n", opName, err)
		return err
	}

	return nil
}
