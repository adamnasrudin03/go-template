package service

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/stretchr/testify/mock"
)

func (srv *UserServiceTestSuite) Test_userService_GetDetail() {
	user := models.User{
		ID:   1,
		Name: "Hello World",
	}
	reqEnv := map[string]string{
		"USE_RABBIT": "false",
	}

	tests := []struct {
		name     string
		envVars  map[string]string
		input    dto.DetailReq
		mockFunc func()
		want     *models.User
		wantErr  bool
	}{
		{
			name:    "err validate params",
			envVars: reqEnv,
			input: dto.DetailReq{
				ID: 0,
			},
			mockFunc: func() {
				defer srv.repoLog.On("CreateLogActivity", mock.Anything, mock.Anything).Return(nil).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "success with cache",
			envVars: reqEnv,
			input: dto.DetailReq{
				ID: 1,
			},
			mockFunc: func() {
				defer srv.repoLog.On("CreateLogActivity", mock.Anything, mock.Anything).Return(nil).Once()
				key := models.GenerateKeyCacheUserDetail(1)
				res := models.User{
					ID: 1,
				}
				srv.repoCache.On("GetCache", mock.Anything, key, &models.User{
					ID: 0,
				}).Return(true).Run(func(args mock.Arguments) {
					target := args.Get(2).(*models.User)
					*target = res
				}).Once()

			},
			want: &models.User{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name:    "failed get db",
			envVars: reqEnv,
			input: dto.DetailReq{
				ID: 101,
			},
			mockFunc: func() {
				defer srv.repoLog.On("CreateLogActivity", mock.Anything, mock.Anything).Return(nil).Once()
				key := models.GenerateKeyCacheUserDetail(101)

				srv.repoCache.On("GetCache", mock.Anything, key, &models.User{ID: 0}).Return(false).Once()
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: 101}).Return(nil, errors.New("failed")).Once()

			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not found",
			envVars: reqEnv,
			input: dto.DetailReq{
				ID: 101,
			},
			mockFunc: func() {
				defer srv.repoLog.On("CreateLogActivity", mock.Anything, mock.Anything).Return(nil).Once()
				key := models.GenerateKeyCacheUserDetail(101)

				srv.repoCache.On("GetCache", mock.Anything, key, &models.User{ID: 0}).Return(false).Once()
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: 101}).Return(nil, nil).Once()

			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "success",
			envVars: reqEnv,
			input: dto.DetailReq{
				ID: 1,
			},
			mockFunc: func() {
				defer srv.repoLog.On("CreateLogActivity", mock.Anything, mock.Anything).Return(nil).Once()
				key := models.GenerateKeyCacheUserDetail(1)

				srv.repoCache.On("GetCache", mock.Anything, key, &models.User{ID: 0}).Return(false).Once()
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: 1}).Return(&user, nil).Once()
				srv.repoCache.On("CreateCache", mock.Anything, key, &user, time.Minute*5).Return(nil).Once()
			},
			want:    &user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		srv.T().Run(tt.name, func(t *testing.T) {
			srv.setupEnvTesting(tt.envVars)
			defer srv.cleanupEnvTesting(t, tt.envVars)

			if tt.mockFunc != nil {
				tt.mockFunc()
			}

			got, err := srv.service.GetDetail(srv.ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (srv *UserServiceTestSuite) setupEnvTesting(envVars map[string]string) {
	prevEnv := os.Environ()
	for _, entry := range prevEnv {
		parts := strings.SplitN(entry, "=", 2)
		os.Unsetenv(parts[0])
	}
	for k, v := range envVars {
		os.Setenv(k, v)
	}

}

func (srv *UserServiceTestSuite) cleanupEnvTesting(t *testing.T, envVars map[string]string) {
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
