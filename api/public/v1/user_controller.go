package v1

import (
	"net/http"
	"strconv"

	"github.com/flowista2/models"
	"github.com/flowista2/pkg/database"
	"github.com/gin-gonic/gin"
)

type UserController struct{

}

func (user_controller *UserController) Index(c *gin.Context){
	var users []models.User
	result := database.DB.Find(&users)
	if result.RowsAffected > 0 && result.Error == nil {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, []models.User{})
	}
}

func (user_controller *UserController) Show(c *gin.Context){
	var user models.User
	user_id, err := strconv.Atoi(c.Param("id"))
	if err!=nil{
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