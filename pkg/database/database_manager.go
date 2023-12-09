package database

import (
	"fmt"
)

func DbManager(args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide an action (migrate, create, or drop)")
		return
	}

	action := args[0]

	switch action {
	case "migrate":
		runMigrations()
	case "create":
		createDatabase()
	case "drop":
		dropDatabase()
	default:
		fmt.Println("Invalid action. Please provide migrate, create, or drop")
	}
}
