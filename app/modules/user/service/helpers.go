package service

import (
	"context"
	"strconv"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/google/uuid"
)

func (srv *UserSrv) getDetail(ctx context.Context, input dto.DetailReq) (*models.User, error) {
	const opName = "UserService-getDetail"
	var (
		res = new(models.User)
		err error
	)

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	res, err = srv.Repo.GetDetail(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v", opName, err)
		return nil, response_mapper.ErrDB()
	}

	isExist := res != nil && res.ID > 0
	if !isExist {
		return nil, response_mapper.ErrDataNotFound("Pengguna", "User")
	}

	if input.Columns == "" {
		go srv.RepoCache.CreateCache(context.Background(), models.GenerateKeyCacheUserDetail(res.ID), res, time.Minute*5)
	}

	return res, nil
}

func (srv *UserSrv) convertModelsToListResponse(data []models.User) []dto.UserRes {
	var records = []dto.UserRes{}

	for i := 0; i < len(data); i++ {
		temp := dto.UserRes{
			ID:        data[i].ID,
			Name:      data[i].Name,
			Role:      help.ToTitle(data[i].Role),
			Username:  data[i].Username,
			Email:     data[i].Email,
			CreatedAt: data[i].CreatedAt,
			UpdatedAt: data[i].UpdatedAt,
		}
		records = append(records, temp)
	}

	return records
}

func (srv *UserSrv) generateOTP() (resp dto.VerifyOtpRes, err error) {
	const opName = "UserService-generateOTP"
	reqID, err := uuid.NewV7()
	if err != nil {
		srv.Logger.Errorf("%v error generate uuid: %v", opName, err)
		return dto.VerifyOtpRes{}, response_mapper.ErrGenerateOtp()
	}

	resp = dto.VerifyOtpRes{
		RequestID: reqID.String(),
		Otp:       strconv.Itoa(help.GenerateRandomNumber(srv.Cfg.App.OtpLength)),
	}

	if resp.Otp == "" {
		return dto.VerifyOtpRes{}, response_mapper.ErrGenerateOtp()
	}

	return resp, nil
}

func (srv *UserSrv) checkOTP(otp []byte, reqOtp string) error {
	sourceOtp := string(otp)
	if sourceOtp == "" {
		return response_mapper.ErrOtpExpired()
	}

	if sourceOtp != reqOtp {
		return response_mapper.ErrOtpInvalid()
	}

	return nil
}

func (srv *UserSrv) checkEmailIsVerified(user models.User) bool {
	return !help.CheckTimePointerValue(user.EmailVerifiedAt).IsZero()
}
