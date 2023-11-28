package http

import (
	"net/http"
	"time"

	"github.com/flowista2/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	Engine *gin.Engine
	HttpServer *http.Server
}

func (server *Server) InitializeServer() {
	server.Engine = gin.Default()
	server.HttpServer = &http.Server{
		Addr:           viper.GetString("server.Addr"),
		Handler:        server.Engine,
		ReadTimeout:    time.Duration(viper.GetInt("server.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("server.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (server *Server) RunServer() {
	server.HttpServer.ListenAndServe()
}

func (server *Server) AssignRoutes() {
	router.LoadRoutes(server.Engine)
}
