package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/blake2b"
)

type url struct {
	Key      string `json:"key"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

type urlArgs struct {
	URL string `json:"url" binding:"required"`
}

var db *sql.DB

func getURLs(c *gin.Context) {
	var urls []url

	rows, err := db.Query("SELECT key, long_url, short_url FROM url_model")
	if err != nil {
		log.Fatalf("Error: %q", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var row url
		err = rows.Scan(&row.Key, &row.LongUrl, &row.ShortUrl)
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, row)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, urls)
}

func postURLs(c *gin.Context) {
	var newURL, result url
	var inputURL urlArgs

	if err := c.BindJSON(&inputURL); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newURL.LongUrl = inputURL.URL
	newURL.Key = hashKey(inputURL.URL)

	row := db.QueryRow("SELECT key, long_url FROM url_model WHERE key = ?", newURL.Key)
	err := row.Scan(&result.Key, &result.LongUrl)
	if err != nil {
		newURL.ShortUrl = "http://localhost:8080/" + newURL.Key

		_, err := db.Exec("INSERT INTO url_model (key, long_url, short_url) VALUES (?,?,?)", newURL.Key, newURL.LongUrl, newURL.ShortUrl)
		if err != nil {
			log.Print(err)
			return
		}

		c.IndentedJSON(http.StatusCreated, newURL)
		return
	}

	// for _, a := range urls {
	// 	if a.Key == newURL.Key {
	// 		if a.LongUrl != newURL.LongUrl {
	// 			randomStr := randomString(4)
	// 			newURL.Key = hashKey(inputURL.URL + randomStr)
	// 		} else {
	// 			c.IndentedJSON(http.StatusOK, a)
	// 			return
	// 		}
	// 	}
	// }
	c.IndentedJSON(http.StatusOK, result)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func getURLByKey(c *gin.Context) {
	key := c.Param("key")
	var result url

	row := db.QueryRow("SELECT key, long_url, short_url FROM url_model WHERE key = ?", key)
	err := row.Scan(&result.Key, &result.LongUrl, &result.ShortUrl)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func deleteURLByKey(c *gin.Context) {
	key := c.Param("key")

	_, err := db.Exec("DELETE FROM url_model WHERE key = ?", key)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}

func hashKey(url string) string {
	hash, _ := blake2b.New(4, nil)
	hash.Write([]byte(url))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./test_db.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/api/urls/", getURLs)
	router.GET("/:key", getURLByKey)
	router.DELETE("/:key", deleteURLByKey)
	router.POST("/api/urls/", postURLs)

	router.Run("localhost:8080")
}
