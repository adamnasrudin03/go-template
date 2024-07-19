package service

import (
	"errors"
	"reflect"
	"testing"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/stretchr/testify/mock"
)

func (srv *UserServiceTestSuite) Test_userService_GetList() {
	users := []models.User{
		{
			ID:   1,
			Name: "user 1",
			Role: models.USER,
		},
		{
			ID:   2,
			Name: "user 2",
			Role: models.USER,
		},
	}
	var records = []dto.UserRes{}

	for i := 0; i < len(users); i++ {
		v := users[i]
		temp := dto.UserRes{
			ID:        v.ID,
			Name:      v.Name,
			Role:      help.ToTitle(v.Role),
			Username:  v.Username,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		records = append(records, temp)
	}

	tests := []struct {
		name     string
		params   *dto.ListUserReq
		mockFunc func(params *dto.ListUserReq)
		want     *response_mapper.Pagination
		wantErr  bool
	}{
		{
			name: "invalid params",
			params: &dto.ListUserReq{
				BasedFilter: models.BasedFilter{
					OrderBy: "invalid",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed get records",
			params: &dto.ListUserReq{
				UserRole:       models.ADMIN,
				NotIncRoleRoot: true,
				BasedFilter: models.BasedFilter{
					Page:  1,
					Limit: 10,
				},
			},
			mockFunc: func(params *dto.ListUserReq) {
				params.BasedFilter = params.BasedFilter.DefaultQuery()

				srv.repo.On("GetList", mock.Anything, *params).Return(nil, errors.New("failed get record")).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed get total records",
			params: &dto.ListUserReq{
				UserRole:       models.ADMIN,
				NotIncRoleRoot: true,
				BasedFilter: models.BasedFilter{
					Page:  1,
					Limit: 2,
				},
			},
			mockFunc: func(params *dto.ListUserReq) {
				params.BasedFilter = params.BasedFilter.DefaultQuery()

				srv.repo.On("GetList", mock.Anything, *params).Return(users, nil).Once()

				input := *params
				input.CustomColumns = "id"
				input.IsNotDefaultQuery = true
				input.Offset = (input.Page - 1) * input.Limit
				input.Limit = models.DefaultLimitIsTotalDataTrue * input.Limit
				srv.repo.On("GetList", mock.Anything, input).Return(nil, errors.New("failed get total record")).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success not get total records",
			params: &dto.ListUserReq{
				UserRole:       models.ADMIN,
				NotIncRoleRoot: true,
				BasedFilter: models.BasedFilter{
					Page:  1,
					Limit: 10,
				},
			},
			mockFunc: func(params *dto.ListUserReq) {
				params.BasedFilter = params.BasedFilter.DefaultQuery()

				srv.repo.On("GetList", mock.Anything, *params).Return(users, nil).Once()
			},
			want: &response_mapper.Pagination{
				Meta: response_mapper.Meta{
					Page:         1,
					Limit:        10,
					TotalRecords: len(records),
				},
				Data: records,
			},
			wantErr: false,
		},
		{
			name: "success with get total records",
			params: &dto.ListUserReq{
				UserRole:       models.ADMIN,
				NotIncRoleRoot: true,
				BasedFilter: models.BasedFilter{
					Page:  1,
					Limit: 2,
				},
			},
			mockFunc: func(params *dto.ListUserReq) {
				params.BasedFilter = params.BasedFilter.DefaultQuery()

				srv.repo.On("GetList", mock.Anything, *params).Return(users, nil).Once()

				input := *params
				input.CustomColumns = "id"
				input.IsNotDefaultQuery = true
				input.Offset = (input.Page - 1) * input.Limit
				input.Limit = models.DefaultLimitIsTotalDataTrue * input.Limit
				srv.repo.On("GetList", mock.Anything, input).Return([]models.User{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
				}, nil).Once()
			},
			want: &response_mapper.Pagination{
				Meta: response_mapper.Meta{
					Page:         1,
					Limit:        2,
					TotalRecords: len(records),
				},
				Data: records,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		srv.T().Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(tt.params)
			}

			got, err := srv.service.GetList(srv.ctx, tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}
