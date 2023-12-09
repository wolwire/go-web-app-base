package main

import (
	"os"
	"fmt"
	"github.com/flowista2/config"
	"github.com/flowista2/pkg/database"
	"github.com/flowista2/pkg/http"
)

var server http.Server

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an action (server, db)")
		return
	}
	config.LoadConfig()
	action := os.Args[1]

	switch action {
	case "server":
		startServer()
	case "db":
		database.DbManager(os.Args[2:])
	default:
		fmt.Println("Invalid action. Please provide server, db")
	}
}

func startServer() {
		fmt.Println("Starting the server...")
		server.InitializeServer()
		server.AssignRoutes()
		database.Initialize()
		server.RunServer()
}
