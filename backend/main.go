package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Database connection string
	dsn := "host=localhost user=postgres password=1234 dbname=amstrongDB port=5432 sslmode=disable"

	// Connect DB
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = database

	// Migrate tables
	DB.AutoMigrate(&User{}, &ArmstrongNumber{})

	// Initialize router
	r := mux.NewRouter()
	RegisterRoutes(r, DB)

	// Start server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
