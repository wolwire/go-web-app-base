package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
    ID           int    `json:"id" gorm:"primaryKey"`
    NAME         string `json:"name"`
    EMAIL        string `json:"email" gorm:"type:varchar(100);uniqueIndex:idx_email"`
    PHONE_NUMBER string `json:"phone_number" gorm:"type:varchar(100);uniqueIndex:idx_phone_number"`
    USERNAME     string `json:"username" gorm:"type:varchar(100);uniqueIndex:idx_username"`
    PASSWORD     string
}

func (User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	return user.hashPassword()
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return user.hashPassword()
}

func (user *User) hashPassword() error {
	if user.PASSWORD != "" {
		hashedPassword, err := HashPassword(user.PASSWORD)
		if err != nil {
			return err
		}
		user.PASSWORD = hashedPassword
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(password))
}

