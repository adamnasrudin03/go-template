package payload

type LoginReq struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=4"`
}

func (req *LoginReq) Validate() error {
	return nil
}

type LoginRes struct {
	Token string `json:"token"`
}
