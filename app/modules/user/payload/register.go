package payload

type RegisterReq struct {
	Name      string `json:"name" validate:"required"`
	Role      string `json:"role"  validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=4"`
	CreatedBy uint64 `json:"created_by"`
}

type RegisterUserReq struct {
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}
