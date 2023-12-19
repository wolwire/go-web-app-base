package routes

import (
	v1 "github.com/flowista2/internal/http/controllers/public/v1"
	"github.com/flowista2/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func LoadPublicRoutes(engine *gin.Engine) error {
	publicGroup := engine.Group("/api")
	{
		v1RouterGroup := publicGroup.Group("/v1")
		{
			var userController v1.UserController
			{
				v1RouterGroup.GET("/users/details", middleware.UserAuth(), userController.Show)
				v1RouterGroup.POST("/users", userController.Create)
			}

			var loginController v1.LoginController
			{
				v1RouterGroup.POST("/login", loginController.Login)
			}
		}
	}
	return	nil
}
