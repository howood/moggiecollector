package caches

import (
	"context"
	"fmt"
	"time"

	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/patrickmn/go-cache"
)

const (
	// NumInstance is number of instance
	NumInstance = 5
	// PurgeExpiredTime is time to purge cache
	PurgeExpiredTime = 10
)

//nolint:gochecknoglobals
var gocacheConnectionMap map[int]*cache.Cache

//nolint:gochecknoinits
func init() {
	gocacheConnectionMap = make(map[int]*cache.Cache, 0)
	for i := range NumInstance {
		//nolint:mnd
		gocacheConnectionMap[i] = cache.New(60*time.Minute, PurgeExpiredTime*time.Minute)
	}
}

// GoCacheClient struct
type GoCacheClient struct{}

// NewGoCacheClient creates a new GoCacheClient
func NewGoCacheClient() *GoCacheClient {
	ret := &GoCacheClient{}
	return ret
}

// Get gets from cache
func (cc *GoCacheClient) Get(ctx context.Context, key string) (interface{}, bool, error) {
	val, ok := cc.getInstance(ctx, key).Get(key)
	return val, ok, nil
}

// Set puts to cache
func (cc *GoCacheClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	cc.getInstance(ctx, key).Set(key, value, ttl)
	return nil
}

// Del deletes from cache
func (cc *GoCacheClient) Del(ctx context.Context, key string) error {
	cc.getInstance(ctx, key).Delete(key)
	return nil
}

// DelBulk bulk deletes from cache
func (cc *GoCacheClient) DelBulk(ctx context.Context, key string) error {
	cc.getInstance(ctx, key).Delete(key)
	return nil
}

// CloseConnect close connection
func (cc *GoCacheClient) CloseConnect() error {
	return nil
}

func (cc *GoCacheClient) getInstance(ctx context.Context, key string) *cache.Cache {
	// djb2 algorithm
	//nolint:mnd
	hash := uint32(5381)
	for _, c := range key {
		//nolint:mnd
		hash = ((hash << 5) + hash) + uint32(c)
	}
	i := int(hash) % NumInstance
	log.Info(ctx, fmt.Sprintf("get_instance: %d", i))
	return gocacheConnectionMap[i]
}
