package models

// Log represents the model for an log
type Log struct {
	ID          uint64       `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"not null"`
	Action      string       `json:"action" gorm:"not null"`
	TableNameID uint64       `json:"table_name_id"`
	TableName   string       `json:"table_name"`
	UserID      uint64       `json:"user_id"`
	User        UserRelation `json:"user" gorm:"ForeignKey:UserID"`
	DefaultModel
}

type UserRelation struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

func (UserRelation) TableName() string {
	return "users"
}
