package dto

import (
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type ChangePasswordReq struct {
	ID              uint64 `json:"id"`
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
	UpdatedBy       uint64 `json:"updated_by"`
}

func (req *ChangePasswordReq) Validate() error {
	if req.ID == 0 {
		return helpers.ErrIsRequired("ID Pengguna", "User ID")
	}

	if len(req.Password) == 0 {
		return helpers.ErrIsRequired("Kata sandi", "Password")
	}

	if req.NewPassword != req.ConfirmPassword {
		return helpers.ErrNewPasswordNotMatchWithConfirmPassword()
	}
	if req.UpdatedBy == 0 {
		req.UpdatedBy = req.ID
	}
	return nil
}
