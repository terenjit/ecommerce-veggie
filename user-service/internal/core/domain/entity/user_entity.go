package entity

import "time"

type UserEntity struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	RoleName   string    `json:"role_name"`
	Phone      string    `json:"phone"`
	Password   string    `json:"-"`
	Address    string    `json:"address"`
	Photo      string    `json:"photo"`
	Lat        string    `json:"lat"`
	Lng        string    `json:"lng"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"softDelete"`
}
