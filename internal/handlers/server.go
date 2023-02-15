package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedreamwork/ChatGPT-PROXY/internal/api"
)

func Proxy(c *gin.Context) {
	// Send request to OpenAI API and stream data to client
	api.Proxy(c)
}
