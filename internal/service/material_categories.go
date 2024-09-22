package service

import (
	"context"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/warehouse"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"time"
)

type Category interface {
	Create(ctx context.Context, category domain.MaterialCategory) (int64, error)
	GetById(ctx context.Context, id, companyId int64) (domain.MaterialCategory, error)
	Update(ctx context.Context, inp domain.UpdateMaterialCategory) error
	Delete(ctx context.Context, id, companyId int64) error
	List(ctx context.Context, param domain.MaterialParams) ([]domain.MaterialCategory, error)
	Search(ctx context.Context, param domain.MaterialParams) ([]domain.MaterialCategory, error)
}

type MaterialCategoriesService struct {
	cfg            *config.Config
	materialClient *warehouse.MaterialsClient
}

func NewMaterialCategoriesService(cfg *config.Config, materialClient *warehouse.MaterialsClient) *MaterialCategoriesService {
	return &MaterialCategoriesService{
		cfg:            cfg,
		materialClient: materialClient,
	}
}

func (s *MaterialCategoriesService) Create(ctx context.Context, category domain.MaterialCategory) (int64, error) {
	return s.materialClient.CreateMaterialCategory(ctx, category)
}

func (s *MaterialCategoriesService) GetById(ctx context.Context, id, companyId int64) (domain.MaterialCategory, error) {
	return s.materialClient.GetByIdMaterialCategory(ctx, id, companyId)
}

func (s *MaterialCategoriesService) Update(ctx context.Context, inp domain.UpdateMaterialCategory) error {
	category, err := s.materialClient.GetByIdMaterialCategory(ctx, inp.ID, inp.CompanyID)
	if err != nil {
		return err
	}

	if inp.Name != nil {
		category.Name = *inp.Name
	}

	if inp.Description != nil {
		category.Description = *inp.Description
	}

	if inp.Slug != nil {
		category.Slug = *inp.Slug
	}

	if inp.IsActive != nil {
		category.IsActive = *inp.IsActive
	}

	if inp.ImgURL != nil {
		category.ImgURL = *inp.ImgURL
	}

	category.UpdatedAt = time.Now().UTC()

	return s.materialClient.UpdateMaterialCategory(ctx, category)
}

func (s *MaterialCategoriesService) Delete(ctx context.Context, id, companyId int64) error {
	return s.materialClient.DeleteMaterialCategory(ctx, id, companyId)
}

func (s *MaterialCategoriesService) List(ctx context.Context, param domain.MaterialParams) ([]domain.MaterialCategory, error) {
	return s.materialClient.GetListMaterialCategory(ctx, param)
}

func (s *MaterialCategoriesService) Search(ctx context.Context, param domain.MaterialParams) ([]domain.MaterialCategory, error) {
	return s.materialClient.SearchMaterialCategory(ctx, param)
}
