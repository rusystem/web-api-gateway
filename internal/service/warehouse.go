package service

import (
	"context"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type Warehouse interface {
	GetById(ctx context.Context, id int64) (domain.Warehouse, error)
	Create(ctx context.Context, wh domain.Warehouse) (int64, error)
}

type WarehouseServices struct {
	cfg *config.Config
	wc  *grpc.WarehouseClient
}

func NewWarehouseServices(cfg *config.Config, wc *grpc.WarehouseClient) *WarehouseServices {
	return &WarehouseServices{
		cfg: cfg,
		wc:  wc,
	}
}

func (s *WarehouseServices) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	return s.wc.GetById(ctx, id)
}

func (s *WarehouseServices) Create(ctx context.Context, wh domain.Warehouse) (int64, error) {
	return s.wc.Create(ctx, wh)
}
