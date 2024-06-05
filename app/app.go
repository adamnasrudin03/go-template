package app

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	logDelivery "github.com/adamnasrudin03/go-template/app/modules/log/delivery"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	logSrv "github.com/adamnasrudin03/go-template/app/modules/log/service"

	userDelivery "github.com/adamnasrudin03/go-template/app/modules/user/delivery"
	userRepo "github.com/adamnasrudin03/go-template/app/modules/user/repository"
	userSrv "github.com/adamnasrudin03/go-template/app/modules/user/service"

	messageDelivery "github.com/adamnasrudin03/go-template/app/modules/message/delivery"
	"github.com/adamnasrudin03/go-template/app/registry"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) *registry.Repositories {
	return &registry.Repositories{
		User: userRepo.NewUserRepository(db, *cache, cfg, logger),
		Log:  logRepo.NewLogRepository(db, *cache, cfg, logger),
	}
}

func WiringService(repo *registry.Repositories, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) *registry.Services {
	return &registry.Services{
		User: userSrv.NewUserService(repo.User, *cache, logger),
		Log:  logSrv.NewLogService(repo.Log, logger),
	}
}

func WiringDelivery(srv *registry.Services, cfg *configs.Configs, logger *logrus.Logger) *registry.Deliveries {
	return &registry.Deliveries{
		User:    userDelivery.NewUserDelivery(srv.User, logger),
		Log:     logDelivery.NewLogDelivery(srv.Log, logger),
		Message: messageDelivery.NewMessageDelivery(srv.Log, cfg, logger),
	}
}
