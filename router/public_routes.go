package router

import (
	v1 "github.com/flowista2/api/public/v1"
	"github.com/gin-gonic/gin"
)

func LoadExternalRoutes(engine *gin.Engine) {
	publicGroup := engine.Group("/api")
	{
		v1RouterGroup := publicGroup.Group("/v1")
		{
			var userController v1.UserController
			{
				v1RouterGroup.GET("/users/:id", userController.Show)
				v1RouterGroup.POST("/users", userController.Create)
			}
			
			var loginController v1.LoginController
			{
				v1RouterGroup.POST("/login", loginController.Login)
			}
		}
	}
}
