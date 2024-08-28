package repository

import (
	"context"
	"database/sql"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository/database"
	"github.com/rusystem/web-api-gateway/internal/repository/inmemory"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type User interface {
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	GetSections(ctx context.Context, id int64) ([]string, error)
	GetById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, user domain.User) error
	UpdateLastLogin(ctx context.Context, id int64) error
}

type UserRepository struct {
	cfg   *config.Config
	cache inmemory.User
	db    database.User
}

func NewUserRepository(cfg *config.Config, cache *memcache.Client, db *sql.DB) *UserRepository {
	return &UserRepository{
		cfg:   cfg,
		cache: inmemory.NewUserCacheRepository(cfg, cache),
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

func (ur *UserRepository) Create(ctx context.Context, user domain.User) (int64, error) {
	return ur.db.Create(ctx, user)
}

func (ur *UserRepository) Update(ctx context.Context, user domain.User) error {
	return ur.db.Update(ctx, user)
}

func (ur *UserRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	return ur.db.UpdateLastLogin(ctx, id)
}
