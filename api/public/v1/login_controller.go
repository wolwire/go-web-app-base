package v1

import (
	"net/http"

	"github.com/flowista2/models"
	"github.com/flowista2/pkg/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
}

/* 
Login is a controller function that handles user login.
It fetches user data from the gin context, such as username and password,
and compares the hashed password with the stored hashed password in the database.
If the passwords match, it returns a success response, otherwise, it returns an error response.
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
	err := bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
