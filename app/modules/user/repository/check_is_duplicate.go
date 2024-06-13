package repository

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/user/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (r *userRepo) CheckIsDuplicate(ctx context.Context, input dto.DetailReq) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}

	err = r.checkIsExistEmail(ctx, input)
	if err != nil {
		return err
	}

	err = r.checkIsExistUsername(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
func (r *userRepo) checkIsExistEmail(ctx context.Context, input dto.DetailReq) (err error) {
	const opName = "UserRepository-checkIsExistEmail"
	checkUser := new(models.User)
	if len(input.Email) > 0 {
		checkUser, _ = r.GetDetail(ctx, dto.DetailReq{Columns: "id", NotID: input.NotID, Email: input.Email})
		if checkUser != nil && checkUser.ID > 0 {
			r.Logger.Errorf("%v Email has be registered", opName)
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

func (r *userRepo) checkIsExistUsername(ctx context.Context, input dto.DetailReq) (err error) {
	const opName = "UserRepository-checkIsExistUsername"
	checkUser := new(models.User)
	if len(input.Username) > 0 {
		checkUser, _ = r.GetDetail(ctx, dto.DetailReq{Columns: "id", NotID: input.NotID, Username: input.Username})
		if checkUser != nil && checkUser.ID > 0 {
			r.Logger.Errorf("%v Username has be registered", opName)
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
