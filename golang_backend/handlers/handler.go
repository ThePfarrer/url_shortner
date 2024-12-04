package handlers

import (
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"

	"thepfarrer/url-shortner/database"
	"thepfarrer/url-shortner/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/blake2b"
)

func GetURLs(c *gin.Context) {
	var urls []models.URL

	rows, err := database.DB.Query("SELECT key, long_url, short_url FROM url_model")
	if err != nil {
		log.Fatalf("Error: %q", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var row models.URL
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

func PostURLs(c *gin.Context) {
	var inputURL models.UrlArgs

	if err := c.BindJSON(&inputURL); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	longUrl := inputURL.URL
	baseUrl := "http://localhost:8080/"

	key := hashKey(inputURL.URL)
	var result models.URL

	row := database.DB.QueryRow("SELECT key, long_url, short_url FROM url_model WHERE key = ?", key)
	err := row.Scan(&result.Key, &result.LongUrl, &result.ShortUrl)
	if err == nil {
		if longUrl == result.LongUrl {
			c.IndentedJSON(http.StatusOK, result)
			return

		}
		key = hashKey(inputURL.URL + randomString(4))
	}

	shortUrl := baseUrl + key

	_, err = database.DB.Exec("INSERT INTO url_model (key, long_url, short_url) VALUES (?,?,?)", key, longUrl, shortUrl)
	if err != nil {
		log.Print(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, models.URL{Key: key, LongUrl: longUrl, ShortUrl: shortUrl})
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GetURLByKey(c *gin.Context) {
	key := c.Param("key")
	var result models.URL

	row := database.DB.QueryRow("SELECT key, long_url, short_url FROM url_model WHERE key = ?", key)
	err := row.Scan(&result.Key, &result.LongUrl, &result.ShortUrl)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
		return
	}

	// c.IndentedJSON(http.StatusOK, result)
	c.Redirect(http.StatusFound, result.LongUrl)
}

func DeleteURLByKey(c *gin.Context) {
	key := c.Param("key")

	_, err := database.DB.Exec("DELETE FROM url_model WHERE key = ?", key)
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
