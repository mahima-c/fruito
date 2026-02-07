package configs

import (
	"os"

	"github.com/mahima-c/fruito/common"
)

type ServiceConfig struct {
	ServicePort string
}

type APIConfig struct {
	ApiKey string
}

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

type RedisConfig struct {
	Host     string
	Addr     string
	Username string
	Password string
	DB       int
}

type AppConfig struct {

	// api configs
	APIConfig

	// database configs
	DBConfig

	// redis configs
	RedisConfig

	// service configs
	ServiceConfig
}

func (db DBConfig) new() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", ""),
		Name:     getEnv("DB_NAME", ""),
		User:     getEnv("DB_ADMIN", ""),
		Password: getEnv("DB_ADMIN_PASSWORD", ""),
	}
}

func (db RedisConfig) new() RedisConfig {
	return RedisConfig{
		Addr:     getEnv("REDIS_ADDR", ""),
		Username: getEnv("REDIS_USER", ""),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       0,
	}

}

func (api APIConfig) new() APIConfig {
	return APIConfig{
		ApiKey: getEnv("API_SECRET_KEY", ""),
	}
}

func (s ServiceConfig) new() ServiceConfig {
	return ServiceConfig{
		ServicePort: string(common.LOCAL_SERVICE_PORT),
	}
}

func New() *AppConfig {
	return &AppConfig{
		DBConfig:      DBConfig{}.new(),
		RedisConfig:   RedisConfig{}.new(),
		APIConfig:     APIConfig{}.new(),
		ServiceConfig: ServiceConfig{}.new(),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
