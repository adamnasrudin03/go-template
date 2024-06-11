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

func (srv *UserServiceTestSuite) Test_userService_Register() {

	reqEnv := map[string]string{
		"USE_RABBIT": "false",
	}
	tests := []struct {
		name     string
		envVars  map[string]string
		input    payload.RegisterReq
		mockFunc func(params payload.RegisterReq)
		wantRes  *models.User
		wantErr  bool
	}{
		{
			name:    "err validate params",
			envVars: reqEnv,
			input: payload.RegisterReq{
				Name: "",
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "duplicated email",
			envVars: reqEnv,
			input: payload.RegisterReq{
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				Password:  "password",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			mockFunc: func(params payload.RegisterReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", Email: params.Email}).
					Return(&models.User{ID: 1}, nil).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "duplicated username",
			envVars: reqEnv,
			input: payload.RegisterReq{
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				Password:  "password",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			mockFunc: func(params payload.RegisterReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", Email: params.Email}).
					Return(nil, nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", Username: params.Username}).
					Return(&models.User{ID: 1}, nil).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "failed register user",
			envVars: reqEnv,
			input: payload.RegisterReq{
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				Password:  "password",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			mockFunc: func(params payload.RegisterReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", Email: params.Email}).
					Return(nil, nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", Username: params.Username}).
					Return(nil, nil).Once()

				user := models.User{
					Name:     params.Name,
					Password: params.Password,
					Email:    params.Email,
					Username: params.Username,
					Role:     params.Role,
					DefaultModel: models.DefaultModel{
						CreatedBy: params.CreatedBy,
						UpdatedBy: params.CreatedBy,
					},
				}
				srv.repo.On("Register", mock.Anything, user).Return(nil, errors.New("failed register user")).Once()
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "success",
			envVars: reqEnv,
			input: payload.RegisterReq{
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				Password:  "password",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			mockFunc: func(params payload.RegisterReq) {
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", Email: params.Email}).
					Return(nil, nil).Once()
				srv.repo.On("GetDetail", mock.Anything, payload.DetailReq{
					Columns: "id", Username: params.Username}).
					Return(nil, nil).Once()

				user := models.User{
					Name:     params.Name,
					Password: params.Password,
					Email:    params.Email,
					Username: params.Username,
					Role:     params.Role,
					DefaultModel: models.DefaultModel{
						CreatedBy: params.CreatedBy,
						UpdatedBy: params.CreatedBy,
					},
				}
				srv.repo.On("Register", mock.Anything, user).Return(&user, nil).Once()
				srv.repo.On("InsertLog", mock.Anything, mock.Anything).Return(nil).Once()
			},
			wantRes: &models.User{
				Name:     "Hello world",
				Role:     helpers.ToLower(models.ADMIN),
				Email:    "hello-world@email.com",
				Username: "hello-world",
				DefaultModel: models.DefaultModel{
					CreatedBy: 1,
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

			gotRes, err := srv.service.Register(srv.ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("userService.Register() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
