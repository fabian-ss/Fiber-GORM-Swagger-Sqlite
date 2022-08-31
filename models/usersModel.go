package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	FirtsName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
}

type UserResponse struct {
	FirtsName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
