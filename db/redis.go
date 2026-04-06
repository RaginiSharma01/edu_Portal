package db

import (
	"smp/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Config) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	return rdb
}
