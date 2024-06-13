package registry

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	authRepo "github.com/adamnasrudin03/go-template/app/modules/auth/repository"
	cacheRepo "github.com/adamnasrudin03/go-template/app/modules/cache/repository"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	userRepo "github.com/adamnasrudin03/go-template/app/modules/user/repository"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repositories all repo object injected here
type Repositories struct {
	Auth  authRepo.AuthRepository
	Cache cacheRepo.CacheRepository
	User  userRepo.UserRepository
	Log   logRepo.LogRepository
}

func WiringRepository(db *gorm.DB, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) *Repositories {
	return &Repositories{
		Auth:  regAuthRepo(db, cache, cfg, logger),
		Cache: cacheRepo.NewCacheRepository(*cache, cfg, logger),
		User:  userRepo.NewUserRepository(db, *cache, cfg, logger),
		Log:   logRepo.NewLogRepository(db, cfg, logger),
	}
}

func regAuthRepo(db *gorm.DB, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) authRepo.AuthRepository {
	return authRepo.NewAuthRepository(authRepo.AuthRepo{
		DB:     db,
		Cache:  *cache,
		Cfg:    cfg,
		Logger: logger,
	})
}
