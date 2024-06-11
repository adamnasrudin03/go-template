package service

import (
	"context"
	"errors"
	"testing"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/repository/mocks"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/stretchr/testify/mock"
)

func Test_userService_createLogActivity(t *testing.T) {

	type args struct {
		ctx   context.Context
		input models.Log
	}
	tests := []struct {
		name       string
		args       args
		beforeFunc func(params models.Log) *userService
		wantErr    bool
	}{
		{
			name: "failed insert log",
			args: args{
				ctx: context.Background(),
				input: models.Log{
					Name:        "Login user hello world(helloworld@email.com)",
					Action:      models.Read,
					TableNameID: 1,
					TableName:   "users",
					UserID:      1,
				},
			},
			beforeFunc: func(params models.Log) *userService {
				var (
					cfg      = configs.GetInstance()
					logger   = driver.Logger(cfg)
					mockUser = new(mocks.UserRepository)
				)
				srv := &userService{
					userRepository: mockUser,
					Cfg:            cfg,
					Logger:         logger,
				}
				mockUser.On("InsertLog", mock.Anything, params).Return(errors.New("failed insert log")).Once()
				return srv
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := tt.beforeFunc(tt.args.input)

			if err := srv.createLogActivity(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("userService.createLogActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
