package repository

import (
	"database/sql"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/web-api-gateway/internal/config"
)

type Auth interface {
}

type AuthRepository struct {
	cfg   *config.Config
	cache *memcache.Client
	db    *sql.DB
}

func NewAuthRepository(cfg *config.Config, cache *memcache.Client, db *sql.DB) *AuthRepository {
	return &AuthRepository{
		cfg:   cfg,
		cache: cache,
		db:    db,
	}
}
