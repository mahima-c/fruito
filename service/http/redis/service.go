package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	rdb *redis.Client
}

type Service interface {
	NewClient() (*Client, error)
	GetClient() (*Client, error)
}

type service struct {
	redisConn string
	client    *Client
}

func NewService(connStr string, client *Client) Service {
	return &service{redisConn: connStr, client: client}
}

func (s *service) NewClient() (*Client, error) {
	opt, err := redis.ParseURL(s.redisConn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	rdb := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	redisClient := &Client{rdb: rdb}
	s.client = redisClient

	log.Println("Redis connection established")
	return redisClient, nil
}

func (s *service) GetClient() (*Client, error) {
	redisClient := s.client
	if redisClient == nil {
		var err error
		redisClient, err = s.NewClient()
		if err != nil {
			return nil, err
		}
		return redisClient, nil
	}
	return redisClient, nil
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.rdb.Set(ctx, key, value, expiration).Err()
}

func (c *Client) Del(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

func (c *Client) Close() error {
	return c.rdb.Close()
}
