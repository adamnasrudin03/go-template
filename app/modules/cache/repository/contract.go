package repository

import (
	"context"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"github.com/sirupsen/logrus"
)

type CacheRepository interface {
	CreateCache(ctx context.Context, key string, data interface{}, ttl time.Duration)
	GetCache(ctx context.Context, key string, res interface{}) bool
	DelCache(ctx context.Context, key string) error
}

type cacheRepo struct {
	Cache  driver.RedisClient
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewCacheRepository(
	cache driver.RedisClient,
	cfg *configs.Configs,
	logger *logrus.Logger,
) CacheRepository {
	return &cacheRepo{
		Cache:  cache,
		Cfg:    cfg,
		Logger: logger,
	}
}
