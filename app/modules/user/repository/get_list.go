package repository

import (
	"context"
	"fmt"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"gorm.io/gorm"
)

func (r *userRepo) GetList(ctx context.Context, params dto.ListUserReq) (res []models.User, err error) {
	var (
		opName = "UserRepository-GetList"
	)

	column := `id,name,role,username,email,created_at,updated_at`
	if params.CustomColumns != `` {
		column = params.CustomColumns
	}

	db := r.DB.Select(column).Model(models.User{})
	if params.Search != `` {
		db = db.Where("name LIKE ?", "%"+params.Search+"%")
	}
	if params.Role != `` {
		db = db.Where("role = ?", params.Role)
	}
	if params.NotIncRoleRoot {
		db = db.Where("role != ?", models.ROOT)
	}

	db = r.defaultQuery(db, params.BasedFilter)
	err = db.WithContext(ctx).Find(&res).Error
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return []models.User{}, err
	}

	return res, nil
}

func (r *userRepo) defaultQuery(db *gorm.DB, params models.BasedFilter) *gorm.DB {

	if models.IsValidOrderBy[params.OrderBy] && params.SortBy != `` {
		db = db.Order(fmt.Sprintf("%s %s ", params.SortBy, params.OrderBy))
	}

	if !params.IsNotDefaultQuery {
		params = params.DefaultQuery()
	}

	if !params.IsNoLimit {
		db = db.Offset(params.Offset).Limit(params.Limit)
	}

	return db
}
