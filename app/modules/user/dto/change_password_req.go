package dto

import (
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
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
		return response_mapper.ErrIsRequired("ID Pengguna", "User ID")
	}

	if len(req.Password) == 0 {
		return response_mapper.ErrIsRequired("Kata sandi", "Password")
	}

	if req.NewPassword != req.ConfirmPassword {
		return response_mapper.ErrNewPasswordNotMatchWithConfirmPassword()
	}
	if req.UpdatedBy == 0 {
		req.UpdatedBy = req.ID
	}
	return nil
}
