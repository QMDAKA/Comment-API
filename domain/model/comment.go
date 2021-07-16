package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Content string `gorm:"not null"`
	PostID  uint64 `gorm:"not null"`
	Post    Post
	UserID  uint64 `gorm:"not null"`
	User    User
}
