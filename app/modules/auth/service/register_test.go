package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	userDto "github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/stretchr/testify/mock"
)

func (srv *AuthServiceTestSuite) TestAuthSrv_Register() {
	reqEnv := map[string]string{
		"USE_RABBIT": "false",
	}
	tests := []struct {
		name     string
		envVars  map[string]string
		input    dto.RegisterReq
		mockFunc func(params dto.RegisterReq)
		wantRes  *models.User
		wantErr  bool
	}{
		{
			name:    "err validate params",
			envVars: reqEnv,
			input: dto.RegisterReq{
				Name: "",
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "duplicated username",
			envVars: reqEnv,
			input: dto.RegisterReq{
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				Password:  "password",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			mockFunc: func(params dto.RegisterReq) {

				srv.repoUser.On("CheckIsDuplicate", mock.Anything, userDto.DetailReq{
					Email:    params.Email,
					Username: params.Username}).Return(errors.New("duplicated username")).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "failed register user",
			envVars: reqEnv,
			input: dto.RegisterReq{
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				Password:  "password",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			mockFunc: func(params dto.RegisterReq) {
				srv.repoUser.On("CheckIsDuplicate", mock.Anything, userDto.DetailReq{
					Email:    params.Email,
					Username: params.Username}).Return(nil).Once()

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
			input: dto.RegisterReq{
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				Password:  "password",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			mockFunc: func(params dto.RegisterReq) {
				srv.repoUser.On("CheckIsDuplicate", mock.Anything, userDto.DetailReq{
					Email:    params.Email,
					Username: params.Username}).Return(nil).Once()

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
				srv.repoLog.On("CreateLogActivity", mock.Anything, mock.Anything).Return(nil).Once()
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
				t.Errorf("AuthService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("AuthService.Register() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
