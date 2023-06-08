package dbaccess

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/piabirajdar/Golang-Gorm-Database/types"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
	return db
}

func Migrate() error {
	db := InitDB()
	return db.AutoMigrate(&types.Author{})
}
