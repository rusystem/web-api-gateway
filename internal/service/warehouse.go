package service

import (
	"context"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/warehouse"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type Warehouse interface {
	GetById(ctx context.Context, id int64) (domain.Warehouse, error)
	Create(ctx context.Context, wh domain.Warehouse) (int64, error)
	Update(ctx context.Context, wh domain.Warehouse) error
	Delete(ctx context.Context, id int64) error
	GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.Warehouse, error)
}

type WarehouseServices struct {
	cfg *config.Config
	wc  *warehouse.WarehouseClient
}

func NewWarehouseServices(cfg *config.Config, warehouseClient *warehouse.WarehouseClient) *WarehouseServices {
	return &WarehouseServices{
		cfg: cfg,
		wc:  warehouseClient,
	}
}

func (s *WarehouseServices) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	return s.wc.GetById(ctx, id)
}

func (s *WarehouseServices) Create(ctx context.Context, wh domain.Warehouse) (int64, error) {
	return s.wc.Create(ctx, wh)
}

func (s *WarehouseServices) Update(ctx context.Context, wh domain.Warehouse) error {
	return s.wc.Update(ctx, wh)
}

func (s *WarehouseServices) Delete(ctx context.Context, id int64) error {
	return s.wc.Delete(ctx, id)
}

func (s *WarehouseServices) GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.Warehouse, error) {
	return s.wc.GetList(ctx, companyId)
}
