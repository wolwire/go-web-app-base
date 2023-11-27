package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(engine *gin.Engine){
	fmt.Println("Loading Routes.............")
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("Routes Loaded..............")
}