package database

import (
	"fmt"

	"github.com/flowista2/models"
)

func runMigrations() {
	Initialize()
	fmt.Println("Running migrations...")
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Errorf("failed to run migrations: %w", err))
	}
	fmt.Println("Migrations complete")
}
