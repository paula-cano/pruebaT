package controllers

import (
	"Backend/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	data, err := request.Request()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
