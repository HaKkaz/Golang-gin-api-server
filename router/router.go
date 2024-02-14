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
	router.GET("/ping", ping)
	router.POST("/api/v1/ad", adController.CreateAd)
	router.Run(":8080")
	return router
}
