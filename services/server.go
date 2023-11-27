package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flowista2/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitializeServer() {
	fmt.Println("Starting Server............")
	ginHandler := gin.Default()
	fmt.Println("Loading Routes.............")
	server := &http.Server{
		Addr:           viper.GetString("server.Addr"),
		Handler:        ginHandler,
		ReadTimeout:    time.Duration(viper.GetInt("server.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("server.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.LoadRoutes(&server)
	server.ListenAndServe()
	fmt.Println("Stopping Server............")
}
