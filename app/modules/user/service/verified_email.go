package service

import (
	"context"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *UserSrv) VerifiedEmail(ctx context.Context, req *dto.VerifyOtpReq) (err error) {
	const opName = "UserService-VerifiedEmail"
	defer helpers.PanicRecover(opName)
	var (
		keyUser = models.GenerateKeyCacheUserDetail(req.UserID)
		keyOtp  = models.GenerateKeyCacheOtp(req.UserID, req.RequestID)
		user    = new(models.User)
	)

	err = req.Validate()
	if err != nil {
		return err
	}

	srv.RepoCache.GetCache(ctx, keyUser, user)
	useCache := user != nil && user.ID > 0
	if !useCache {
		user, err = srv.getDetail(ctx, dto.DetailReq{ID: req.UserID})
		if err != nil {
			srv.Logger.Errorf("%v error check data: %v", opName, err)
			return err
		}
	}

	if ok := srv.checkEmailIsVerified(*user); ok {
		return helpers.ErrEmailIsVerified()
	}

	temp := []byte("")
	srv.RepoCache.GetCache(ctx, keyOtp, &temp)
	if err = srv.checkOTP(temp, req.Otp); err != nil {
		return err
	}

	now := time.Now()
	user.EmailVerifiedAt = &now
	err = srv.Repo.UpdateSpecificField(ctx, models.User{ID: user.ID, EmailVerifiedAt: &now})
	if err != nil {
		srv.Logger.Errorf("%v error update data: %v", opName, err)
		return helpers.ErrUpdatedDB()
	}

	go srv.RepoCache.CreateCache(context.Background(), keyUser, user, time.Minute*5)
	return nil
}
