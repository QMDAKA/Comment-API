package mysql

import (
	"gorm.io/gorm"
)

type Post struct {
	db *gorm.DB
}

func ProvidePostRepo(db *gorm.DB) *Post {
	return &Post{db: db}
}

func (p *Post) GetByID() {

}