package repository

import (
	"context"
	"errors"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"gorm.io/gorm"
)

func (r *userRepo) GetDetail(ctx context.Context, input dto.DetailReq) (res *models.User, err error) {
	const opName = "UserRepository-GetDetail"

	column := "*"
	if input.Columns != "" {
		column = input.Columns
	}

	db := r.whereGetDetail(r.DB.Select(column), input)

	if err = db.WithContext(ctx).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Logger.Errorf("%v error get db: %v ", opName, err)
		return nil, err
	}
	return res, nil
}

// whereGetDetail sets the where clause for the GetDetail function based on the
// input dto.
func (r *userRepo) whereGetDetail(db *gorm.DB, input dto.DetailReq) *gorm.DB {
	// Input is strongly typed
	if input.ID > 0 {
		db = db.Where("id = ?", input.ID)
	}
	if input.NotID > 0 {
		db = db.Where("id != ?", input.NotID)
	}
	if input.Email != "" {
		db = db.Where("email = ?", input.Email)
	}
	if input.Username != "" {
		db = db.Where("username = ?", input.Username)
	}
	if input.Name != "" {
		db = db.Where("name = ?", input.Name)
	}
	if input.Role != "" {
		db = db.Where("role = ?", input.Role)
	}

	return db
}
