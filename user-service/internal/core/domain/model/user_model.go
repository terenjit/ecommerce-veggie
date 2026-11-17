package model

import "time"

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Photo      string    `json:"photo"`
	Address    string    `json:"address"`
	Lat        string    `json:"lat"`
	Lng        string    `json:"lng"`
	IsVerified bool      `json:"is_verified"`
	RoleID     int64     `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	Role       Role      `json:"role"`
}

func (User) TableName() string {
	return "users"
}
