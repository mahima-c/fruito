package transport

import (
	"net/http"

	"github.com/mahima-c/fruito/common"
	"github.com/mahima-c/fruito/service/http/auth"
	"github.com/mahima-c/fruito/service/http/database"
	"github.com/mahima-c/fruito/service/http/health"
	"github.com/mahima-c/fruito/service/http/redis"
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

	// upsert product
	r.Methods(common.POST.ToString()).Path("/product").Handler(common.NewServer(
		serviceEndpoints.UpsertProductEndpoint,
		DecodeUpsertProductRequest,
		EncodeGenericResponse,
		optionsWithoutAuth...,
	))

	// get product by id
	r.Methods(common.GET.ToString()).Path("/product").Queries("id", "{id}").Handler(common.NewServer(
		serviceEndpoints.GetProductByIdEndpoint,
		DecodeGetProductByIdRequest,
		EncodeGenericResponse,
		optionsWithoutAuth...,
	))

	// get all products
	r.Methods(common.GET.ToString()).Path("/product").Handler(common.NewServer(
		serviceEndpoints.GetAllProductsEndpoint,
		DecodeGetAllProductsRequest,
		EncodeGenericResponse,
		optionsWithoutAuth...,
	))

	return r
}
