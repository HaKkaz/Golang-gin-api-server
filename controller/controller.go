package controller

import (
	"go-gin-api-server/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
		Start_At:  ad.StartAt,
		End_At:    ad.EndAt,
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

type StringArray []string

func (a *StringArray) Scan(src interface{}) error {
	decoder := pq.Array(&a)
	return decoder.Scan(src)
}

func (a StringArray) Value() (interface{}, error) {
	return pq.StringArray(a), nil
}
