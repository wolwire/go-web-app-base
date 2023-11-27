package main

import (
	"fmt"
	"github.com/flowista2/config"
	"github.com/flowista2/services"
)

func main() {
	fmt.Println("Starting server........")
	config.LoadConfig()
	services.InitializeServer()
}
