package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	UUID string `gorm:"not null"`
	Post []Post
}
