package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/mahima-c/fruito/common"
	"github.com/mahima-c/fruito/configs"
	"github.com/mahima-c/fruito/service/http/auth"
	"github.com/mahima-c/fruito/service/http/database"
	"github.com/mahima-c/fruito/service/http/health"
	"github.com/mahima-c/fruito/service/http/redis"
	"github.com/mahima-c/fruito/transport"
)

func main() {

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

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

	fmt.Println("Listening on", *httpAddr)
	// start server
	common.Go(func() {
		func() {
			err := http.ListenAndServe(*httpAddr, h)
			if err != nil {
				log.Fatal("Failed to start server!", err)
			}
		}()
	})
	fmt.Println("Listening on ---", *httpAddr)
	log.Println("Service startup ended!", fmt.Sprintf("Startup time: %d", time.Since(serviceStart).Milliseconds()))
}
