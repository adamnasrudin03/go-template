package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) GetList(ctx context.Context, params *payload.ListUserReq) (*helpers.Pagination, error) {
	var (
		opName       = "UserService-GetList"
		records      = []payload.UserRes{}
		totalRecords = len(records)
	)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	dataDB, err := srv.userRepository.GetList(ctx, *params)
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

		totalData, err := srv.userRepository.GetList(ctx, *params)
		if err != nil {
			srv.Logger.Errorf("%v error get total records: %v ", opName, err)
			return nil, helpers.ErrDB()
		}

		totalRecords = len(totalData)
		resp.Meta.TotalRecords = totalRecords
	}

	return resp, nil
}
