package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(engine *gin.Engine) {
	fmt.Println("Loading Routes.............")
	err := LoadExternalRoutes(engine)
	if err != nil {
		panic(fmt.Errorf("fatal error loading external routes: %w", err))
	}
	fmt.Println("Routes Loaded..............")
}
