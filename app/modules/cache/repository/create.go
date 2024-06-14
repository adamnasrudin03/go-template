package repository

import (
	"context"
	"time"

	"github.com/adamnasrudin03/go-template/app/models"
)

func (r *cacheRepo) CreateCache(ctx context.Context, key string, data interface{}, ttl time.Duration) {
	var (
		opName = "CacheRepository-CreateCache"
		err    error
	)
	if ttl == 0 || ttl <= models.TimeDurationZero {
		ttl = r.Cfg.Redis.DefaultCacheTimeOut
	}

	err = r.Cache.Set(key, data, ttl)
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return
	}
}
