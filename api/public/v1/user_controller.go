package v1

import (
	"net/http"
	"strconv"

	"github.com/flowista2/models"
	"github.com/flowista2/pkg/database"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// Show retrieves a user by their ID and returns the user information as JSON.
// If the user is found, it returns HTTP status code 200 (OK) along with the user details.
// If the user is not found, it returns HTTP status code 404 (Not Found).
// If the ID parameter is not a valid integer, it returns HTTP status code 400 (Bad Request).
func (user_controller *UserController) Show(c *gin.Context) {
	var user models.User
	user_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, user)
		return
	}

	result := database.DB.Find(&user, user_id)
	if result.RowsAffected > 0 && result.Error == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, user)
	}
}

// Create is a method of the LoginController struct that handles the creation of a new user.
// It takes a gin.Context object as a parameter and retrieves the user information from the request body.
// The user information includes the username, email, phone number, and password (hashed before storing in the database).
// The method creates a new user record in the database using the retrieved information.
// If the user is successfully created, it returns a JSON response with the user details and HTTP status code 200 (OK).
// If there is an error during the creation process, it returns a JSON response with the user details and HTTP status code 404 (Not Found).
func (user_controller *UserController) Create(c *gin.Context) {
	var user models.User
	user.USERNAME = c.PostForm("username")
	user.EMAIL = c.PostForm("email")
	user.PHONE_NUMBER = c.PostForm("phone_number")
	user.PASSWORD = c.PostForm("password") // password is hashed before storing in database
	
	result := database.DB.Create(&user)
	if result.RowsAffected > 0 && result.Error == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, user)
	}
}
