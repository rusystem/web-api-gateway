package service

import (
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository"
	"github.com/rusystem/web-api-gateway/pkg/auth"
	grpc "github.com/rusystem/web-api-gateway/pkg/client/grpc/accounts"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/warehouse"
)

type Config struct {
	Config          *config.Config
	Repo            *repository.Repository
	TokenManager    auth.TokenManager
	SuppliersClient *warehouse.SuppliersClient
	WarehouseClient *warehouse.WarehouseClient
	UserClient      *grpc.UserAccountsClient
	CompanyClient   *grpc.CompanyAccountsClient
	SectionsClient  *grpc.SectionsAccountsClient
}

type Service struct {
	Auth      Auth
	Supplier  Supplier
	Warehouse Warehouse
	User      User
	Company   Company
	Sections  Sections
}

func New(cfg Config) *Service {
	return &Service{
		Auth:      NewAuthServices(cfg.Config, cfg.Repo, cfg.TokenManager, cfg.UserClient),
		Supplier:  NewSupplierService(cfg.Config, cfg.SuppliersClient),
		Warehouse: NewWarehouseServices(cfg.Config, cfg.WarehouseClient),
		User:      NewUserServices(cfg.Config, cfg.UserClient),
		Company:   NewCompanyService(cfg.Config, cfg.CompanyClient),
		Sections:  NewSectionsService(cfg.Config, cfg.SectionsClient),
	}
}
