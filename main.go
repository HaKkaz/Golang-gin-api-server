package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.GET("/ping", ping)
	router.POST("/api/v1/ad", create_advertisement)
	router.Run(":8080")
}
