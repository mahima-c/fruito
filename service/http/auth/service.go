package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/configs"
	"github.com/Mrhb787/hospital-ward-manager/model"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/redis"
)

type Service interface {
	SignIn(ctx context.Context, req SignInUserRequest) (resp SignInUserResponse, err error)
	ValidateToken(ctx context.Context, req ValidateTokenRequest) (resp ValidateTokenResponse, err error)
}

type service struct {
	dbService    database.Service
	redisService redis.Service
}

func NewService(dbService database.Service, redisService redis.Service) Service {
	return &service{dbService: dbService, redisService: redisService}
}

func (s *service) SignIn(ctx context.Context, req SignInUserRequest) (resp SignInUserResponse, err error) {

	log.Println("signIn req", req)

	// 1. get user by phone number
	var user model.User
	user, err = s.dbService.GetUserByPhone(req.PhoneNumber)
	if err != nil {
		return resp, err
	}

	log.Println("user found", user)

	// 2. get token from cache or db
	var token string
	var is_token_expired bool
	token, is_token_expired, err = s.GetTokenFromCache(ctx, req.PhoneNumber, user.ID)
	if err != nil {
		return resp, err
	}

	log.Println("token", token)
	log.Println("token expired", is_token_expired)

	// generate new token and session
	if is_token_expired {

		appConfig := configs.New()
		if appConfig == nil {
			return resp, errors.New("something went wrong")
		}

		// new jwt token
		token, err = common.GenerateJWT(int64(user.ID), appConfig.ApiKey, time.Hour*24*60)
		if err != nil {
			return resp, err
		}

		userSession := model.UserSession{
			UserID: user.ID,
			Token:  token,
		}

		// create new session
		err = s.dbService.CreateUserSession(userSession)
		if err != nil {
			return resp, err
		}
	}

	resp.User = user
	resp.Token = token
	return resp, nil
}

func (s *service) ValidateToken(ctx context.Context, req ValidateTokenRequest) (resp ValidateTokenResponse, err error) {

	appConfig := configs.New()
	if appConfig == nil {
		return resp, errors.New("something went wrong")
	}

	claims, cErr := common.VerifyJWT(req.Token, appConfig.ApiKey)
	if cErr != nil {
		return resp, cErr
	}

	// token valid
	resp.IsValid = true

	// token expired
	if claims.ExpiresAt.UTC().Before(time.Now().UTC()) {
		resp.IsValid = false
		return
	}

	return resp, err
}
