package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=amstrongDB port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = database

	DB.AutoMigrate(&User{}, &ArmstrongNumber{})

	r := mux.NewRouter()
	RegisterRoutes(r, DB)

	// Enable CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // allow all origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
