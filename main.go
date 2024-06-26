package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/registry"
	"github.com/adamnasrudin03/go-template/app/router"
	"github.com/adamnasrudin03/go-template/pkg/database"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

func init() {
	// set timezone local
	time.Local = helpers.TimeZoneJakarta()

	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load env file")
	}

}

func main() {
	var (
		cfg                  = configs.GetInstance()
		logger               = driver.Logger(cfg)
		cache                = driver.Redis(cfg)
		db          *gorm.DB = database.SetupDbConnection(cfg, logger)
		repo                 = registry.WiringRepository(db, &cache, cfg, logger)
		services             = registry.WiringService(repo, &cache, cfg, logger)
		controllers          = registry.WiringDelivery(services, cfg, logger)
	)

	defer database.CloseDbConnection(db, logger)

	if cfg.App.UseRabbitMQ {
		go services.Msg.Consume(context.Background())
	}

	r := router.NewRoutes(*controllers)

	listen := fmt.Sprintf(":%v", cfg.App.Port)
	r.Run(listen)
}
