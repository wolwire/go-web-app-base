package v1

import (
	"net/http"
	"strconv"

	"github.com/flowista2/models"
	"github.com/flowista2/pkg/database"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

/* 
Login is a controller function that handles user login.
It fetches user data from the gin context, such as username and password,
and compares the hashed password with the stored hashed password in the database.
If the passwords match, it sets a session cookie and returns a success response, otherwise, it returns an error response.
*/
func (user_controller *LoginController) Login(c *gin.Context) {
	var user models.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid username or password"})
		return
	}
	err := user.ComparePassword(password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    strconv.Itoa(int(user.ID)),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(c.Writer, sessionCookie)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

