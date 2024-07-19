package service

import (
	"context"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
)

func (srv *UserSrv) VerifiedEmail(ctx context.Context, req *dto.VerifyOtpReq) (err error) {
	const opName = "UserService-VerifiedEmail"
	defer help.PanicRecover(opName)
	var (
		keyUser = models.GenerateKeyCacheUserDetail(req.UserID)
		keyOtp  = models.GenerateKeyCacheOtp(req.UserID, req.RequestID)
		user    = new(models.User)
	)

	err = req.Validate()
	if err != nil {
		return err
	}

	if ok := srv.RepoCache.GetCache(ctx, keyUser, user); !ok {
		user, err = srv.getDetail(ctx, dto.DetailReq{ID: req.UserID})
		if err != nil {
			srv.Logger.Errorf("%v error check data: %v", opName, err)
			return err
		}
	}

	if ok := srv.checkEmailIsVerified(*user); ok {
		return response_mapper.ErrEmailIsVerified()
	}

	temp := []byte("")
	srv.RepoCache.GetCache(ctx, keyOtp, &temp)
	if err = srv.checkOTP(temp, req.Otp); err != nil {
		return err
	}

	now := time.Now()
	err = srv.Repo.UpdateSpecificField(ctx, models.User{
		ID:              user.ID,
		EmailVerifiedAt: &now,
		DefaultModel:    models.DefaultModel{UpdatedBy: req.UserID}})
	if err != nil {
		srv.Logger.Errorf("%v error update data: %v", opName, err)
		return response_mapper.ErrUpdatedDB()
	}

	user.EmailVerifiedAt = &now
	user.UpdatedBy = req.UserID
	go func(usr models.User, params dto.VerifyOtpReq) {
		newCtx := context.Background()
		srv.RepoCache.DelCache(newCtx, models.GenerateKeyCacheOtp(params.UserID, params.RequestID))
		srv.RepoCache.CreateCache(newCtx, models.GenerateKeyCacheUserDetail(usr.ID), usr, time.Minute*5)
	}(*user, *req)

	return nil
}
