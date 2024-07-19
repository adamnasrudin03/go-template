package service

import (
	"errors"
	"reflect"
	"testing"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"
	"github.com/stretchr/testify/mock"
)

func (srv *LogServiceTestSuite) Test_logSrv_GetList() {
	input := &dto.ListLogReq{
		UserID:     1,
		UsePreload: true,
		BasedFilter: models.BasedFilter{
			Page:  1,
			Limit: 2,
		},
	}
	input.BasedFilter = input.BasedFilter.DefaultQuery()

	inputErr := &dto.ListLogReq{
		UserID:     1,
		UsePreload: true,
		BasedFilter: models.BasedFilter{
			Page:  1,
			Limit: 2,
		},
	}
	inputErr.BasedFilter = input.BasedFilter.DefaultQuery()

	logs := []models.Log{
		{
			ID:          1,
			Name:        "Login with user admin",
			Action:      "Read",
			TableNameID: 1,
			TableName:   "users",
			UserID:      1,
		},
		{
			ID:          2,
			Name:        "Login with user admin",
			Action:      "Read",
			TableNameID: 1,
			TableName:   "users",
			UserID:      1,
		},
	}

	records := []dto.LogRes{}
	for i := 0; i < len(logs); i++ {
		data := logs[i]
		temp := dto.LogRes{
			ID:        data.ID,
			DateTime:  data.LogDateTime,
			Name:      data.Name,
			Action:    data.Action,
			UserID:    data.UserID,
			UserName:  data.User.Name,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		records = append(records, temp)
	}

	logsLessThanLimit := logs[:1]
	recordsLessThanLimit := []dto.LogRes{}
	for i := 0; i < len(logsLessThanLimit); i++ {
		data := logsLessThanLimit[i]
		temp := dto.LogRes{
			ID:        data.ID,
			DateTime:  data.LogDateTime,
			Name:      data.Name,
			Action:    data.Action,
			UserID:    data.UserID,
			UserName:  data.User.Name,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		recordsLessThanLimit = append(recordsLessThanLimit, temp)
	}

	tests := []struct {
		name     string
		mockFunc func()
		params   *dto.ListLogReq
		want     *response_mapper.Pagination
		wantErr  bool
	}{
		{
			name: "invalid params user id",
			params: &dto.ListLogReq{
				UserID: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid order_by",
			params: &dto.ListLogReq{
				UserID: 1,
				BasedFilter: models.BasedFilter{
					OrderBy: "invalid",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed get list",
			mockFunc: func() {
				srv.repoLog.On("GetList", mock.Anything, *input).Return(nil, response_mapper.ErrDB()).Once()
			},
			params:  input,
			want:    nil,
			wantErr: true,
		},
		{
			name: "success case total records less than limit",
			mockFunc: func() {
				srv.repoLog.On("GetList", mock.Anything, *input).Return(logsLessThanLimit, nil).Once()
			},
			params: input,
			want: &response_mapper.Pagination{
				Meta: response_mapper.Meta{
					Page:         input.Page,
					Limit:        input.Limit,
					TotalRecords: len(recordsLessThanLimit),
				},
				Data: recordsLessThanLimit,
			},
			wantErr: false,
		},
		{
			name: "error case get total records",
			mockFunc: func() {
				srv.repoLog.On("GetList", mock.Anything, *inputErr).Return(logs, nil).Once()

				paramsTotal := dto.ListLogReq{
					UserID:      inputErr.UserID,
					UsePreload:  inputErr.UsePreload,
					BasedFilter: inputErr.BasedFilter,
				}
				paramsTotal.CustomColumns = "id"
				paramsTotal.IsNotDefaultQuery = true
				paramsTotal.Offset = (paramsTotal.Page - 1) * paramsTotal.Limit
				paramsTotal.Limit = models.DefaultLimitIsTotalDataTrue * paramsTotal.Limit
				srv.repoLog.On("GetList", mock.Anything, paramsTotal).Return(nil, errors.New("error get total records")).Once()
			},
			params:  inputErr,
			want:    nil,
			wantErr: true,
		},
		{
			name: "success case get total records",
			mockFunc: func() {
				srv.repoLog.On("GetList", mock.Anything, *input).Return(logs, nil).Once()

				paramsTotal := dto.ListLogReq{
					UserID:      input.UserID,
					UsePreload:  inputErr.UsePreload,
					BasedFilter: input.BasedFilter,
				}
				paramsTotal.CustomColumns = "id"
				paramsTotal.IsNotDefaultQuery = true
				paramsTotal.Offset = (paramsTotal.Page - 1) * paramsTotal.Limit
				paramsTotal.Limit = models.DefaultLimitIsTotalDataTrue * paramsTotal.Limit
				srv.repoLog.On("GetList", mock.Anything, paramsTotal).Return(logs, nil).Once()
			},
			params: input,
			want: &response_mapper.Pagination{
				Meta: response_mapper.Meta{
					Page:         input.Page,
					Limit:        input.Limit,
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
				tt.mockFunc()
			}

			got, err := srv.service.GetList(srv.ctx, tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("logSrv.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logSrv.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}
