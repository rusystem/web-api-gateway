package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"net/http"
)

func (h *Handler) initWarehouseRoutes(api *gin.RouterGroup) {
	//supplier := api.Group("/warehouse", h.userIdentity) //todo вернуть после тестирования аутентификации
	spl := api.Group("/warehouse")
	{
		spl.GET("/:id", h.getWarehouse)
		spl.POST("/", h.createWarehouse)
	}
}

// todo добавить второй строкой @Security ApiKeyAuth

// @Summary Get warehouse by id
// @Tags warehouse
// @Description Получение склада по id
// @ID get-warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /warehouse/{id} [GET]
func (h *Handler) getWarehouse(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wh, err := h.services.Warehouse.GetById(c, int64(id))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wh)
}

// todo добавить второй строкой @Security ApiKeyAuth

// @Summary Create warehouse
// @Tags warehouse
// @Description Создание склада
// @ID create-warehouse
// @Accept json
// @Produce json
// @Param input body domain.Warehouse true "Необходимо указать данные склада."
// @Success 200 {object} domain.IdResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /supplier [POST]
func (h *Handler) createWarehouse(c *gin.Context) {
	var inp domain.Warehouse
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	id, err := h.services.Warehouse.Create(c, inp)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.IdResponse{ID: id})
}
