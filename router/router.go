package router

import (
	"go-gin-api-server/controller"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func SetupRouter(gormDB *gorm.DB, redisCache *redis.Client) *gin.Engine {
	// Create a new router
	router := gin.New()

	adController := controller.NewAdController(gormDB, redisCache)
	router.POST("/api/v1/ad", adController.CreateAd)
	router.GET("/api/v1/ad", adController.GetAds)
	router.Run(":8080")
	return router
}
