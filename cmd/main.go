package main

import (
	"github.com/flowista2/config"
	"github.com/flowista2/pkg/database"
	"github.com/flowista2/pkg/http"
)

var server http.Server

func main() {
	config.LoadConfig()
	server.InitializeServer()
	server.AssignRoutes()
	database.Initialize()
	server.RunServer()
}
