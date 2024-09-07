package repository

import (
	"context"
	"database/sql"
	"github.com/rusystem/cache"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository/database"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type User interface {
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	GetSections(ctx context.Context, id int64) ([]string, error)
	GetById(ctx context.Context, id int64) (domain.User, error)
	UpdateLastLogin(ctx context.Context, id int64) error
}

type UserRepository struct {
	cfg   *config.Config
	cache *cache.MemoryCache
	db    database.User
}

func NewUserRepository(cfg *config.Config, cache *cache.MemoryCache, db *sql.DB) *UserRepository {
	return &UserRepository{
		cfg:   cfg,
		cache: cache,
		db:    database.NewUserDatabase(db),
	}
}

func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	return ur.db.GetByUsername(ctx, username)
}

func (ur *UserRepository) GetSections(ctx context.Context, id int64) ([]string, error) {
	return ur.db.GetSections(ctx, id)
}

func (ur *UserRepository) GetById(ctx context.Context, id int64) (domain.User, error) {
	return ur.db.GetById(ctx, id)
}

func (ur *UserRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	return ur.db.UpdateLastLogin(ctx, id)
}
