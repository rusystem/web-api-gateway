package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"net/http"
)

func (h *Handler) initWarehouseRoutes(api *gin.RouterGroup) {
	wh := api.Group("/warehouse")
	{
		wh.GET("/:id", h.userIdentity, h.getWarehouse)

		// only admin can create, update, delete warehouse
		wh.POST("/", h.adminIdentity, h.createWarehouse)
	}
}

// @Summary Get warehouse by id
// @Security ApiKeyAuth
// @Tags warehouse
// @Description Получение склада по id
// @ID get-warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} response
// @Router /warehouse/{id} [GET]
func (h *Handler) getWarehouse(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wh, err := h.services.Warehouse.GetById(c, id)
	if err != nil {
		if errors.Is(err, domain.ErrWarehouseNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wh)
}

// @Summary Create warehouse
// @Security ApiKeyAuth
// @Tags warehouse
// @Description Создание склада
// @ID create-warehouse
// @Accept json
// @Produce json
// @Param input body domain.InputWarehouse true "Необходимо указать данные склада."
// @Success 200 {object} domain.IdResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /warehouse [POST]
func (h *Handler) createWarehouse(c *gin.Context) {
	var inp domain.Warehouse
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	id, err := h.services.Warehouse.Create(c, inp) //todo добавить привязку склада к компании
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.IdResponse{ID: id})
}
