package router

import (
	"go-gin-api-server/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(gormDB *gorm.DB) *gin.Engine {
	// Create a new router
	router := gin.New()

	adController := controller.NewAdController(gormDB)
	router.POST("/api/v1/ad", adController.CreateAd)
	router.GET("/api/v1/ad", adController.GetAds)
	router.Run(":8080")
	return router
}
