package service

import (
	"context"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/warehouse"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	tools "github.com/rusystem/web-api-gateway/tools"
	"time"
)

type Materials interface {
	CreatePlanning(ctx context.Context, material domain.Material) (int64, error)
	GetPlanningById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error)
	UpdatePlanningById(ctx context.Context, inp domain.UpdatePlanningMaterial, info domain.JWTInfo) error
	DeletePlanningById(ctx context.Context, id int64, info domain.JWTInfo) error
	GetPlanningList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	MovePlanningToPurchased(ctx context.Context, id int64, info domain.JWTInfo) (int64, int64, error)

	CreatePurchased(ctx context.Context, material domain.Material) (int64, int64, error)
	GetPurchasedById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error)
	UpdatePurchasedById(ctx context.Context, inp domain.UpdatePurchasedMaterial, info domain.JWTInfo) error
	DeletePurchasedById(ctx context.Context, id int64, info domain.JWTInfo) error
	GetPurchasedList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	MovePurchasedToArchive(ctx context.Context, id int64, info domain.JWTInfo) error

	GetPlanningArchiveById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error)
	GetPurchasedArchiveById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error)
	GetPlanningArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	GetPurchasedArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	DeletePlanningArchiveById(ctx context.Context, id int64, info domain.JWTInfo) error
	DeletePurchasedArchiveById(ctx context.Context, id int64, info domain.JWTInfo) error
}

type MaterialsService struct {
	cfg            *config.Config
	materialClient *warehouse.MaterialsClient
}

func NewMaterialsService(cfg *config.Config, materialClient *warehouse.MaterialsClient) *MaterialsService {
	return &MaterialsService{
		cfg:            cfg,
		materialClient: materialClient,
	}
}

func (s *MaterialsService) CreatePlanning(ctx context.Context, material domain.Material) (int64, error) {
	return s.materialClient.CreatePlanning(ctx, material)
}

func (s *MaterialsService) UpdatePlanningById(ctx context.Context, inp domain.UpdatePlanningMaterial, info domain.JWTInfo) error {
	material, err := s.materialClient.GetPlanningById(ctx, inp.ID)
	if err != nil {
		return err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	if inp.WarehouseID != nil {
		material.WarehouseID = *inp.WarehouseID
	}

	if inp.Name != nil {
		material.Name = *inp.Name
	}

	if inp.ByInvoice != nil {
		material.ByInvoice = *inp.ByInvoice
	}

	if inp.Article != nil {
		material.Article = *inp.Article
	}

	if inp.ProductCategory != nil {
		material.ProductCategory = *inp.ProductCategory
	}

	if inp.Unit != nil {
		material.Unit = *inp.Unit
	}

	if inp.TotalQuantity != nil {
		material.TotalQuantity = *inp.TotalQuantity
	}

	if inp.Volume != nil {
		material.Volume = *inp.Volume
	}

	if inp.PriceWithoutVAT != nil {
		material.PriceWithoutVAT = *inp.PriceWithoutVAT
	}

	if inp.TotalWithoutVAT != nil {
		material.TotalWithoutVAT = *inp.TotalWithoutVAT
	}

	if inp.SupplierID != nil {
		material.SupplierID = *inp.SupplierID
	}

	if inp.Location != nil {
		material.Location = *inp.Location
	}

	if inp.Contract != nil {
		material.Contract = *inp.Contract
	}

	if inp.File != nil {
		material.File = *inp.File
	}

	if inp.Status != nil {
		material.Status = *inp.Status
	}

	if inp.Comments != nil {
		material.Comments = *inp.Comments
	}

	if inp.Reserve != nil {
		material.Reserve = *inp.Reserve
	}

	if inp.ReceivedDate != nil {
		material.ReceivedDate = *inp.ReceivedDate
	}

	material.LastUpdated = time.Now().UTC()

	if inp.MinStockLevel != nil {
		material.MinStockLevel = *inp.MinStockLevel
	}

	if inp.ExpirationDate != nil {
		material.ExpirationDate = *inp.ExpirationDate
	}

	if inp.ResponsiblePerson != nil {
		material.ResponsiblePerson = *inp.ResponsiblePerson
	}

	if inp.StorageCost != nil {
		material.StorageCost = *inp.StorageCost
	}

	if inp.WarehouseSection != nil {
		material.WarehouseSection = *inp.WarehouseSection
	}

	if inp.IncomingDeliveryNumber != nil {
		material.IncomingDeliveryNumber = *inp.IncomingDeliveryNumber
	}

	if inp.OtherFields != nil {
		material.OtherFields = *inp.OtherFields
	}

	return s.materialClient.UpdatePlanningById(ctx, material)
}

func (s *MaterialsService) DeletePlanningById(ctx context.Context, id int64, info domain.JWTInfo) error {
	material, err := s.materialClient.GetPlanningById(ctx, id)
	if err != nil {
		return err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	return s.materialClient.DeletePlanningById(ctx, id)
}

func (s *MaterialsService) GetPlanningById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error) {
	material, err := s.materialClient.GetPlanningById(ctx, id)
	if err != nil {
		return domain.Material{}, err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.Material{}, domain.ErrNotAllowed
	}

	return material, nil
}

func (s *MaterialsService) GetPlanningList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return s.materialClient.GetListPlanning(ctx, params)
}

func (s *MaterialsService) MovePlanningToPurchased(ctx context.Context, id int64, info domain.JWTInfo) (int64, int64, error) {
	material, err := s.materialClient.GetPlanningById(ctx, id)
	if err != nil {
		return 0, 0, err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return 0, 0, domain.ErrNotAllowed
	}

	return s.materialClient.MovePlanningToPurchased(ctx, id)
}

func (s *MaterialsService) CreatePurchased(ctx context.Context, material domain.Material) (int64, int64, error) {
	return s.materialClient.CreatePurchased(ctx, material)
}

func (s *MaterialsService) UpdatePurchasedById(ctx context.Context, inp domain.UpdatePurchasedMaterial, info domain.JWTInfo) error {
	material, err := s.materialClient.GetPurchasedById(ctx, inp.ID)
	if err != nil {
		return err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	if inp.WarehouseID != nil {
		material.WarehouseID = *inp.WarehouseID
	}

	if inp.Name != nil {
		material.Name = *inp.Name
	}

	if inp.ByInvoice != nil {
		material.ByInvoice = *inp.ByInvoice
	}

	if inp.Article != nil {
		material.Article = *inp.Article
	}

	if inp.ProductCategory != nil {
		material.ProductCategory = *inp.ProductCategory
	}

	if inp.Unit != nil {
		material.Unit = *inp.Unit
	}

	if inp.TotalQuantity != nil {
		material.TotalQuantity = *inp.TotalQuantity
	}

	if inp.Volume != nil {
		material.Volume = *inp.Volume
	}

	if inp.PriceWithoutVAT != nil {
		material.PriceWithoutVAT = *inp.PriceWithoutVAT
	}

	if inp.TotalWithoutVAT != nil {
		material.TotalWithoutVAT = *inp.TotalWithoutVAT
	}

	if inp.SupplierID != nil {
		material.SupplierID = *inp.SupplierID
	}

	if inp.Location != nil {
		material.Location = *inp.Location
	}

	if inp.Contract != nil {
		material.Contract = *inp.Contract
	}

	if inp.File != nil {
		material.File = *inp.File
	}

	if inp.Status != nil {
		material.Status = *inp.Status
	}

	if inp.Comments != nil {
		material.Comments = *inp.Comments
	}

	if inp.Reserve != nil {
		material.Reserve = *inp.Reserve
	}

	if inp.ReceivedDate != nil {
		material.ReceivedDate = *inp.ReceivedDate
	}

	material.LastUpdated = time.Now().UTC()

	if inp.MinStockLevel != nil {
		material.MinStockLevel = *inp.MinStockLevel
	}

	if inp.ExpirationDate != nil {
		material.ExpirationDate = *inp.ExpirationDate
	}

	if inp.ResponsiblePerson != nil {
		material.ResponsiblePerson = *inp.ResponsiblePerson
	}

	if inp.StorageCost != nil {
		material.StorageCost = *inp.StorageCost
	}

	if inp.WarehouseSection != nil {
		material.WarehouseSection = *inp.WarehouseSection
	}

	if inp.IncomingDeliveryNumber != nil {
		material.IncomingDeliveryNumber = *inp.IncomingDeliveryNumber
	}

	if inp.OtherFields != nil {
		material.OtherFields = *inp.OtherFields
	}

	return s.materialClient.UpdatePurchasedById(ctx, material)
}

func (s *MaterialsService) DeletePurchasedById(ctx context.Context, id int64, info domain.JWTInfo) error {
	material, err := s.materialClient.GetPurchasedById(ctx, id)
	if err != nil {
		return err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	return s.materialClient.DeletePurchasedById(ctx, id)
}

func (s *MaterialsService) GetPurchasedById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error) {
	material, err := s.materialClient.GetPurchasedById(ctx, id)
	if err != nil {
		return domain.Material{}, err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.Material{}, domain.ErrNotAllowed
	}

	return material, nil
}

func (s *MaterialsService) GetPurchasedList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return s.materialClient.GetListPurchased(ctx, params)
}

func (s *MaterialsService) MovePurchasedToArchive(ctx context.Context, id int64, info domain.JWTInfo) error {
	material, err := s.materialClient.GetPurchasedById(ctx, id)
	if err != nil {
		return err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	return s.materialClient.MovePurchasedToArchive(ctx, id)
}

func (s *MaterialsService) GetPlanningArchiveById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error) {
	material, err := s.materialClient.GetPlanningArchiveById(ctx, id)
	if err != nil {
		return domain.Material{}, err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.Material{}, domain.ErrNotAllowed
	}

	return material, nil
}

func (s *MaterialsService) GetPurchasedArchiveById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Material, error) {
	material, err := s.materialClient.GetPurchasedArchiveById(ctx, id)
	if err != nil {
		return domain.Material{}, err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.Material{}, domain.ErrNotAllowed
	}

	return material, nil
}

func (s *MaterialsService) GetPlanningArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return s.materialClient.GetListPlanningArchive(ctx, params)
}

func (s *MaterialsService) GetPurchasedArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return s.materialClient.GetListPurchasedArchive(ctx, params)
}

func (s *MaterialsService) DeletePlanningArchiveById(ctx context.Context, id int64, info domain.JWTInfo) error {
	material, err := s.materialClient.GetPlanningArchiveById(ctx, id)
	if err != nil {
		return err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	return s.materialClient.DeletePlanningArchiveById(ctx, id)
}

func (s *MaterialsService) DeletePurchasedArchiveById(ctx context.Context, id int64, info domain.JWTInfo) error {
	material, err := s.materialClient.GetPurchasedArchiveById(ctx, id)
	if err != nil {
		return err
	}

	if material.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	return s.materialClient.DeletePurchasedArchiveById(ctx, id)
}
