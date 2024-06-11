package repository

import (
	"context"
	"fmt"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/payload"
)

func (r *logRepo) GetList(ctx context.Context, params payload.ListLogReq) (res []models.Log, err error) {
	var (
		opName = "LogRepository-GetList"
	)

	column := `*`
	if params.CustomColumns != `` {
		column = params.CustomColumns
	}

	db := r.DB.Select(column).Model(models.Log{}).WithContext(ctx)

	if params.UserID > 0 {
		db = db.Where(`user_id = ?`, params.UserID)
	}

	if models.IsValidOrderBy[params.OrderBy] && params.SortBy != `` {
		db = db.Order(fmt.Sprintf("%s %s ", params.SortBy, params.OrderBy))
	}

	if !params.IsNotDefaultQuery {
		params.BasedFilter = params.BasedFilter.DefaultQuery()
	}

	if !params.IsNoLimit {
		db = db.Offset(params.Offset).Limit(params.Limit)
	}

	if params.UsePreload {
		db = db.Preload("User")
	}

	err = db.Find(&res).Error
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return []models.Log{}, err
	}

	return res, nil
}
