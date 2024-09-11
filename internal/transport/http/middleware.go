package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	allowedOrigins := []string{"http://localhost", "http://127.0.0.1", "http://91.243.71.100"}

	// Проверяем, разрешён ли запрос с данного origin
	for _, o := range allowedOrigins {
		if origin == o {
			c.Header("Access-Control-Allow-Origin", origin)
			break
		}
	}

	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Content-Type", "application/json")

	// Для OPTIONS-запросов возвращаем статус OK
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}

/*func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}*/
