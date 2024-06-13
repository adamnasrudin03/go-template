package dto

type LoginReq struct {
	Username string `json:"username" validate:"required,min=4"` // username or email
	Password string `json:"password" validate:"required,min=4"`
}
