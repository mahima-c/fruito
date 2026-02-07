package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/configs"
	"github.com/Mrhb787/hospital-ward-manager/service/http/auth"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/service/http/redis"
	"github.com/Mrhb787/hospital-ward-manager/transport"
)

func main() {

	// service begin
	serviceStart := time.Now()

	// get app config
	appConfig := configs.New()

	// initalize services
	healthService := health.NewService()

	// database service
	dbService := database.NewService(appConfig.DBConfig.Host, nil)
	dbClient, err := dbService.NewClient()
	if err != nil || dbClient == nil {
		log.Fatal("Failed to connect database")
	}

	// redis service
	redisService := redis.NewService(appConfig.RedisConfig.Host, nil)
	redisClient, err := redisService.NewClient()
	if err != nil || redisClient == nil {
		log.Fatal("Failed to connect redis")
	}

	// auth service
	authService := auth.NewService(dbService, redisService)

	// http handler
	h := transport.NewHandler(transport.HttpHandlerRequest{
		HealthService: healthService,
		DbService:     dbService,
		RedisService:  redisService,
		AuthService:   authService,
	})
	addr := fmt.Sprintf(":%s", appConfig.ServiceConfig.ServicePort)
	httpAddr := flag.String("http.addr", addr, "HTTP listen address")
	flag.Parse()

	// start server
	common.Go(func() {
		func() {
			err := http.ListenAndServe(*httpAddr, h)
			if err != nil {
				log.Fatal("Failed to start server!", err)
			}
		}()
	})
	log.Println("Service startup ended!", fmt.Sprintf("Startup time: %d", time.Since(serviceStart).Milliseconds()))
}
