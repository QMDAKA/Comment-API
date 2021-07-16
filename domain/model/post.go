package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Content  string `gorm:"not null"`
	Comments []Comment
	UserID   uint64 `gorm:"not null"`
	User     User
}
