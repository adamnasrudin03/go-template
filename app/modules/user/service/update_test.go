package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
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
		input    dto.UpdateReq
		mockFunc func(input dto.UpdateReq)
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
			input: dto.UpdateReq{
				ID:        1,
				UpdatedBy: 1,
			},
			mockFunc: func(input dto.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: input.ID, Columns: "id"}).
					Return(nil, nil).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "duplicate username",
			envVars: reqEnv,
			input: dto.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input dto.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()

				srv.repo.On("CheckIsDuplicate", mock.Anything, dto.DetailReq{
					NotID:    input.ID,
					Email:    input.Email,
					Username: input.Username}).Return(errors.New("duplicate username")).Once()

			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "failed update data",
			envVars: reqEnv,
			input: dto.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input dto.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()

				srv.repo.On("CheckIsDuplicate", mock.Anything, dto.DetailReq{
					NotID:    input.ID,
					Email:    input.Email,
					Username: input.Username}).Return(nil).Once()

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
			input: dto.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input dto.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()

				srv.repo.On("CheckIsDuplicate", mock.Anything, dto.DetailReq{
					NotID:    input.ID,
					Email:    input.Email,
					Username: input.Username}).Return(nil).Once()

				user := input.ConvertToUser()
				// update
				srv.repo.On("UpdateSpecificField", mock.Anything, user).
					Return(nil).Once()
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: input.ID}).
					Return(nil, errors.New("failed get detail data after updated")).Once()
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name:    "success",
			envVars: reqEnv,
			input: dto.UpdateReq{
				ID:        1,
				Name:      "Hello world",
				Role:      models.ADMIN,
				Email:     "hello-world@email.com",
				Username:  "hello-world",
				UpdatedBy: 1,
			},
			mockFunc: func(input dto.UpdateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{
					ID: input.ID, Columns: "id"}).
					Return(&models.User{ID: 1}, nil).Once()

				srv.repo.On("CheckIsDuplicate", mock.Anything, dto.DetailReq{
					NotID:    input.ID,
					Email:    input.Email,
					Username: input.Username}).Return(nil).Once()

				user := input.ConvertToUser()
				// update
				srv.repo.On("UpdateSpecificField", mock.Anything, user).
					Return(nil).Once()
				srv.repo.On("GetDetail", mock.Anything, dto.DetailReq{ID: input.ID}).
					Return(&user, nil).Once()
				srv.repoCache.On("CreateCache", mock.Anything, mock.Anything, mock.Anything, time.Minute*5).Return(nil).Once()
				srv.repoLog.On("CreateLogActivity", mock.Anything, mock.Anything).Return(nil).Once()
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
