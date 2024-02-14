package main

import (
	"go-gin-api-server/database"
	"go-gin-api-server/router"

	_ "github.com/lib/pq" // Import PostgreSQL driverad
)

func main() {
	db := database.Connect()
	defer db.Close()
	gormDB := database.ConnectORM(db)
	router.SetupRouter(gormDB)
}
