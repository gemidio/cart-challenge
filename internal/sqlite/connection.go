package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	var err error
	database, err = gorm.Open(sqlite.Open("./build/docker/sqlite/database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
