package models

import (
	"github.com/flowista2/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
    ID           int    `json:"id"`
    NAME         string `json:"name"`
    EMAIL        string `json:"email"`
    PHONE_NUMBER string `json:"phone_number"`
    USERNAME     string `json:"username"`
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
		hashedPassword, err := pkg.HashPassword(user.PASSWORD)
		if err != nil {
			return err
		}
		user.PASSWORD = hashedPassword
	}
	return nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(password))
}

