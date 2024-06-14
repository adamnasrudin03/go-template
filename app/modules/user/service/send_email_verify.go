package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *UserSrv) SendEmailVerify(ctx context.Context, userID uint64) (*dto.VerifyOtpRes, error) {
	const opName = "UserService-SendEmailVerify"
	defer helpers.PanicRecover(opName)
	var (
		keyUser = models.GenerateKeyCacheUserDetail(userID)
		user    = new(models.User)
		err     error
	)

	srv.RepoCache.GetCache(ctx, keyUser, user)
	useCache := user != nil && user.ID > 0
	if !useCache {
		user, err = srv.getDetail(ctx, dto.DetailReq{ID: userID})
		if err != nil {
			srv.Logger.Errorf("%v error check data: %v", opName, err)
			return nil, err
		}
	}

	if ok := srv.checkEmailIsVerified(*user); ok {
		return nil, helpers.ErrEmailIsVerified()
	}

	resp, err := srv.generateOTP()
	if err != nil {
		srv.Logger.Errorf("%v error generate otp: %v", opName, err)
		return nil, err
	}

	keyOtp := models.GenerateKeyCacheOtp(userID, resp.RequestID)
	srv.RepoCache.CreateCache(ctx, keyOtp, []byte(resp.Otp), srv.Cfg.App.OtpExpired)
	srv.Logger.Infof("request id: %v => otp: %v", resp.RequestID, resp.Otp)
	// err = srv.RepoMessage.SendEmail(ctx, messageDto.SendEmailReq{
	// 	To:      []string{user.Email},
	// 	Subject: models.SubjectEmailVerify,
	// 	Message: "Your OTP: " + otp,
	// })
	// if err != nil {
	// 	srv.Logger.Errorf("%v error send email: %v", opName, err)
	// 	return nil, err
	// }

	resp.Otp = ""
	return &resp, nil
}
