package controller

import (
	"go-gin-api-server/types"
	"go-gin-api-server/utils"
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

	adObj := Advertisements{
		Title:     ad.Title,
		Start_At:  utils.DateToTimestamp(ad.StartAt),
		End_At:    utils.DateToTimestamp(ad.EndAt),
		Age_Start: ad.Conditions.AgeStart,
		Age_End:   ad.Conditions.AgeEnd,
		Gender:    ad.Conditions.Gender,
		Platform:  ad.Conditions.Platform,
		Country:   ad.Conditions.Country,
	}

	// Create New Ad
	result := cont.gormDB.Create(&adObj)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Advertisement created successfully"})
}
