package handlers

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
)

func ParseParams(c *gin.Context) (map[string]interface{}, error) {
    var params map[string]interface{}
    err := json.NewDecoder(c.Request.Body).Decode(&params)
    return params, err
}