package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

var DB Db

func Initialize() Db {
	db := connectDatabase()
	dbname := viper.GetString("database.dbname")

	err := db.Exec(fmt.Sprintf("USE %s", dbname)).Error
	if err != nil {
		panic(fmt.Errorf("failed to create database: %w", err))
	}
	DB = Db{DB: db}
	return DB
}
