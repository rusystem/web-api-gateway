package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/service"
	"github.com/rusystem/web-api-gateway/pkg/auth"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"strconv"
)

type Handler struct {
	services     *service.Service
	tokenManager auth.TokenManager
	cfg          *config.Config
}

func NewHandler(services *service.Service, tokenManager auth.TokenManager, cfg *config.Config) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		cfg:          cfg,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)

		// warehouse routes
		h.initSupplierRoutes(v1)
		h.initWarehouseRoutes(v1)

		// accounts routes
		h.initCompanyRoutes(v1)
		h.initUserRoutes(v1)
		h.initSectionRoutes(v1)
	}
}

func parseSkipQueryParam(c *gin.Context) (int, error) {
	var skip int
	var err error

	skipParam := c.Query("skip")
	if skipParam != "" {
		skip, err = strconv.Atoi(skipParam)
		if err != nil {
			return 0, domain.ErrInvalidSkipParam
		}
	}

	return skip, nil
}

func parseTakeQueryParam(c *gin.Context) (int, error) {
	var take = 100
	var err error

	takeParam := c.Query("take")
	if takeParam != "" {
		take, err = strconv.Atoi(takeParam)
		if err != nil {
			return 0, domain.ErrInvalidTakeParam
		}
	}

	return take, nil
}

func parseIdIntPathParam(c *gin.Context) (int64, error) {
	idParam := c.Param("id")
	if idParam == "" {
		return 0, domain.ErrInvalidIdParam
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, domain.ErrInvalidIdParam
	}

	return int64(id), nil
}

func parseIdStringPathParam(c *gin.Context) (string, error) {
	id := c.Param("id")
	if id == "" {
		return "", domain.ErrInvalidIdParam
	}

	return id, nil
}

func parseEmailPathParam(c *gin.Context) (string, error) {
	id := c.Param("email")
	if id == "" {
		return "", domain.ErrInvalidEmailParam
	}

	return id, nil
}
