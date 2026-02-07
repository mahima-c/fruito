package configs

import (
	"os"

	"github.com/Mrhb787/hospital-ward-manager/common"
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
	Host string
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
		Host: getEnv("REDIS_URL", ""),
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
