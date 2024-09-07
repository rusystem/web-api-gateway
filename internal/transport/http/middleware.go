package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if origin != "http://localhost" && origin != "http://127.0.0.1" {
		// Устанавливаем CORS-заголовки, если запрос не с localhost
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
	}

	c.Header("Content-Type", "application/json")

	// Для OPTIONS-запросов возвращаем статус OK
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
