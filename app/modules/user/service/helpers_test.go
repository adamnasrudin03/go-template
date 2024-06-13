package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
)

func Test_userService_getDetail(t *testing.T) {

	type args struct {
		ctx   context.Context
		input dto.DetailReq
	}
	tests := []struct {
		name       string
		args       args
		beforeFunc func() *userService
		want       *models.User
		wantErr    bool
	}{
		{
			name: "failed params",
			args: args{
				ctx: context.Background(),
			},
			beforeFunc: func() *userService {
				return &userService{}
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := tt.beforeFunc()
			got, err := srv.getDetail(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.getDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.getDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
