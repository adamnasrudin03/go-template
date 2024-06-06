package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (srv *userService) getDetail(ctx context.Context, input payload.DetailReq) (*models.User, error) {
	const opName = "UserService-getDetail"
	var (
		res = new(models.User)
		err error
	)

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	res, err = srv.userRepository.GetDetail(ctx, input)
	if err != nil {
		srv.Logger.Errorf("%v error: %v", opName, err)
		return nil, helpers.ErrDB()
	}

	isExist := res != nil && res.ID > 0
	if !isExist {
		return nil, helpers.ErrDataNotFound("Pengguna", "User")
	}

	return res, nil
}

func (srv *userService) checkIsNotDuplicate(ctx context.Context, input payload.DetailReq) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}

	err = srv.checkIsExistEmail(ctx, input)
	if err != nil {
		return err
	}

	err = srv.checkIsExistUsername(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (srv *userService) checkIsExistEmail(ctx context.Context, input payload.DetailReq) (err error) {
	const opName = "UserService-checkIsExistEmail"
	checkUser := new(models.User)
	if len(input.Email) > 0 {
		checkUser, _ = srv.userRepository.GetDetail(ctx, payload.DetailReq{Columns: "id", NotID: input.NotID, Email: input.Email})
		if checkUser != nil && checkUser.ID > 0 {
			srv.Logger.Errorf("%v Email has be registered", opName)
			return helpers.NewError(helpers.ErrConflict, helpers.NewResponseMultiLang(
				helpers.MultiLanguages{
					ID: "Surel Sudah Terdafar",
					EN: "Email Already Registered",
				},
			))
		}
	}

	return nil
}

func (srv *userService) checkIsExistUsername(ctx context.Context, input payload.DetailReq) (err error) {
	const opName = "UserService-checkIsExistUsername"
	checkUser := new(models.User)
	if len(input.Username) > 0 {
		checkUser, _ = srv.userRepository.GetDetail(ctx, payload.DetailReq{Columns: "id", NotID: input.NotID, Username: input.Username})
		if checkUser != nil && checkUser.ID > 0 {
			srv.Logger.Errorf("%v Username has be registered", opName)
			return helpers.NewError(helpers.ErrConflict, helpers.NewResponseMultiLang(
				helpers.MultiLanguages{
					ID: "Username Sudah Terdafar",
					EN: "Username Already Registered",
				},
			))
		}
	}
	return nil
}

func (srv *userService) convertModelsToListResponse(data []models.User) []payload.UserRes {
	var records = []payload.UserRes{}

	for i := 0; i < len(data); i++ {
		temp := payload.UserRes{
			ID:        data[i].ID,
			Name:      data[i].Name,
			Role:      helpers.ToTitle(data[i].Role),
			Username:  data[i].Username,
			Email:     data[i].Email,
			CreatedAt: data[i].CreatedAt,
			UpdatedAt: data[i].UpdatedAt,
		}
		records = append(records, temp)
	}

	return records
}
