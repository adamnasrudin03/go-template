package service

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	"github.com/stretchr/testify/mock"
)

func (srv *AuthServiceTestSuite) TestAuthSrv_Login() {
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
				t.Errorf("AuthService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("AuthService.Login() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func (srv *AuthServiceTestSuite) setupEnvTesting(envVars map[string]string) {
	prevEnv := os.Environ()
	for _, entry := range prevEnv {
		parts := strings.SplitN(entry, "=", 2)
		os.Unsetenv(parts[0])
	}
	for k, v := range envVars {
		os.Setenv(k, v)
	}

}

func (srv *AuthServiceTestSuite) cleanupEnvTesting(t *testing.T, envVars map[string]string) {
	prevEnv := os.Environ()
	t.Cleanup(func() {
		for k := range envVars {
			os.Unsetenv(k)
		}
		for _, entry := range prevEnv {
			parts := strings.SplitN(entry, "=", 2)
			os.Setenv(parts[0], parts[1])
		}
	})
}
