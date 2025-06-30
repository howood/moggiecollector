package service

import (
	"context"
	"time"

	"github.com/howood/moggiecollector/infrastructure/client/caches"
	"github.com/howood/moggiecollector/library/utils"
)

type AuthCacheService interface {
	Set(ctx context.Context, key string, value interface{}, expired time.Duration) error
	Get(ctx context.Context, key string) (interface{}, bool, error)
	Del(ctx context.Context, key string) error
	DelBulk(ctx context.Context, key string) error
}

type authCacheService struct {
	cacheInstance caches.CacheInstance
}

// NewAuthCacheService returns AuthCacheService interface.
//
//nolint:ireturn,nolintlint
func NewAuthCacheService() AuthCacheService {
	instance := caches.NewRedis(true, utils.GetOsEnvInt("AUTH_CACHED_DB", 0))

	return &authCacheService{
		cacheInstance: instance,
	}
}

func (d *authCacheService) Set(ctx context.Context, key string, value interface{}, expired time.Duration) error {
	//nolint:wrapcheck
	return d.cacheInstance.Set(ctx, key, value, expired)
}

func (d *authCacheService) Get(ctx context.Context, key string) (interface{}, bool, error) {
	//nolint:wrapcheck
	return d.cacheInstance.Get(ctx, key)
}

func (d *authCacheService) Del(ctx context.Context, key string) error {
	//nolint:wrapcheck
	return d.cacheInstance.Del(ctx, key)
}

func (d *authCacheService) DelBulk(ctx context.Context, key string) error {
	//nolint:wrapcheck
	return d.cacheInstance.DelBulk(ctx, key)
}
