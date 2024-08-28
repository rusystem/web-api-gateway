package inmemory

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type User interface {
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	Add(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id int64) error
}

type UserCacheRepository struct {
	cfg   *config.Config
	cache *memcache.Client
}

func NewUserCacheRepository(cfg *config.Config, cache *memcache.Client) *UserCacheRepository {
	return &UserCacheRepository{
		cfg:   cfg,
		cache: cache,
	}
}

func (u *UserCacheRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	//todo
	return domain.User{}, nil
}

func (u *UserCacheRepository) Add(ctx context.Context, user domain.User) error {
	//todo
	return nil
}

func (u *UserCacheRepository) Delete(ctx context.Context, id int64) error {
	//todo
	return nil
}
