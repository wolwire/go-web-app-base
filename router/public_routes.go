package router

import (
	v1 "github.com/flowista2/api/public/v1"
	"github.com/gin-gonic/gin"
)

func LoadExternalRoutes(engine *gin.Engine) {
	public_group := engine.Group("/api")
	{
		v1_router_group := public_group.Group("/v1")
		{
			var user_controller v1.UserController
			{
				v1_router_group.GET("/users", user_controller.Index)
				v1_router_group.GET("/users/:id", user_controller.Show)
			}
		}
	}
}
