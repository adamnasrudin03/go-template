package main

import (
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app"
	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/router"
	"github.com/adamnasrudin03/go-template/pkg/database"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var (
	cfg                  = configs.GetInstance()
	logger               = driver.Logger(cfg)
	cache                = driver.Redis(cfg)
	db          *gorm.DB = database.SetupDbConnection()
	repo                 = app.WiringRepository(db, &cache, cfg, logger)
	services             = app.WiringService(repo, &cache, cfg, logger)
	controllers          = app.WiringDelivery(services, cfg, logger)
)

func init() {
	// set timezone local
	time.Local = helpers.TimeZoneJakarta()
}

func main() {
	defer database.CloseDbConnection(db)

	if cfg.App.UseRabbitMQ {
		go controllers.Message.Consume(&gin.Context{})
	}

	r := router.NewRoutes(*controllers)

	listen := fmt.Sprintf(":%v", cfg.App.Port)
	r.Run(listen)
}
