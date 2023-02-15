package api

import (
	_ "embed"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Proxy(c *gin.Context) {
	url := os.Getenv("OPENAI_URL") + c.Param("path")
	// POST request with all data and headers
	var req *http.Request
	var err error
	if c.Request.Method == "POST" {
		req, err = http.NewRequest("POST", url, c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error", "error": err})
			return
		}
	} else if c.Request.Method == "GET" {
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error", "error": err})
			return
		}
	} else if c.Request.Method == "PATCH" {
		req, err = http.NewRequest("PATCH", url, c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error", "error": err})
			return
		}
	} else if c.Request.Method == "DELETE" {
		req, err = http.NewRequest("DELETE", url, c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error", "error": err})
			return
		}
	} else {
		c.JSON(500, gin.H{"message": "Internal server error", "error": "Invalid HTTP method"})
		return
	}
	// Add headers
	for key, value := range c.Request.Header {
		req.Header.Set(key, value[0])
	}
	// Add content type JSON
	req.Header.Set("Content-Type", "application/json")
	// Set keep alive and timeout
	req.Close = false
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Keep-Alive", "timeout=360")
	// Send request
	client := &http.Client{Timeout: time.Second * 360}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error", "error": err})
		return
	}
	// Stream response to client
	defer resp.Body.Close()
	// Return stream of data to client
	c.Stream(func(w io.Writer) bool {
		// Write data to client
		io.Copy(w, resp.Body)
		return false
	})
}
