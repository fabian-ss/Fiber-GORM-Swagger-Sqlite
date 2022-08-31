package models

import "time"

type Nota struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"-"`
	UserRefer   int       `json:"user_id"`
	User        User      `gorm:"foreignKey:UserRefer"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type NotaResponse struct {
	User        User   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
