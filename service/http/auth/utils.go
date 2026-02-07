package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Mrhb787/hospital-ward-manager/model"
	redisUtil "github.com/redis/go-redis/v9"
)

func (s *service) GetTokenFromCache(ctx context.Context, phone string, userId int) (token string, is_expired bool, err error) {
	redisClient, err := s.redisService.GetClient()
	if err != nil {
		log.Println("Redis connection failed")
		return token, false, err
	}

	// eg: USER_TOKEN_<phonenumber>
	cacheKey := fmt.Sprintf("USER_TOKEN_%s", phone)
	cacheToken, gErr := redisClient.Get(ctx, cacheKey)
	if gErr != nil && gErr != redisUtil.Nil {
		return token, false, gErr
	}

	if gErr == redisUtil.Nil || cacheToken == "" {
		// get from db and set cache

		log.Println("cache key not found")

		// db client
		dbClient, dErr := s.dbService.GetClient()
		if dErr != nil {
			return token, false, dErr
		}

		tx, txErr := dbClient.DB.Begin()
		if txErr != nil {
			log.Println("Failed to begin transaction", txErr)
			return token, false, txErr
		}

		defer func() {
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()

		var userSession *model.UserSession
		userSession, err = model.GetLatestUserSessionByUserId(tx, userId)
		if err != nil {
			return token, false, err
		}

		// session doesn't exists
		if userSession == nil {
			return token, true, nil
		}

		// session expired
		if time.Now().UTC().After(userSession.CreatedAt.UTC().Add(time.Hour * 24 * 60)) {
			return token, true, nil
		}

		// set cache for 1 day
		err = redisClient.Set(ctx, cacheKey, userSession.Token, time.Hour*24)
		if err != nil {
			return token, false, err
		}

		log.Println("cache set successfully", cacheKey)

		return userSession.Token, false, nil
	}

	log.Println("token served from redis cache", cacheKey, cacheToken)

	return cacheToken, false, nil
}
