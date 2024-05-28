package app

import (
	logDelivery "github.com/adamnasrudin03/go-template/app/modules/log/delivery"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	logSrv "github.com/adamnasrudin03/go-template/app/modules/log/service"

	userDelivery "github.com/adamnasrudin03/go-template/app/modules/user/delivery"
	userRepo "github.com/adamnasrudin03/go-template/app/modules/user/repository"
	userSrv "github.com/adamnasrudin03/go-template/app/modules/user/service"

	"github.com/adamnasrudin03/go-template/app/registry"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB, cache *driver.RedisClient) *registry.Repositories {
	return &registry.Repositories{
		User: userRepo.NewUserRepository(db, *cache),
		Log:  logRepo.NewLogRepository(db, *cache),
	}
}

func WiringService(repo *registry.Repositories, cache *driver.RedisClient) *registry.Services {
	return &registry.Services{
		User: userSrv.NewUserService(repo.User, *cache),
		Log:  logSrv.NewLogService(repo.Log),
	}
}

func WiringDelivery(srv *registry.Services) *registry.Deliveries {
	return &registry.Deliveries{
		User: userDelivery.NewUserDelivery(srv.User),
		Log:  logDelivery.NewLogDelivery(srv.Log),
	}
}
