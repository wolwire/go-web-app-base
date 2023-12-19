package middleware

import (
	"net/http"

	"github.com/flowista2/internal/repository"
	"github.com/flowista2/models"
	"github.com/flowista2/pkg/cookie"
	"github.com/flowista2/pkg/database"
	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Fetch user ID from session cookie
        session_cookie, err := c.Request.Cookie("session")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Extract user ID from cookie
		session_user, err := cookie.SessionUser(session_cookie)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Fetch user from database
        var user models.User
        repository.UserRep(&database.DB).Find(session_user.ID, &user)
        if &user == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Store user in context
        c.Set("current_user", user)

        c.Next()
    }
}