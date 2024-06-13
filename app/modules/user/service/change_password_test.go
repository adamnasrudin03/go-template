package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/stretchr/testify/mock"
)

func (srv *UserServiceTestSuite) Test_userService_ChangePassword() {

	reqEnv := map[string]string{
		"USE_RABBIT": "false",
	}
	tests := []struct {
		name     string
		envVars  map[string]string
		input    dto.ChangePasswordReq
		mockFunc func(input dto.ChangePasswordReq)
		wantErr  bool
	}{
		{
			name:    "err invalid params",
			envVars: reqEnv,
			input: dto.ChangePasswordReq{
				ID: 0,
			},
			wantErr: true,
		},
		{
			name:    "user not found",
			envVars: reqEnv,
			input: dto.ChangePasswordReq{
				ID:              1,
				Password:        "password123",
				ConfirmPassword: "password456",
				NewPassword:     "password456",
			},
			mockFunc: func(input dto.ChangePasswordReq) {
				key := fmt.Sprintf("%v-%d", models.CacheUserDetail, 1)
				srv.repo.On("GetCache", mock.Anything, key, &models.User{ID: 0}).Return(nil).Once()
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: 1}).Return(nil, nil).Once()

			},
			wantErr: true,
		},
		{
			name:    "invalid old password",
			envVars: reqEnv,
			input: dto.ChangePasswordReq{
				ID:              1,
				Password:        "invalid-old-password",
				ConfirmPassword: "password456",
				NewPassword:     "password456",
			},
			mockFunc: func(input dto.ChangePasswordReq) {
				key := fmt.Sprintf("%v-%d", models.CacheUserDetail, 1)
				srv.repo.On("GetCache", mock.Anything, key, &models.User{ID: 0}).Return(nil).Once()
				user := models.User{
					ID:   1,
					Name: "Hello World",
				}
				user.Password, _ = helpers.HashPassword("password123")
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: 1}).Return(&user, nil).Once()

			},
			wantErr: true,
		},
		{
			name:    "failed update password",
			envVars: reqEnv,
			input: dto.ChangePasswordReq{
				ID:              1,
				Password:        "password123",
				ConfirmPassword: "password456",
				NewPassword:     "password456",
				UpdatedBy:       1,
			},
			mockFunc: func(input dto.ChangePasswordReq) {
				key := fmt.Sprintf("%v-%d", models.CacheUserDetail, 1)
				srv.repo.On("GetCache", mock.Anything, key, &models.User{ID: 0}).Return(nil).Once()

				user := models.User{
					ID:       1,
					Name:     "Hello World",
					Password: input.Password,
				}
				user.Password, _ = helpers.HashPassword(user.Password)
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: input.ID}).Return(&user, nil).Once()
				srv.repo.On("Updates", mock.Anything, mock.Anything).Return(nil, errors.New("invalid update")).Once()

			},
			wantErr: true,
		},
		{
			name:    "success",
			envVars: reqEnv,
			input: dto.ChangePasswordReq{
				ID:              1,
				Password:        "password123",
				ConfirmPassword: "password456",
				NewPassword:     "password456",
				UpdatedBy:       1,
			},
			mockFunc: func(input dto.ChangePasswordReq) {
				key := fmt.Sprintf("%v-%d", models.CacheUserDetail, 1)
				srv.repo.On("GetCache", mock.Anything, key, &models.User{ID: 0}).Return(nil).Once()

				user := models.User{
					ID:       1,
					Name:     "Hello World",
					Password: input.Password,
				}
				user.Password, _ = helpers.HashPassword(user.Password)
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: 1}).Return(&user, nil).Once()
				srv.repo.On("Updates", mock.Anything, mock.Anything).Return(&user, nil).Once()
				srv.repo.On("CreateCache", mock.Anything, key, mock.Anything).Return(nil).Once()
				srv.repo.On("InsertLog", mock.Anything, mock.Anything).Return(nil).Once()

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
			if err := srv.service.ChangePassword(srv.ctx, tt.input); (err != nil) != tt.wantErr {
				t.Errorf("userService.ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
