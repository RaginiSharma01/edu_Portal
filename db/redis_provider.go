package db

import (
	"smp/config"

	"github.com/redis/go-redis/v9"
)

func ProvideRedis(cfg *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	return client
}
