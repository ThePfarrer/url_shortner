package main

import (
	"thepfarrer/url-shortner/database"
	"thepfarrer/url-shortner/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database
	database.InitDB()
	defer database.DB.Close()

	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/api/urls/", handlers.GetURLs)
	router.GET("/:key", handlers.GetURLByKey)
	router.DELETE("/:key", handlers.DeleteURLByKey)
	router.POST("/api/urls/", handlers.PostURLs)

	router.Run("localhost:8080")
}
