package routes

import (
	"thepfarrer/url-shortner/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/api/urls/", handlers.GetURLs)
	router.GET("/:key", handlers.GetURLByKey)
	router.DELETE("/:key", handlers.DeleteURLByKey)
	router.POST("/api/urls/", handlers.PostURLs)

	return router
}
