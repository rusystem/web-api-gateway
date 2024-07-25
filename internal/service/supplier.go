package service

import (
	"context"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type Supplier interface {
	GetById(ctx context.Context, id int64) (domain.Supplier, error)
	Create(ctx context.Context, spl domain.Supplier) (int64, error)
}

type SupplierService struct {
	cfg            *config.Config
	supplierClient *grpc.SuppliersClient
}

func NewSupplierService(cfg *config.Config, supplierClient *grpc.SuppliersClient) *SupplierService {
	return &SupplierService{
		cfg:            cfg,
		supplierClient: supplierClient,
	}
}

func (s *SupplierService) GetById(ctx context.Context, id int64) (domain.Supplier, error) {
	return s.supplierClient.GetById(ctx, id)
}

func (s *SupplierService) Create(ctx context.Context, spl domain.Supplier) (int64, error) {
	return s.supplierClient.Create(ctx, spl)
}
