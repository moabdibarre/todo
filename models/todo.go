package models

import "time"

type ToDoList struct {
	ID        string   `gorm:"primaryKey"`
	Name      string `json:"name"`
	UserID    string  `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
