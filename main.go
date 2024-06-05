package main

import (
	"fmt"

	"github.com/adamnasrudin03/go-template/app"
	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/router"
	"github.com/adamnasrudin03/go-template/pkg/database"
	"github.com/adamnasrudin03/go-template/pkg/driver"

	"gorm.io/gorm"
)

var (
	config               = configs.GetInstance()
	logger               = driver.Logger(config)
	cache                = driver.Redis(config)
	db          *gorm.DB = database.SetupDbConnection()
	repo                 = app.WiringRepository(db, &cache, logger)
	services             = app.WiringService(repo, &cache, logger)
	controllers          = app.WiringDelivery(services, logger)
)

func main() {
	defer database.CloseDbConnection(db)

	r := router.NewRoutes(*controllers)

	listen := fmt.Sprintf(":%v", config.App.Port)
	r.Run(listen)
}
