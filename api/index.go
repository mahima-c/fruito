package api

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/configs"
	"github.com/Mrhb787/hospital-ward-manager/service/http/auth"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/service/http/redis"
	"github.com/Mrhb787/hospital-ward-manager/transport"
)

//go:embed main.html
var mainHTML []byte

var httpTransportRequest transport.HttpHandlerRequest

func init() {
	// app config
	appConfig := configs.New()

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

	// health service
	healthService := health.NewService()

	// map http services
	httpTransportRequest.DbService = dbService
	httpTransportRequest.HealthService = healthService
	httpTransportRequest.AuthService = authService
	httpTransportRequest.HealthService = healthService
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		// Serve index.html
		serveMainHTML(w)
		return
	}

	// http handler
	h := transport.NewHandler(httpTransportRequest)

	// serve http
	h.ServeHTTP(w, r)

}

func serveMainHTML(w http.ResponseWriter) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the embedded HTML to the response
	w.Write(mainHTML)
}
