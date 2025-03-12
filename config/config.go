package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	RedisClient   *redis.Client
	Ctx           = context.Background()
	CurrentConfig *Config
)

// Config holds application configuration values.
type Config struct {
	RedisAddr     string `mapstructure:"redis_addr"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
	JWTSecret     string `mapstructure:"jwt_secret"`
}

// LoadConfig loads configuration using Viper.
func LoadConfig() (*Config, error) {
	// Set the config file name (without extension) and type.
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// Look for the config file in the current directory.
	viper.AddConfigPath(".")

	// Set default values.
	viper.SetDefault("redis_addr", "localhost:6379")
	viper.SetDefault("redis_password", "")
	viper.SetDefault("redis_db", 0)
	viper.SetDefault("jwt_secret", "my_secret_key")

	// Read in the config file.
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error reading config file: %w", err)
	}

	// Unmarshal the config into the Config struct.
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	CurrentConfig = &cfg
	return &cfg, nil
}

// InitRedis initializes the Redis client using configuration values.
func InitRedis(cfg *Config) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	_, err := RedisClient.Ping(Ctx).Result()
	return err
}
