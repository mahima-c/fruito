package transport

import (
	"context"

	"github.com/mahima-c/fruito/common"
	"github.com/mahima-c/fruito/service/http/auth"
	"github.com/mahima-c/fruito/service/http/database"
	"github.com/mahima-c/fruito/service/http/health"
	"github.com/mahima-c/fruito/service/http/redis"
)

type Endpoints struct {
	HeathCheckEndpoint     common.Endpoint
	SignInUserEndpoint     common.Endpoint
	UpsertProductEndpoint  common.Endpoint
	GetProductByIdEndpoint common.Endpoint
	GetAllProductsEndpoint common.Endpoint
}

func MakeHealthEndpoints(s health.Service) Endpoints {
	return Endpoints{
		HeathCheckEndpoint: MakeHealthEndpoint(s),
	}
}

func MakeHttpServiceEndpoints(dbService database.Service, redisService redis.Service, authService auth.Service) Endpoints {
	return Endpoints{
		SignInUserEndpoint:     MakeSignInUserEndpoint(authService),
		UpsertProductEndpoint:  MakeUpsertProductEndpoint(dbService),
		GetProductByIdEndpoint: MakeGetProductByIdEndpoint(dbService),
		GetAllProductsEndpoint: MakeGetAllProductsEndpoint(dbService),
	}
}

func MakeHealthEndpoint(s health.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := s.Health(ctx)
		return healthResponse{
			Status: status,
		}, err
	}
}

func MakeSignInUserEndpoint(authService auth.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SignInUserRequest)

		resp, err := authService.SignIn(ctx, auth.SignInUserRequest{PhoneNumber: req.Phone, Password: req.Password})
		if err != nil {
			return resp, err
		}

		return resp, nil
	}
}
