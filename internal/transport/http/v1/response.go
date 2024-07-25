package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/logger"
)

type idResponse struct {
	ID interface{} `json:"id"`
}

type response struct {
	ID                string   `json:"id,omitempty"`
	AllowRegistration bool     `json:"allow_registration,omitempty"`
	Message           string   `json:"message,omitempty"`
	Token             string   `json:"token,omitempty"`
	Sections          []string `json:"sections,omitempty"`
	Status            string   `json:"status,omitempty"`
	Name              string   `json:"name,omitempty,omitempty"`
	IsAdmin           bool     `json:"is_admin,omitempty"`
	CreatedUserID     string   `json:"created_user_id,omitempty"`
	Avatar            string   `json:"avatar,omitempty"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{Message: message})
}
