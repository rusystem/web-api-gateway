package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"github.com/rusystem/web-api-gateway/tools"
	"net/http"
	"time"
)

func (h *Handler) initMaterialsRoutes(api *gin.RouterGroup) {
	materials := api.Group("/materials", h.userIdentity) //todo добавить секции в мидлвейер
	{
		planning := materials.Group("/planning")
		{
			planning.POST("/", h.createPlanning)
			planning.GET("/:id", h.getPlanningById)
			planning.PUT("/:id", h.updatePlanningById)
			planning.DELETE("/:id", h.deletePlanningById)
			planning.GET("/list", h.getPlanningList)
			planning.PUT("/move-to-purchased/:id", h.movePlanningToPurchased)
		}

		purchased := materials.Group("/purchased")
		{
			purchased.POST("/", h.createPurchased)
			purchased.GET("/:id", h.getPurchasedById)
			purchased.PUT("/:id", h.updatePurchasedById)
			purchased.DELETE("/:id", h.deletePurchasedById)
			purchased.GET("/list", h.getPurchasedList)
			purchased.GET("/:id/qr-code", h.getPurchasedQrCode)
			purchased.GET("/:id/barcode", h.getPurchasedBarcode)
			purchased.PUT("/move-to-archive/:id", h.movePurchasedToArchive)
		}

		archive := materials.Group("/archive")
		{
			planning := archive.Group("/planning")
			{
				planning.GET("/:id", h.getPlanningArchiveById)
				planning.GET("/list", h.getPlanningArchiveList)
				planning.DELETE("/:id", h.deletePlanningArchiveById)
			}

			purchased := archive.Group("/purchased")
			{
				purchased.GET("/:id", h.getPurchasedArchiveById)
				purchased.GET("/list", h.getPurchasedArchiveList)
				purchased.DELETE("/:id", h.deletePurchasedArchiveById)
			}
		}
	}
}

// @Summary Create planning material
// @Security ApiKeyAuth
// @Tags materials planning
// @Description Создание планируемого материала
// @ID create-planning-material
// @Accept json
// @Produce json
// @Param input body domain.CreatePlanningMaterial true "Необходимо указать данные планируемого материала"
// @Success 200 {object} domain.IdResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/planning [POST]
func (h *Handler) createPlanning(c *gin.Context) {
	var inp domain.CreatePlanningMaterial
	if err := c.ShouldBindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Materials.CreatePlanning(c, domain.Material{
		WarehouseID:            inp.WarehouseID,
		Name:                   inp.Name,
		ByInvoice:              inp.ByInvoice,
		Article:                inp.Article,
		ProductCategory:        inp.ProductCategory,
		Unit:                   inp.Unit,
		TotalQuantity:          inp.TotalQuantity,
		Volume:                 inp.Volume,
		PriceWithoutVAT:        inp.PriceWithoutVAT,
		TotalWithoutVAT:        inp.TotalWithoutVAT,
		SupplierID:             inp.SupplierID,
		Contract:               inp.Contract,
		File:                   inp.File,
		Status:                 inp.Status,
		Comments:               inp.Comments,
		Reserve:                inp.Reserve,
		ReceivedDate:           inp.ReceivedDate,
		LastUpdated:            time.Now().UTC(),
		MinStockLevel:          inp.MinStockLevel,
		ExpirationDate:         inp.ExpirationDate,
		ResponsiblePerson:      inp.ResponsiblePerson,
		StorageCost:            inp.StorageCost,
		WarehouseSection:       inp.WarehouseSection,
		IncomingDeliveryNumber: inp.IncomingDeliveryNumber,
		OtherFields:            inp.OtherFields,
		CompanyID:              info.CompanyId,
	})
	if err != nil {
		if errors.Is(err, domain.ErrWarehouseNotFound) || errors.Is(err, domain.ErrSupplierNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.IdResponse{ID: id})
}

// @Summary Get planning material
// @Security ApiKeyAuth
// @Tags materials planning
// @Description Получение планируемого материала
// @ID get-planning-material
// @Accept json
// @Produce json
// @Param id path int true "ID планируемого материала"
// @Success 200 {object} domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/planning/{id} [GET]
func (h *Handler) getPlanningById(c *gin.Context) {
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

	material, err := h.services.Materials.GetPlanningById(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

	c.JSON(http.StatusOK, material)
}

// @Summary Update planning material
// @Security ApiKeyAuth
// @Tags materials planning
// @Description Обновление планируемого материала
// @ID update-planning-material
// @Accept json
// @Produce json
// @Param id path int true "ID планируемого материала"
// @Param input body domain.UpdatePlanningMaterial true "Необходимо указать данные планируемого материала"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/planning/{id} [PUT]
func (h *Handler) updatePlanningById(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.UpdatePlanningMaterial
	if err := c.ShouldBindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	inp.ID = id

	if err = h.services.Materials.UpdatePlanningById(c, inp, info); err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if errors.Is(err, domain.ErrNotAllowed) {
			newResponse(c, http.StatusForbidden, err.Error())
			return
		}

		if errors.Is(err, domain.ErrWarehouseNotFound) || errors.Is(err, domain.ErrSupplierNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete planning material
// @Security ApiKeyAuth
// @Tags materials planning
// @Description Удаление планируемого материала
// @ID delete-planning-material
// @Accept json
// @Produce json
// @Param id path int true "ID планируемого материала"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/planning/{id} [DELETE]
func (h *Handler) deletePlanningById(c *gin.Context) {
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

	if err = h.services.Materials.DeletePlanningById(c, id, info); err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

// @Summary Get planning list
// @Security ApiKeyAuth
// @Tags materials planning
// @Description Список планируемых материалов
// @ID get-planning-list
// @Accept json
// @Produce json
// @Param limit query int false "limit query param"
// @Param offset query int false "offset query param"
// @Success 200 {object} []domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/planning/list [GET]
func (h *Handler) getPlanningList(c *gin.Context) {
	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	limit, err := parseLimitQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := parseOffsetQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mtrls, err := h.services.Materials.GetPlanningList(c.Request.Context(), domain.MaterialParams{
		Limit:     limit,
		Offset:    offset,
		CompanyId: info.CompanyId,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mtrls)
}

// @Summary Move planning material to purchased
// @Security ApiKeyAuth
// @Tags materials planning
// @Description Перемещение планируемого материала в закупленные
// @ID move-planning-to-purchased
// @Accept json
// @Produce json
// @Param id path int true "ID планируемого материала"
// @Success 200 {object} domain.PurchasedIdResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/planning/move-to-purchased/{id} [PUT]
func (h *Handler) movePlanningToPurchased(c *gin.Context) {
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

	newId, itemId, err := h.services.Materials.MovePlanningToPurchased(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

	c.JSON(http.StatusOK, domain.PurchasedIdResponse{ID: newId, ItemId: itemId})
}

// @Summary Create purchased material
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Создание закупленного материала
// @ID create-purchased-material
// @Accept json
// @Produce json
// @Param input body domain.CreatePurchasedMaterial true "Необходимо указать данные закупленного материала"
// @Success 200 {object} domain.PurchasedIdResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased [POST]
func (h *Handler) createPurchased(c *gin.Context) {
	var inp domain.CreatePurchasedMaterial
	if err := c.ShouldBindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, itemId, err := h.services.Materials.CreatePurchased(c, domain.Material{
		WarehouseID:            inp.WarehouseID,
		Name:                   inp.Name,
		ByInvoice:              inp.ByInvoice,
		Article:                inp.Article,
		ProductCategory:        inp.ProductCategory,
		Unit:                   inp.Unit,
		TotalQuantity:          inp.TotalQuantity,
		Volume:                 inp.Volume,
		PriceWithoutVAT:        inp.PriceWithoutVAT,
		TotalWithoutVAT:        inp.TotalWithoutVAT,
		SupplierID:             inp.SupplierID,
		Location:               inp.Location,
		Contract:               inp.Contract,
		File:                   inp.File,
		Status:                 inp.Status,
		Comments:               inp.Comments,
		Reserve:                inp.Reserve,
		ReceivedDate:           inp.ReceivedDate,
		LastUpdated:            time.Now().UTC(),
		MinStockLevel:          inp.MinStockLevel,
		ExpirationDate:         inp.ExpirationDate,
		ResponsiblePerson:      inp.ResponsiblePerson,
		StorageCost:            inp.StorageCost,
		WarehouseSection:       inp.WarehouseSection,
		IncomingDeliveryNumber: inp.IncomingDeliveryNumber,
		OtherFields:            inp.OtherFields,
		CompanyID:              info.CompanyId,
	})
	if err != nil {
		if errors.Is(err, domain.ErrWarehouseNotFound) || errors.Is(err, domain.ErrSupplierNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.PurchasedIdResponse{ID: id, ItemId: itemId})
}

// @Summary Get purchased material
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Получение закупленного материала
// @ID get-purchased-material
// @Accept json
// @Produce json
// @Param id path int true "ID закупленного материала"
// @Success 200 {object} domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased/{id} [GET]
func (h *Handler) getPurchasedById(c *gin.Context) {
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

	material, err := h.services.Materials.GetPurchasedById(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

	c.JSON(http.StatusOK, material)
}

// @Summary Update purchased material
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Обновление закупленного материала
// @ID update-purchased-material
// @Accept json
// @Produce json
// @Param id path int true "ID закупленного материала"
// @Param input body domain.UpdatePurchasedMaterial true "Необходимо указать данные закупленного материала"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased/{id} [PUT]
func (h *Handler) updatePurchasedById(c *gin.Context) {
	id, err := parseIdIntPathParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.UpdatePurchasedMaterial
	if err := c.ShouldBindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	inp.ID = id

	if err = h.services.Materials.UpdatePurchasedById(c, inp, info); err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if errors.Is(err, domain.ErrNotAllowed) {
			newResponse(c, http.StatusForbidden, err.Error())
			return
		}

		if errors.Is(err, domain.ErrWarehouseNotFound) || errors.Is(err, domain.ErrSupplierNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete purchased material
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Удаление закупленного материала
// @ID delete-purchased-material
// @Accept json
// @Produce json
// @Param id path int true "ID закупленного материала"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased/{id} [DELETE]
func (h *Handler) deletePurchasedById(c *gin.Context) {
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

	if err = h.services.Materials.DeletePurchasedById(c, id, info); err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

// @Summary Get purchased list
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Получение списка закупленных материалов
// @ID get-purchased-list
// @Accept json
// @Produce json
// @Param limit query int false "limit query param"
// @Param offset query int false "offset query param"
// @Success 200 {object} []domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased/list [GET]
func (h *Handler) getPurchasedList(c *gin.Context) {
	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	limit, err := parseLimitQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := parseOffsetQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mtrls, err := h.services.Materials.GetPurchasedList(c.Request.Context(), domain.MaterialParams{
		Limit:     limit,
		Offset:    offset,
		CompanyId: info.CompanyId,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mtrls)
}

// @Summary Get purchased QR code
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Получение QR кода закупленного материала
// @ID get-purchased-qr-code
// @Accept json
// @Produce  image/png
// @Param id path int true "ID закупленного материала"
// @Success 200 {file} png "QR-код в формате PNG"
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased/{id}/qr-code [GET]
func (h *Handler) getPurchasedQrCode(c *gin.Context) {
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

	material, err := h.services.Materials.GetPurchasedById(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

	qrCode, err := tools.GenerateQRCodePNG(domain.CodeInfo{
		Id:     material.ID,
		ItemId: material.ItemID,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", qrCode)
}

// @Summary Get purchased barcode
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Получение штрихкода закупленного материала
// @ID get-purchased-barcode
// @Accept json
// @Produce  image/png
// @Param id path int true "ID закупленного материала"
// @Success 200 {file} png "Штрихкод"
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased/{id}/barcode [GET]
func (h *Handler) getPurchasedBarcode(c *gin.Context) {
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

	material, err := h.services.Materials.GetPurchasedById(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

	width := 300
	height := 100 //todo возможно вынести в параметры

	barCode, err := tools.GenerateBarcode(domain.CodeInfo{
		Id:     material.ID,
		ItemId: material.ItemID,
	}, width, height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", barCode)
}

// @Summary Move purchased to archive
// @Security ApiKeyAuth
// @Tags materials purchased
// @Description Перемещение закупленного материала в архив
// @ID move-purchased-to-archive
// @Accept json
// @Produce json
// @Param id path int true "ID закупленного материала"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/purchased/move-to-archive/{id} [PUT]
func (h *Handler) movePurchasedToArchive(c *gin.Context) {
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

	if err = h.services.Materials.MovePurchasedToArchive(c, id, info); err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

// @Summary Get planning archive by id
// @Security ApiKeyAuth
// @Tags materials archive
// @Description Получение запланированного материала из архива по ID
// @ID get-planning-archive-by-id
// @Accept json
// @Produce json
// @Param id path int true "ID архиввного планируемого материала"
// @Success 200 {object} domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/archive/planning/{id} [GET]
func (h *Handler) getPlanningArchiveById(c *gin.Context) {
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

	material, err := h.services.Materials.GetPlanningArchiveById(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

	c.JSON(http.StatusOK, material)
}

// @Summary Get planning archive list
// @Security ApiKeyAuth
// @Tags materials archive
// @Description Получение списка запланированных материалов из архива
// @ID get-planning-archive-list
// @Accept json
// @Produce json
// @Success 200 {array} domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/archive/planning/list [GET]
func (h *Handler) getPlanningArchiveList(c *gin.Context) {
	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	limit, err := parseLimitQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := parseOffsetQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mtrls, err := h.services.Materials.GetPlanningArchiveList(c.Request.Context(), domain.MaterialParams{
		Limit:     limit,
		Offset:    offset,
		CompanyId: info.CompanyId,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mtrls)
}

// @Summary Delete planning archive by id
// @Security ApiKeyAuth
// @Tags materials archive
// @Description Удаление запланированного материала из архива по ID
// @ID delete-planning-archive-by-id
// @Accept json
// @Produce json
// @Param id path int true "ID архиввного планируемого материала"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/archive/planning/{id} [DELETE]
func (h *Handler) deletePlanningArchiveById(c *gin.Context) {
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

	if err = h.services.Materials.DeletePlanningArchiveById(c, id, info); err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

// @Summary Get purchased archive by id
// @Security ApiKeyAuth
// @Tags materials archive
// @Description Получение закупленного материала из архива по ID
// @ID get-purchased-archive-by-id
// @Accept json
// @Produce json
// @Param id path int true "ID архиввного закупленного материала"
// @Success 200 {object} domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/archive/purchased/{id} [GET]
func (h *Handler) getPurchasedArchiveById(c *gin.Context) {
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

	material, err := h.services.Materials.GetPurchasedArchiveById(c, id, info)
	if err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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

	c.JSON(http.StatusOK, material)
}

// @Summary Get purchased archive list
// @Security ApiKeyAuth
// @Tags materials archive
// @Description Получение списка закупленных материалов из архива
// @ID get-purchased-archive-list
// @Accept json
// @Produce json
// @Success 200 {array} domain.Material
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/archive/purchased/list [GET]
func (h *Handler) getPurchasedArchiveList(c *gin.Context) {
	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	limit, err := parseLimitQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := parseOffsetQueryParam(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mtrls, err := h.services.Materials.GetPurchasedArchiveList(c.Request.Context(), domain.MaterialParams{
		Limit:     limit,
		Offset:    offset,
		CompanyId: info.CompanyId,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mtrls)
}

// @Summary Delete purchased archive by id
// @Security ApiKeyAuth
// @Tags materials archive
// @Description Удаление закупленного материала из архива по ID
// @ID delete-purchased-archive-by-id
// @Accept json
// @Produce json
// @Param id path int true "ID архиввного закупленного материала"
// @Success 200
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /materials/archive/purchased/{id} [DELETE]
func (h *Handler) deletePurchasedArchiveById(c *gin.Context) {
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

	if err = h.services.Materials.DeletePurchasedArchiveById(c, id, info); err != nil {
		if errors.Is(err, domain.ErrMaterialNotFound) {
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
