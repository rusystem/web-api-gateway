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
		wh.GET("/", h.userIdentity, h.getWarehouses)

		// only admin can create, update, delete warehouse
		wh.POST("/", h.adminIdentity, h.createWarehouse)
		wh.PUT("/:id", h.adminIdentity, h.updateWarehouse)
		wh.DELETE("/:id", h.adminIdentity, h.deleteWarehouse)
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

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	wh, err := h.services.Warehouse.GetById(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrNotAllowed) {
			newResponse(c, http.StatusForbidden, err.Error())
			return
		}

		if errors.Is(err, domain.ErrWarehouseNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wh)
}

// @Summary Get warehouses
// @Security ApiKeyAuth
// @Tags warehouse
// @Description Получение списка складов
// @ID get-warehouses
// @Accept json
// @Produce json
// @Success 200 {array} domain.Warehouse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} response
// @Router /warehouse [GET]
func (h *Handler) getWarehouses(c *gin.Context) {
	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	whs, err := h.services.Warehouse.GetListByCompanyId(c, info.CompanyId)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, whs)
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
	var inp domain.InputWarehouse
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Warehouse.Create(c, domain.Warehouse{
		Name:              inp.Name,
		Address:           inp.Address,
		ResponsiblePerson: inp.ResponsiblePerson,
		Phone:             inp.Phone,
		Email:             inp.Email,
		MaxCapacity:       inp.MaxCapacity,
		CurrentOccupancy:  inp.CurrentOccupancy,
		OtherFields:       inp.OtherFields,
		Country:           inp.Country,
		CompanyId:         info.CompanyId,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.IdResponse{ID: id})
}

// @Summary Update warehouse
// @Security ApiKeyAuth
// @Tags warehouse
// @Description Обновление склада своей компании
// @ID update-warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Param input body domain.InputWarehouse true "Необходимо указать данные склада."
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /warehouse/{id} [PUT]
func (h *Handler) updateWarehouse(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.WarehouseUpdate
	if err := c.ShouldBindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	inp.ID = id

	if err := h.services.Warehouse.Update(c, inp, info); err != nil {
		if errors.Is(err, domain.ErrWarehouseNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if errors.Is(err, domain.ErrNotAllowed) {
			newResponse(c, http.StatusForbidden, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete warehouse
// @Security ApiKeyAuth
// @Tags warehouse
// @Description Удаление склада своей компании
// @ID delete-warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /warehouse/{id} [DELETE]
func (h *Handler) deleteWarehouse(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.services.Warehouse.Delete(c, id, info); err != nil {
		if errors.Is(err, domain.ErrWarehouseNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if errors.Is(err, domain.ErrNotAllowed) {
			newResponse(c, http.StatusForbidden, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
