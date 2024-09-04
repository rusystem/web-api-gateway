package repository

import (
	"context"
	"database/sql"
	"github.com/rusystem/cache"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository/database"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type Company interface {
	GetById(ctx context.Context, id int64) (domain.Company, error)
	IsExist(ctx context.Context, id int64) (bool, error)
}

type CompanyRepository struct {
	cfg   *config.Config
	cache *cache.MemoryCache
	db    database.Company
}

func NewCompanyRepository(cfg *config.Config, cache *cache.MemoryCache, db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		cfg:   cfg,
		cache: cache,
		db:    database.NewCompanyDatabase(db),
	}
}

func (c *CompanyRepository) GetById(ctx context.Context, id int64) (domain.Company, error) {
	return c.db.GetById(ctx, id)
}

func (c *CompanyRepository) IsExist(ctx context.Context, id int64) (bool, error) {
	return c.db.IsExist(ctx, id)
}
