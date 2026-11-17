package model

import "time"

type UserRole struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (UserRole) TableName() string {
	return "user_role"
}
