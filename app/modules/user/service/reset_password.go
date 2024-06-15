package service

import (
	"context"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *UserSrv) ResetPassword(ctx context.Context, input *dto.ResetPasswordReq) (err error) {
	const opName = "UserService-ResetPassword"
	defer helpers.PanicRecover(opName)
	var (
		keyUser = models.GenerateKeyCacheUserDetail(input.ID)
		keyOtp  = models.GenerateKeyCacheForgotPassword(input.ID, input.RequestID)
		user    = new(models.User)
	)

	err = input.Validate()
	if err != nil {
		return err
	}

	if ok := srv.RepoCache.GetCache(ctx, keyUser, user); !ok {
		user, err = srv.getDetail(ctx, dto.DetailReq{ID: input.ID})
		if err != nil {
			srv.Logger.Errorf("%v error check data: %v", opName, err)
			return err
		}
	}

	if ok := srv.checkEmailIsVerified(*user); !ok {
		return helpers.ErrEmailNotVerified()
	}

	temp := []byte("")
	srv.RepoCache.GetCache(ctx, keyOtp, &temp)
	if err = srv.checkOTP(temp, input.Otp); err != nil {
		return err
	}

	newPass, err := helpers.HashPassword(input.NewPassword)
	if err != nil {
		srv.Logger.Errorf("%v error hash password: %v ", opName, err)
		return helpers.ErrHashPasswordFailed()
	}
	user.UpdatedBy = input.UpdatedBy
	user.Password = newPass

	user, err = srv.Repo.Updates(ctx, *user)
	if err != nil {
		srv.Logger.Errorf("%v error: %v ", opName, err)
		return helpers.ErrUpdatedDB()
	}

	go func(usr models.User, params dto.ResetPasswordReq) {
		newCtx := context.Background()
		srv.RepoCache.DelCache(newCtx, models.GenerateKeyCacheForgotPassword(params.ID, params.RequestID))
		srv.RepoCache.CreateCache(newCtx, models.GenerateKeyCacheUserDetail(usr.ID), usr, time.Minute*5)
	}(*user, *input)

	return nil
}
