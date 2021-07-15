package app

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/QMDAKA/comment-mock/registry"
)

func Serve() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	s, err := registry.InitializeServer(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize server: %s", err))
	}
	defer s.Close()

	s.Handler()
	s.Start()

}
