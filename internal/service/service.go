package service

import (
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository"
	"github.com/rusystem/web-api-gateway/pkg/auth"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc"
)

type Service struct {
	Auth      Auth
	Supplier  Supplier
	Warehouse Warehouse
}

func New(cfg *config.Config, repo *repository.Repository, tokenManager auth.TokenManager, sc *grpc.SuppliersClient, wc *grpc.WarehouseClient) *Service {
	return &Service{
		Auth:      NewAuthServices(cfg, repo, tokenManager),
		Supplier:  NewSupplierService(cfg, sc),
		Warehouse: NewWarehouseServices(cfg, wc),
	}
}
