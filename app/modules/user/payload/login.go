package payload

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

func (req *LoginReq) Validate() error {
	return nil
}

type LoginRes struct {
	Token string `json:"token"`
}
