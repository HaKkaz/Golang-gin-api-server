package controller

import (
	"fmt"
	"go-gin-api-server/types"
	"go-gin-api-server/utils"
	"net/http"
	"time"

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

func (cont *AdController) GetAds(c *gin.Context) {
	filter := NewAdFilter()
	if err := c.Bind(&filter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	curTimestamp := time.Now().Unix()

	query := cont.gormDB.Model(&Advertisements{})

	query = query.Where("? BETWEEN start_at AND end_at", curTimestamp)

	if filter.Age != nil {
		query = query.Where("? BETWEEN age_start AND age_end", *filter.Age)
	}

	if filter.Gender != nil {
		query = query.Where("? = ANY(gender) OR array_length(gender, 1) IS NULL", *filter.Gender)
	}

	if filter.Country != nil {
		query = query.Where("? = ANY(country) OR array_length(country, 1) IS NULL", *filter.Country)
	}

	if filter.Platform != nil {
		query = query.Where("? = ANY(platform) OR array_length(platform, 1) IS NULL", *filter.Platform)
	}

	ads := []ad{}

	filter.check()
	err := query.Select("title, end_at").Order("end_at ASC").Offset(int(filter.Offset)).Limit(int(filter.Limit)).Find(&ads).Error

	fmt.Println(query.Statement.SQL.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := adResponse{Ads: AdToadResult(ads)}
	c.JSON(http.StatusOK, response)
}
