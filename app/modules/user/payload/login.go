package payload

type LoginReq struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=4"`
}

type LoginRes struct {
	Token string `json:"token"`
}
