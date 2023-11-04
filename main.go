package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"main.go/database"
	"main.go/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize Database
	connection_string := database.GetPostgresConnectionString()
	database.Connect(connection_string)
	database.Migrate()

	router := routes.InitRouter()
	router.Run(":8000")

}
