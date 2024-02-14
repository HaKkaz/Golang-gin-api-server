package router

import (
	"database/sql"
	"go-gin-api-server/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	// Create a new router
	router := gin.New()

	adController := controller.NewAdController(db)
	router.GET("/ping", ping)
	router.POST("/api/v1/ad", adController.CreateAd)
	router.Run(":8080")
	return router
}
