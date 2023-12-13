package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(engine *gin.Engine) {
	fmt.Println("Loading Routes.............")
	err := LoadPublicRoutes(engine)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error loading external routes: %w", err))
	}
	err = LoadInternalRoutes(engine)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error loading internal routes: %w", err))
	}
	fmt.Println("Routes Loaded..............")
}
