package service

import (
	"context"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"
)

func (srv *logSrv) GetList(ctx context.Context, params *dto.ListLogReq) (*response_mapper.Pagination, error) {
	var (
		opName       = "LogService-GetList"
		records      = []dto.LogRes{}
		totalRecords = len(records)
	)
	defer help.PanicRecover(opName)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	dataDB, err := srv.Repo.GetList(ctx, *params)
	if err != nil {
		srv.Logger.Errorf("%v error get records: %v ", opName, err)
		return nil, response_mapper.ErrDB()
	}

	records = srv.convertModelsToListResponse(dataDB)
	totalRecords = len(records)
	resp := &response_mapper.Pagination{
		Meta: response_mapper.Meta{
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
			return nil, response_mapper.ErrDB()
		}

		totalRecords = len(totalData)
		resp.Meta.TotalRecords = totalRecords
	}

	return resp, nil
}
