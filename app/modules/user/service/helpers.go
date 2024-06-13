package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *UserSrv) getDetail(ctx context.Context, input dto.DetailReq) (*models.User, error) {
	const opName = "UserService-getDetail"
	var (
		res = new(models.User)
		err error
	)

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	res, err = srv.Repo.GetDetail(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v", opName, err)
		return nil, helpers.ErrDB()
	}

	isExist := res != nil && res.ID > 0
	if !isExist {
		return nil, helpers.ErrDataNotFound("Pengguna", "User")
	}

	return res, nil
}

func (srv *UserSrv) convertModelsToListResponse(data []models.User) []dto.UserRes {
	var records = []dto.UserRes{}

	for i := 0; i < len(data); i++ {
		temp := dto.UserRes{
			ID:        data[i].ID,
			Name:      data[i].Name,
			Role:      helpers.ToTitle(data[i].Role),
			Username:  data[i].Username,
			Email:     data[i].Email,
			CreatedAt: data[i].CreatedAt,
			UpdatedAt: data[i].UpdatedAt,
		}
		records = append(records, temp)
	}

	return records
}
