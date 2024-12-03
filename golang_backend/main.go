package main

import (
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/blake2b"
)

type url struct {
	Key      string `json:"key"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

type urlArgs struct {
	URL string `json:"url"`
}

var urls = []url{
	{Key: "a105c5c1", LongUrl: "https://www.example.com", ShortUrl: "http://localhost:8080/a105c5c1"},
}

func getURLs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, urls)
}

func postURLs(c *gin.Context) {
	var newURL url
	var inputURL urlArgs

	if err := c.BindJSON(&inputURL); err != nil {
		return
	}

	newURL.LongUrl = inputURL.URL
	newURL.Key = hashKey(inputURL.URL)
	newURL.ShortUrl = "http://localhost:8080/" + newURL.Key

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

func hashKey(url string) string {
	hash, _ := blake2b.New(4, nil)
	hash.Write([]byte(url))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}

func main() {
	router := gin.Default()
	router.GET("/api/urls/", getURLs)
	router.GET("/:key", getURLByKey)
	router.POST("/api/urls/", postURLs)

	router.Run("localhost:8080")
}
