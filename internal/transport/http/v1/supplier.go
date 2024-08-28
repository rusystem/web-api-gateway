package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"net/http"
)

func (h *Handler) initSupplierRoutes(api *gin.RouterGroup) {
	spl := api.Group("/supplier", h.adminIdentity)
	{
		spl.GET("/:id", h.getSupplier)
		spl.POST("/", h.createSupplier)
	}
}

// @Summary Get supplier by id
// @Security ApiKeyAuth
// @Tags supplier
// @Description Получение поставщика по id
// @ID get-supplier
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} domain.Supplier
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /supplier/{id} [GET]
func (h *Handler) getSupplier(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	spl, err := h.services.Supplier.GetById(c, int64(id))
	if err != nil {
		if errors.Is(err, domain.ErrSupplierNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, spl)
}

// @Summary Create supplier
// @Security ApiKeyAuth
// @Tags supplier
// @Description Создание поставщика
// @ID create-supplier
// @Accept json
// @Produce json
// @Param input body domain.InputSupplier true "Необходимо указать данные поставщика."
// @Success 200 {object} domain.IdResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /supplier [POST]
func (h *Handler) createSupplier(c *gin.Context) {
	var inp domain.Supplier
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	id, err := h.services.Supplier.Create(c, inp)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.IdResponse{ID: id})
}
