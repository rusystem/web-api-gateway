package repository

import (
	"database/sql"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/web-api-gateway/internal/config"
)

type Repository struct {
	Auth *AuthRepository
}

func New(cfg *config.Config, cache *memcache.Client, pc *sql.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(cfg, cache, pc),
	}
}
