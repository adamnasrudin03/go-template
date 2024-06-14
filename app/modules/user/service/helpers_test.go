package service

import (
	"context"
	"reflect"
	"testing"
	"time"

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
		beforeFunc func() *UserSrv
		want       *models.User
		wantErr    bool
	}{
		{
			name: "failed params",
			args: args{
				ctx: context.Background(),
			},
			beforeFunc: func() *UserSrv {
				return &UserSrv{}
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

func TestUserSrv_checkEmailIsVerified(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		srv  *UserSrv
		user models.User
		want bool
	}{
		{
			name: "email not verified",
			srv:  &UserSrv{},
			want: false,
		},
		{
			name: "email is verified",
			srv:  &UserSrv{},
			user: models.User{EmailVerifiedAt: &now},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.srv.checkEmailIsVerified(tt.user); got != tt.want {
				t.Errorf("UserSrv.checkEmailIsVerified() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSrv_checkOTP(t *testing.T) {
	otp1 := "1234"
	otp2 := "4321"
	type args struct {
		otp    []byte
		reqOtp string
	}
	tests := []struct {
		name    string
		srv     *UserSrv
		args    args
		wantErr bool
	}{
		{
			name: "otp is not valid",
			srv:  &UserSrv{},
			args: args{
				otp:    []byte(otp1),
				reqOtp: otp2,
			},
			wantErr: true,
		},
		{
			name: "otp is expired or not found",
			srv:  &UserSrv{},
			args: args{
				otp:    []byte(""),
				reqOtp: otp2,
			},
			wantErr: true,
		},
		{
			name: "success",
			srv:  &UserSrv{},
			args: args{
				otp:    []byte(otp2),
				reqOtp: otp2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.srv.checkOTP(tt.args.otp, tt.args.reqOtp); (err != nil) != tt.wantErr {
				t.Errorf("UserSrv.checkOTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
