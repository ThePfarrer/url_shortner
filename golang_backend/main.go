package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type url struct {
	ID       string `json:"id"`
	Key      string `json:"key"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

var urls = []url{
	{ID: "1", Key: "a105c5c1", LongUrl: "https://www.example.com", ShortUrl: "http://localhost:8080/a105c5c1"},
}

func getURLs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, urls)
}

func postURLs(c *gin.Context) {
	var newURL url

	if err := c.BindJSON(&newURL); err != nil {
		return
	}

	urls = append(urls, newURL)
	c.IndentedJSON(http.StatusCreated, newURL)
}

func getURLByKey(c *gin.Context) {
	key := c.Param("key")

	for _, a := range urls {
		if a.Key == key {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
}

func main() {
	router := gin.Default()
	router.GET("/api/urls/", getURLs)
	router.GET("/:key", getURLByKey)
	router.POST("/api/urls/", postURLs)

	router.Run("localhost:8080")
}
