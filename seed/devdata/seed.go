package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/QMDAKA/comment-mock/domain/model"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Comment{})

	// Create
	db.Create(&model.User{
		UUID: "3oxsad",
		Post: []model.Post{
			{
				Content: "lorem ipsum dolor sit amet",
			},
		},
	})

}
