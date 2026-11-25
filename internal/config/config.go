package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Mongo    MongoConfig
	Redis    RedisConfig
	NATS     NATSConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DB)
}

type MongoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func (c MongoConfig) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		c.User, c.Password, c.Host, c.Port, c.DB)
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func (c RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

type NATSConfig struct {
	URL string
}

type AuthConfig struct {
	JWTSecret    string
	APIKeyPrefix string
}

func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getenv("SERVER_PORT", "8080"),
			Env:  getenv("SERVER_ENV", "development"),
		},
		Postgres: PostgresConfig{
			Host:     getenv("PG_HOST", "localhost"),
			Port:     getenv("PG_PORT", "5432"),
			User:     getenv("PG_USER", "gameforge"),
			Password: getenv("PG_PASSWORD", "changeme"),
			DB:       getenv("PG_DB", "gameforge"),
		},
		Mongo: MongoConfig{
			Host:     getenv("MONGO_HOST", "localhost"),
			Port:     getenv("MONGO_PORT", "27017"),
			User:     getenv("MONGO_USER", "gameforge"),
			Password: getenv("MONGO_PASSWORD", "changeme"),
			DB:       getenv("MONGO_DB", "gameforge"),
		},
		Redis: RedisConfig{
			Host:     getenv("REDIS_HOST", "localhost"),
			Port:     getenv("REDIS_PORT", "6379"),
			Password: getenv("REDIS_PASSWORD", "changeme"),
		},
		NATS: NATSConfig{
			URL: getenv("NATS_URL", "nats://localhost:4222"),
		},
		Auth: AuthConfig{
			JWTSecret:    getenv("JWT_SECRET", "changeme-jwt-secret"),
			APIKeyPrefix: getenv("API_KEY_PREFIX", "gf_"),
		},
	}
	return cfg, nil
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

var (
	ShutdownTimeout   = 30 * time.Second
	RuleEvalTimeout   = 100 * time.Millisecond
	WSDebounceWindow  = 100 * time.Millisecond
)
