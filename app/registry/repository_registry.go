package registry

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	authRepo "github.com/adamnasrudin03/go-template/app/modules/auth/repository"
	cacheRepo "github.com/adamnasrudin03/go-template/app/modules/cache/repository"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	messageRepo "github.com/adamnasrudin03/go-template/app/modules/message/repository"
	userRepo "github.com/adamnasrudin03/go-template/app/modules/user/repository"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repositories all repo object injected here
type Repositories struct {
	Auth    authRepo.AuthRepository
	Cache   cacheRepo.CacheRepository
	User    userRepo.UserRepository
	Log     logRepo.LogRepository
	Message messageRepo.MessageRepository
}

func WiringRepository(db *gorm.DB, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) *Repositories {
	return &Repositories{
		Auth:    authRepo.NewAuthRepository(db, cfg, logger),
		Cache:   cacheRepo.NewCacheRepository(*cache, cfg, logger),
		User:    userRepo.NewUserRepository(db, cfg, logger),
		Log:     logRepo.NewLogRepository(db, cfg, logger),
		Message: messageRepo.NewMessageRepository(db, cfg, logger),
	}
}
