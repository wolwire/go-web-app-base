package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

var DB Db

func Initialize() {
	db := connectDatabase()
	dbname := viper.GetString("database.dbname")

	err := db.Exec(fmt.Sprintf("USE %s", dbname)).Error
	if err != nil {
		panic(fmt.Errorf("failed to create database: %w", err))
	}
	DB = Db{DB: db}
}

func createDatabase() {
	db := connectDatabase()
	dbname := viper.GetString("database.dbname")
	err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname)).Error
	if err != nil {
		panic(fmt.Errorf("failed to create database: %w", err))
	}
	fmt.Println("Database created")
}

func dropDatabase() {
	db := connectDatabase()
	dbname := viper.GetString("database.dbname")	
	err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname)).Error
	if err != nil {
		panic(fmt.Errorf("failed to drop database: %w", err))
	}
	fmt.Println("Database dropped")
}

func connectDatabase() *gorm.DB {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port) 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}
	return db
}

