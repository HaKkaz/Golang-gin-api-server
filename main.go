package main

import (
	"go-gin-api-server/router"
	database "go-gin-api-server/storage"

	_ "github.com/lib/pq" // Import PostgreSQL driverad
)

func main() {
	db := database.ConnectDatabase()
	cache := database.ConnectCache()
	defer db.Close()
	gormDB := database.ConnectORM(db)
	router.SetupRouter(gormDB, cache)
}
