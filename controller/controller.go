package controller

import (
	"go-gin-api-server/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdController struct {
	gormDB *gorm.DB
}

func NewAdController(gormDB *gorm.DB) *AdController {
	return &AdController{gormDB: gormDB}
}

func (cont *AdController) CreateAd(c *gin.Context) {
	var ad types.AdCreationRequesetBody
	if err := c.BindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle other validation checks and business logic here
	ad.Print()

	c.JSON(http.StatusOK, gin.H{"message": "Advertisement created successfully"})
}
