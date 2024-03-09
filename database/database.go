package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *sql.DB {
	// Load env file and get the environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: load env file error.")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to PostgreSQL database
	dbPath := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		panic(err)
	}

	// Check if the table exists, if so, drop it
	_, err = db.Exec("DROP TABLE IF EXISTS Advertisements")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Drop table success.")
	}

	// Testing to Create Table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Advertisements (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			start_at BIGINT NOT NULL,
			end_at BIGINT NOT NULL,
			age_start INT,
			age_end INT,
			gender varchar(1)[],
			country varchar(2)[],
			platform varchar(10)[]
		)
	`)

	if err != nil {
		fmt.Println("Error: create table error.")
		fmt.Println(err)
	} else {
		fmt.Println("Create table success.")
	}

	// Create indexes
	_, err = db.Exec("CREATE INDEX idx_gender ON Advertisements USING GIN (gender)")
	if err != nil {
		fmt.Println("Error: create gender index error.")
		fmt.Println(err)
	}

	_, err = db.Exec("CREATE INDEX idx_country ON Advertisements USING GIN (country)")
	if err != nil {
		fmt.Println("Error: create country index error.")
		fmt.Println(err)
	}

	_, err = db.Exec("CREATE INDEX idx_platform ON Advertisements USING GIN (platform)")
	if err != nil {
		fmt.Println("Error: create platform index error.")
		fmt.Println(err)
	}

	return db
}

func ConnectORM(db *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	return gormDB
}
