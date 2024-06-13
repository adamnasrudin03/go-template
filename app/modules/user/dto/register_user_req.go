package dto

type RegisterUserReq struct {
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=4"`
}

func (m *RegisterUserReq) ConvertToRegisterReq() RegisterReq {
	return RegisterReq{
		Name:     m.Name,
		Role:     m.Role,
		Email:    m.Email,
		Username: m.Username,
		Password: m.Password,
	}
}
