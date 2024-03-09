package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin-api-server/types"
	"go-gin-api-server/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type AdController struct {
	gormDB     *gorm.DB
	redisCache *redis.Client
	ctx        context.Context
}

func NewAdController(gormDB *gorm.DB, redisCache *redis.Client) *AdController {
	return &AdController{
		gormDB:     gormDB,
		redisCache: redisCache,
		ctx:        context.Background(),
	}
}

// Create new ad and store it in PostgreSQL
func (cont *AdController) CreateAd(c *gin.Context) {
	// Bind request body to AdCreationRequesetBody,
	// struct binding will do the validation
	var ad types.AdCreationRequesetBody
	if err := c.BindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle other validation checks and business logic here
	ad.Print()

	adObj := Advertisements{
		Title:    ad.Title,
		StartAt:  utils.DateToTimestamp(ad.StartAt),
		EndAt:    utils.DateToTimestamp(ad.EndAt),
		AgeStart: ad.Conditions.AgeStart,
		AgeEnd:   ad.Conditions.AgeEnd,
		Gender:   ad.Conditions.Gender,
		Platform: ad.Conditions.Platform,
		Country:  ad.Conditions.Country,
	}

	// Create New Ad
	result := cont.gormDB.Create(&adObj)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	cont.GetActiveAds()

	c.JSON(http.StatusOK, gin.H{"message": "Advertisement created successfully"})
}

// Get active ads from PostgreSQL
func (cont *AdController) GetActiveAds() error {
	curTimestamp := time.Now().Unix() // Get current int64 timestamp

	query := cont.gormDB.Model(&Advertisements{})
	query = query.Where("? BETWEEN start_at AND end_at", curTimestamp)

	ads := []Advertisements{}
	query.Select("*").Order("end_at ASC").Find(&ads)

	adJSON, err := json.Marshal(ads)
	if err != nil {
		return err
	}

	fmt.Println(string(adJSON))
	fmt.Println("Get active ads successfully")
	// TODO: Send ads to Redis Cache

	// Set ads to Redis Cache
	errr := cont.redisCache.Set("ads", adJSON, 24*time.Hour).Err()
	if errr != nil {
		return errr
	}

	// val, err := cont.redisCache.Get("ads").Result()
	// if err == redis.Nil {
	// 	fmt.Println("Key 'ads' does not exist")
	// } else if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Value of 'ads' key:\n", val)
	// }

	return nil
}

// Get ads from Redis Cache
func (cont *AdController) GetAds(c *gin.Context) {
	filteredAds := []Advertisements{}

	iter := cont.redisCache.Scan(0, "ads", 0).Iterator()

	// Only "ads" key will be iterated
	for iter.Next() {
		// Get the advertisement list
		val, err := cont.redisCache.Get(iter.Val()).Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = json.Unmarshal([]byte(val), &filteredAds)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// iterate through the list
		for _, ad := range filteredAds {
			fmt.Println("ad.Title = ", ad.Title)
			fmt.Println("ad.StartAt = ", ad.StartAt)
			fmt.Println("ad.EndAt = ", ad.EndAt)
		}
	}

	// Convert Advertisements to AdToadResult
	ads := []ad{}
	for _, filterdAd := range filteredAds {
		ads = append(ads, ad{
			Title:  filterdAd.Title,
			End_At: filterdAd.EndAt,
		})
	}

	response := adResponse{Ads: AdToadResult(ads)}
	c.JSON(http.StatusOK, response)
}
