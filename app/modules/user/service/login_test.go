package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/stretchr/testify/mock"
)

func (srv *UserServiceTestSuite) Test_userService_Login() {

	reqEnv := map[string]string{
		"USE_RABBIT": "false",
	}
	tests := []struct {
		name     string
		envVars  map[string]string
		input    dto.LoginReq
		mockFunc func(params dto.LoginReq)
		wantRes  *dto.LoginRes
		wantErr  bool
	}{
		{
			name:    "failed get user",
			envVars: reqEnv,
			input: dto.LoginReq{
				Username: "hello-world",
				Password: "password123",
			},
			mockFunc: func(params dto.LoginReq) {
				srv.repo.On("Login", mock.Anything, params).Return(nil, errors.New("failed get user")).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "user not found",
			envVars: reqEnv,
			input: dto.LoginReq{
				Username: "hello-world",
				Password: "password123",
			},
			mockFunc: func(params dto.LoginReq) {
				srv.repo.On("Login", mock.Anything, params).Return(nil, nil).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		srv.T().Run(tt.name, func(t *testing.T) {
			srv.setupEnvTesting(tt.envVars)
			defer srv.cleanupEnvTesting(t, tt.envVars)

			if tt.mockFunc != nil {
				tt.mockFunc(tt.input)
			}
			gotRes, err := srv.service.Login(srv.ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("userService.Login() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
