package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(engine *gin.Engine) {
	fmt.Println("Loading Routes.............")
	LoadExternalRoutes(engine)
	fmt.Println("Routes Loaded..............")
}
