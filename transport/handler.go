package transport

import (
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/service/http/auth"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/service/http/redis"
)

type HttpHandlerRequest struct {
	HealthService health.Service
	DbService     database.Service
	RedisService  redis.Service
	AuthService   auth.Service
}

func NewHandler(req HttpHandlerRequest) http.Handler {
	// router
	r := common.NewRouter()

	// endpoints initialize

	// health endpoints
	healthEndpoints := MakeHealthEndpoints(req.HealthService)

	// service endpoints
	serviceEndpoints := MakeHttpServiceEndpoints(req.DbService, req.RedisService, req.AuthService)

	// health check
	r.Methods(common.GET.ToString()).Path("/health").Handler(common.NewServer(
		healthEndpoints.HeathCheckEndpoint,
		DecodeHealthRequest,
		EncodeHealthResponse,
		optionsWithoutRouteCheck...,
	))

	// sign in
	r.Methods(common.POST.ToString()).Path("/sign_in").Handler(common.NewServer(
		serviceEndpoints.SignInUserEndpoint,
		DecodeSignInUserRequest,
		EncodeGenericResponse,
		optionsWithoutAuth...,
	))

	return r
}
