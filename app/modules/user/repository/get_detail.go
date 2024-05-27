package repository

import (
	"context"
	"errors"
	"log"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"gorm.io/gorm"
)

func (r *userRepo) GetDetail(ctx context.Context, input payload.DetailReq) (res *models.User, err error) {
	const opName = "UserRepository-GetDetail"

	db := r.DB
	if input.ID > 0 {
		db = db.Where("id = ?", input.ID)
	}
	if input.Email != "" {
		db = db.Where("email = ?", input.Email)
	}
	if input.Name != "" {
		db = db.Where("name = ?", input.Name)
	}
	if input.Role != "" {
		db = db.Where("role = ?", input.Role)
	}

	if err = db.WithContext(ctx).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		log.Printf("%v error get db: %v \n", opName, err)
		return
	}
	return res, nil
}
