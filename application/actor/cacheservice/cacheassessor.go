package cacheservice

import (
	"context"
	"strconv"
	"time"

	"github.com/howood/moggiecollector/infrastructure/client/caches"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/library/utils"
)

// CacheAssessor struct
type CacheAssessor struct {
	instance caches.CacheInstance
}

// NewCacheAssessor creates a new CacheAssessor
func NewCacheAssessor(ctx context.Context) *CacheAssessor {
	var I *CacheAssessor
	log.Debug(ctx, "use:"+utils.GetOsEnv("CACHE_TYPE", "gocache"))
	switch utils.GetOsEnv("CACHE_TYPE", "gocache") {
	case "gocache":
		I = &CacheAssessor{
			instance: caches.NewGoCacheClient(),
		}
	default:
		I = &CacheAssessor{
			instance: caches.NewGoCacheClient(),
		}
	}
	return I
}

// Get returns cache contents
func (ca *CacheAssessor) Get(index string) (interface{}, bool) {
	defer func() {
		//nolint:errcheck
		ca.instance.CloseConnect()
	}()
	cachedvalue, cachedfound := ca.instance.Get(index)
	if cachedfound {
		return cachedvalue, true
	}
	return "", false
}

// Set puts cache contents
func (ca *CacheAssessor) Set(index string, value interface{}, expired time.Duration) error {
	defer func() {
		//nolint:errcheck
		ca.instance.CloseConnect()
	}()
	//nolint:durationcheck
	return ca.instance.Set(index, value, expired*time.Second)
}

// Delete remove cache contents
func (ca *CacheAssessor) Delete(index string) error {
	defer func() {
		//nolint:errcheck
		ca.instance.CloseConnect()
	}()
	return ca.instance.Del(index)
}

// GetChacheExpired get cache expired
func GetChacheExpired() time.Duration {
	expired, err := strconv.Atoi(utils.GetOsEnv("CACHE_EXPIED", "3600"))
	if err != nil {
		panic(err)
	}
	return time.Duration(expired)
}
