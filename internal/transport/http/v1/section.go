package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"net/http"
)

func (h *Handler) initSectionRoutes(api *gin.RouterGroup) {
	sections := api.Group("/sections")
	{
		sections.GET("/:id", h.adminIdentity, h.getSection)
		sections.GET("/", h.adminIdentity, h.getSections)

		// only super admin can create, update, delete section
		sections.POST("/", h.superAdminIdentity, h.createSection)
		sections.PUT("/:id", h.superAdminIdentity, h.updateSection)
		sections.DELETE("/:id", h.superAdminIdentity, h.deleteSection)
	}
}

// @Summary Get section
// @Security ApiKeyAuth
// @Tags sections
// @Description Получение секции по id
// @ID get-section
// @Accept  json
// @Produce  json
// @Param id path int true "Section ID" example(1)
// @Success 200 {object} domain.Section
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /sections/{id} [GET]
func (h *Handler) getSection(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	section, err := h.services.Sections.GetById(c, id)
	if err != nil {
		if errors.Is(err, domain.ErrSectionNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, section)
}

// @Summary Get sections
// @Security ApiKeyAuth
// @Tags sections
// @Description Получение списка секций
// @ID get-sections
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Section
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /sections [GET]
func (h *Handler) getSections(c *gin.Context) {
	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	sections, err := h.services.Sections.List(c, info)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sections)
}

// @Summary Create section
// @Security ApiKeyAuth
// @Tags sections
// @Description Создание секции
// @Description Только super admin может создавать секции
// @ID create-section
// @Accept  json
// @Produce  json
// @Param section body domain.SectionCreate true "Section"
// @Success 200 {int} int
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /sections [POST]
func (h *Handler) createSection(c *gin.Context) {
	var section domain.SectionCreate
	if err := c.BindJSON(&section); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sectionId, err := h.services.Sections.Create(c, domain.Section{
		Name: section.Name,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sectionId)
}

// @Summary Update section
// @Security ApiKeyAuth
// @Tags sections
// @Description Обновление секции
// @ID update-section
// @Accept  json
// @Produce  json
// @Param id path int true "Section ID" example(1)
// @Param section body domain.SectionUpdate true "Section"
// @Success 200
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /sections/{id} [PUT]
func (h *Handler) updateSection(c *gin.Context) {
	var req domain.SectionUpdate
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Sections.Update(c, domain.Section{
		Id:   id,
		Name: req.Name,
	}); err != nil {
		if errors.Is(err, domain.ErrSectionNotFound) {
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

// @Summary Delete section
// @Security ApiKeyAuth
// @Tags sections
// @Description Удаление секции
// @ID delete-section
// @Accept  json
// @Produce  json
// @Param id path int true "Section ID" example(1)
// @Success 200
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /sections/{id} [DELETE]
func (h *Handler) deleteSection(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Sections.Delete(c, id); err != nil {
		if errors.Is(err, domain.ErrSectionNotFound) {
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
