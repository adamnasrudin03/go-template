package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/stretchr/testify/mock"
)

func (srv *UserServiceTestSuite) Test_userService_Update() {

	reqEnv := map[string]string{
		"USE_RABBIT": "false",
	}
	tests := []struct {
		name     string
		envVars  map[string]string
		input    payload.UpdateReq
		mockFunc func(input payload.UpdateReq)
		wantRes  *models.User
		wantErr  bool
	}{
		{
			name:    "invalid params",
			envVars: reqEnv,
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "user not found",
			envVars: reqEnv,
			input: payload.UpdateReq{
				ID:        1,
				UpdatedBy: 1,
			},
			mockFunc: func(input payload.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{ID: input.ID, Columns: "id"}).
					Return(nil, nil).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "duplicate username",
			envVars: reqEnv,
			input: payload.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input payload.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()
				// check duplicate
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Email: input.Email}).
					Return(nil, nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Username: input.Username}).
					Return(&models.User{ID: 1}, nil).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "failed update data",
			envVars: reqEnv,
			input: payload.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input payload.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()
				// check duplicate
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Email: input.Email}).
					Return(nil, nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Username: input.Username}).
					Return(nil, nil).Once()

				// update
				srv.repo.On("UpdateSpecificField", mock.Anything, input.ConvertToUser()).
					Return(errors.New("failed update data")).Once()
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "get detail data after updated",
			envVars: reqEnv,
			input: payload.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input payload.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()
				// check duplicate
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Email: input.Email}).
					Return(nil, nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Username: input.Username}).
					Return(nil, nil).Once()

				user := input.ConvertToUser()
				// update
				srv.repo.On("UpdateSpecificField", mock.Anything, user).
					Return(nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{ID: input.ID}).
					Return(nil, errors.New("failed get detail data after updated")).Once()
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "success",
			envVars: reqEnv,
			input: payload.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input payload.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()
				// check duplicate
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Email: input.Email}).
					Return(nil, nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", NotID: input.ID, Username: input.Username}).
					Return(nil, nil).Once()

				user := input.ConvertToUser()
				// update
				srv.repo.On("UpdateSpecificField", mock.Anything, user).
					Return(nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{ID: input.ID}).
					Return(&user, nil).Once()
				srv.repo.On("CreateCache", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
				srv.repo.On("InsertLog", mock.Anything, mock.Anything).Return(nil).Once()
			},
			wantRes: &models.User{
				ID:       1,
				Name:     "Hello world",
				Role:     helpers.ToLower(models.ADMIN),
				Email:    "hello-world@email.com",
				Username: "hello-world",
				DefaultModel: models.DefaultModel{
					UpdatedBy: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		srv.T().Run(tt.name, func(t *testing.T) {
			srv.setupEnvTesting(tt.envVars)
			defer srv.cleanupEnvTesting(t, tt.envVars)

			if tt.mockFunc != nil {
				tt.mockFunc(tt.input)
			}
			gotRes, err := srv.service.Update(srv.ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("userService.Update() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
