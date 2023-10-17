package redis

import "os"

type RedisConfig struct {
	Host string
	DB   string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		DB:   os.Getenv("REDIS_DB"),
	}
}
