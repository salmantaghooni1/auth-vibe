package main

import (
	// gin-swagger middleware

	"fmt"

	"github.com/salmantaghooni/golang-car-web-api/api"
	"github.com/salmantaghooni/golang-car-web-api/config"
	"github.com/salmantaghooni/golang-car-web-api/data/cache"
	"github.com/salmantaghooni/golang-car-web-api/data/db"
	"github.com/salmantaghooni/golang-car-web-api/data/db/migrations"
	"github.com/salmantaghooni/golang-car-web-api/pkg/logging"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	fmt.Println(logger)
	err := cache.InitRedis(cfg)
	fmt.Println(err)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()
	fmt.Println("err")
	if err := db.InitDb(cfg); err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	fmt.Println(err)
	migrations.Up_1()
	defer db.CloseDb()
	api.InitServer(cfg)

}
