package dto

import "time"

type UserRes struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
