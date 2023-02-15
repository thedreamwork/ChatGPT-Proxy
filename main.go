package main

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thedreamwork/ChatGPT-PROXY/internal/handlers"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := gin.Default()

	// Proxy all requests to /* to proxy if not already handled
	handler.Any("/*path", handlers.Proxy)

	log.Println("listening on", PORT)
	log.Fatal(endless.ListenAndServe(":"+PORT, handler))
}
