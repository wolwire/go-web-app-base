package main

import (
	"github.com/flowista2/config"
	"github.com/flowista2/pkg/database"
	"github.com/flowista2/pkg/http"
)

func main() {
	config.LoadConfig()
	var server http.Server
	server.InitializeServer()
	server.AssignRoutes()
	database.Initialize()
	server.RunServer()
}
