package inmemory

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/web-api-gateway/internal/config"
)

type Auth interface {
}

type AuthCacheRepository struct {
	cfg   *config.Config
	cache *memcache.Client
}

func NewAuthCacheRepository(cfg *config.Config, cache *memcache.Client) *AuthCacheRepository {
	return &AuthCacheRepository{
		cfg:   cfg,
		cache: cache,
	}
}
