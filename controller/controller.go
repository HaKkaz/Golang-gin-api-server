package controller

import (
	"context"
	"encoding/json"
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
	// ad.Print()
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

	// Get active ads from PostgreSQL and set it to Redis Cache
	cont.GetActiveAdsForCache()

	c.JSON(http.StatusOK, gin.H{"message": "Advertisement created successfully"})
}

// Get active ads from PostgreSQL
func (cont *AdController) GetActiveAdsForCache() error {
	curTimestamp := time.Now().Unix() // Get current int64 timestamp

	query := cont.gormDB.Model(&Advertisements{})
	query = query.Where("? BETWEEN start_at AND end_at", curTimestamp)

	ads := []Advertisements{}
	query.Select("*").Order("end_at ASC").Find(&ads) // select all ads and order by end_at in ascending order

	// fmt.Println("ads size: ", len(ads))

	adJSON, err := json.Marshal(ads)
	if err != nil {
		return err
	}

	// Set ads to Redis Cache
	errr := cont.redisCache.Set("ads", adJSON, 24*time.Hour).Err()
	if errr != nil {
		return errr
	}

	return nil
}

// Get ads from Redis Cache
func (cont *AdController) GetAds(c *gin.Context) {
	filter := NewAdFilter()
	if err := c.Bind(&filter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

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
	}

	// fmt.Println("size of filteredAds: ", len(filteredAds))

	// Convert Advertisements to AdToadResult
	ads := []ad{}

	offset := filter.Offset
	limit := filter.Limit

	var skipped int32 = 0

	for _, filteredAd := range filteredAds {
		if filter.Age != nil && (filteredAd.AgeStart > *filter.Age || filteredAd.AgeEnd < *filter.Age) {
			continue
		}

		if filter.Country != nil && !isInCountry(*filter.Country, filteredAd.Country) {
			continue
		}

		if filter.Gender != nil && !isInGender(*filter.Gender, filteredAd.Gender) {
			continue
		}

		if filter.Platform != nil && !isInPlatform(*filter.Platform, filteredAd.Platform) {
			continue
		}

		if skipped < offset {
			skipped++
			continue
		}

		ads = append(ads, ad{
			Title:  filteredAd.Title,
			End_At: filteredAd.EndAt,
		})

		if len(ads) == int(limit) {
			break
		}
	}

	response := adResponse{Ads: AdToadResult(ads)}
	c.JSON(http.StatusOK, response)
}
