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
		h.initMaterialsRoutes(v1)

		// accounts routes
		h.initCompanyRoutes(v1)
		h.initUserRoutes(v1)
		h.initSectionRoutes(v1)
	}
}

func parseOffsetQueryParam(c *gin.Context) (int64, error) {
	var offset int
	var err error

	offsetParam := c.Query("offset")
	if offsetParam != "" {
		offset, err = strconv.Atoi(offsetParam)
		if err != nil {
			return 0, domain.ErrInvalidOffsetParam
		}
	}

	return int64(offset), nil
}

func parseLimitQueryParam(c *gin.Context) (int64, error) {
	var limit = 100
	var err error

	limitParam := c.Query("limit")
	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			return 0, domain.ErrInvalidLimitParam
		}
	}

	return int64(limit), nil
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
