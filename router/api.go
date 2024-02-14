package router

import (
	"go-gin-api-server/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	m := map[string]string{
		"message": "pong",
	}
	c.JSON(http.StatusOK, m)
}

func create_advertisement(c *gin.Context) {
	var ad types.AdCreationRequesetBody
	if err := c.BindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle other validation checks and business logic here
	ad.Print()

	c.JSON(http.StatusOK, gin.H{"message": "Advertisement created successfully"})
}
