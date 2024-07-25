package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"net/http"
)

func (h *Handler) initSupplierRoutes(api *gin.RouterGroup) {
	//supplier := api.Group("/supplier", h.userIdentity) //todo вернуть после тестирования аутентификации
	spl := api.Group("/supplier")
	{
		spl.GET("/:id", h.getSupplier)
		spl.POST("/", h.createSupplier)
	}
}

// todo добавить второй строкой @Security ApiKeyAuth

// @Summary Get supplier by id
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
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, spl)
}

// todo добавить второй строкой @Security ApiKeyAuth

// @Summary Create supplier
// @Tags supplier
// @Description Создание поставщика
// @ID create-supplier
// @Accept json
// @Produce json
// @Param input body domain.Supplier true "Необходимо указать данные поставщика."
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
