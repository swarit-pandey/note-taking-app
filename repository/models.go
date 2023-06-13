package repository

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

type Note struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Content string `gorm:"type:text;not null"`
}

type NoteResponse struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}
