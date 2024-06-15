package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	messageDto "github.com/adamnasrudin03/go-template/app/modules/message/dto"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *UserSrv) SendEmailResetPass(ctx context.Context, userID uint64) (*dto.VerifyOtpRes, error) {
	const opName = "UserService-SendEmailResetPass"
	defer helpers.PanicRecover(opName)
	var (
		keyUser = models.GenerateKeyCacheUserDetail(userID)
		user    = new(models.User)
		err     error
	)

	if ok := srv.RepoCache.GetCache(ctx, keyUser, user); !ok {
		user, err = srv.getDetail(ctx, dto.DetailReq{ID: userID})
		if err != nil {
			srv.Logger.Errorf("%v error check data: %v", opName, err)
			return nil, err
		}
	}

	if ok := srv.checkEmailIsVerified(*user); !ok {
		return nil, helpers.ErrEmailNotVerified()
	}

	resp, err := srv.generateOTP()
	if err != nil {
		srv.Logger.Errorf("%v error generate otp: %v", opName, err)
		return nil, err
	}

	err = srv.RepoMessage.SendEmail(ctx, messageDto.SendEmailReq{
		To:      []string{user.Email},
		Subject: models.SubjectPasswordReset,
		Message: models.GenerateMessagePasswordReset(user.Name, resp.Otp, srv.Cfg.App.OtpExpired),
	})
	if err != nil {
		srv.Logger.Errorf("%v error send email: %v", opName, err)
		return nil, err
	}

	srv.RepoCache.CreateCache(ctx, models.GenerateKeyCacheForgotPassword(userID, resp.RequestID), []byte(resp.Otp), srv.Cfg.App.OtpExpired)
	resp.Otp = ""
	return &resp, nil
}
