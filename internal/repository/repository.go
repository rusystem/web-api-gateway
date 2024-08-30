package repository

import (
	"database/sql"
	"github.com/rusystem/cache"
	"github.com/rusystem/web-api-gateway/internal/config"
)

type Repository struct {
	Auth Auth
	User User
}

func New(cfg *config.Config, cache *cache.MemoryCache, pc *sql.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(cfg, cache, pc),
		User: NewUserRepository(cfg, cache, pc),
	}
}
