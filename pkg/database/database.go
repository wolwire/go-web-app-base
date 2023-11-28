package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Db struct{
	*gorm.DB
}

var DB Db

func Initialize() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = Db{DB: db}
}
