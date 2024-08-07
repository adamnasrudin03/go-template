package repository

import (
	"context"
	"time"

	help "github.com/adamnasrudin03/go-helpers"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *userRepo) UpdateSpecificField(ctx context.Context, input models.User) (err error) {
	const opName = "UserRepository-UpdateSpecificField"
	update := map[string]interface{}{}

	if input.Email != "" {
		update["email"] = input.Email
	}
	if input.Username != "" {
		update["username"] = input.Username
	}
	if input.Name != "" {
		update["name"] = input.Name
	}
	if input.Role != "" {
		update["role"] = input.Email
	}
	if input.UpdatedBy > 0 {
		update["updated_by"] = input.UpdatedBy
	}

	if !help.CheckTimePointerValue(input.EmailVerifiedAt).IsZero() {
		update["email_verified_at"] = input.EmailVerifiedAt
	}

	update["updated_at"] = time.Now()
	err = r.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", input.ID).Updates(update).Error
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return err
	}

	return nil
}
