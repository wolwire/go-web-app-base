package routes

import (
	"errors"

	v1 "github.com/flowista2/internal/http/controllers/public/v1"
	"github.com/gin-gonic/gin"
)

func LoadPublicRoutes(engine *gin.Engine) error {
	publicGroup := engine.Group("/api")
	{
		v1RouterGroup := publicGroup.Group("/v1")
		{
			var userController v1.UserController
			if &userController == nil {
                return errors.New("Failed to initialize UserController")
            }
			{
				v1RouterGroup.GET("/users/:id", userController.Show)
				v1RouterGroup.POST("/users", userController.Create)
			}

			var loginController v1.LoginController
            if &loginController == nil {
                return errors.New("Failed to initialize LoginController")
            }
			{
				v1RouterGroup.POST("/login", loginController.Login)
			}
		}
	}
	return	nil
}
