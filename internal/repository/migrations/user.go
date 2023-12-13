package migrations

import (
	"gorm.io/gorm"
)

type User struct {
    ID           int    `gorm:"primaryKey"`
    NAME         string 
    EMAIL        string `gorm:"type:varchar(100);uniqueIndex:idx_email"`
    PHONE_NUMBER string `gorm:"type:varchar(100);uniqueIndex:idx_phone_number"`
    USERNAME     string `gorm:"type:varchar(100);uniqueIndex:idx_username"`
    PASSWORD     string
}

func userMigrateUp(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

