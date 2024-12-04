package main

import (
	"thepfarrer/url-shortner/database"
	"thepfarrer/url-shortner/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database
	database.InitDB()
	defer database.DB.Close()

	router := routes.SetupRouter()

	router.Run("localhost:8080")
}
