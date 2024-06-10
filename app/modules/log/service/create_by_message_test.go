package service

import (
	"testing"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/stretchr/testify/mock"
)

func (srv *LogServiceTestSuite) Test_logSrv_CreateByMessage() {
	now := time.Now()

	payload := models.Log{
		Name:        "Login user hello world(helloworld@email.com)",
		Action:      models.Read,
		TableNameID: 1,
		TableName:   "user",
		UserID:      1,
		LogDateTime: time.Date(now.Year(), now.Month(), now.Day(), 00, 00, 00, 0, time.UTC),
	}
	payloadByte, _ := helpers.SafeJsonMarshal(payload)
	inputMessage := string(payloadByte)

	type args struct {
		message string
	}
	tests := []struct {
		name     string
		mockFunc func()
		args     args
		wantErr  bool
	}{
		{
			name: "params message is empty",
			mockFunc: func() {
			},
			args: args{
				message: "",
			},
			wantErr: true,
		},
		{
			name: "err unmarshal message when params message is invalid",
			mockFunc: func() {
			},
			args: args{
				message: "invalid message",
			},
			wantErr: true,
		},
		{
			name: "err create data to db",
			mockFunc: func() {
				srv.repoLog.On("Create", mock.Anything, payload).Return(helpers.ErrCreatedDB()).Once()
			},
			args: args{
				message: inputMessage,
			},
			wantErr: true,
		},
		{
			name: "success",
			mockFunc: func() {
				srv.repoLog.On("Create", mock.Anything, payload).Return(nil).Once()
			},
			args: args{
				message: inputMessage,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		srv.T().Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			if err := srv.service.CreateByMessage(srv.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("logSrv.CreateByMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
