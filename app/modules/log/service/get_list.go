package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *logSrv) GetList(ctx context.Context, params *dto.ListLogReq) (*helpers.Pagination, error) {
	var (
		opName       = "LogService-GetList"
		records      = []dto.LogRes{}
		totalRecords = len(records)
	)
	defer helpers.PanicRecover(opName)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	dataDB, err := srv.Repo.GetList(ctx, *params)
	if err != nil {
		srv.Logger.Errorf("%v error get records: %v ", opName, err)
		return nil, helpers.ErrDB()
	}

	records = srv.convertModelsToListResponse(dataDB)
	totalRecords = len(records)
	resp := &helpers.Pagination{
		Meta: helpers.Meta{
			Page:         int(params.Page),
			Limit:        int(params.Limit),
			TotalRecords: totalRecords,
		},
		Data: records,
	}

	// total records in less than limit
	if totalRecords > 0 && totalRecords != params.Limit {
		return resp, nil
	}

	// get total data
	if totalRecords > 0 {
		params.CustomColumns = "id"
		params.IsNotDefaultQuery = true
		params.Offset = (params.Page - 1) * params.Limit
		params.Limit = models.DefaultLimitIsTotalDataTrue * params.Limit

		totalData, err := srv.Repo.GetList(ctx, *params)
		if err != nil {
			srv.Logger.Errorf("%v error get total records: %v ", opName, err)
			return nil, helpers.ErrDB()
		}

		totalRecords = len(totalData)
		resp.Meta.TotalRecords = totalRecords
	}

	return resp, nil
}
