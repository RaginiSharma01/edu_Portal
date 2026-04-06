package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func StoreOTP(ctx context.Context, rdb *redis.Client, email string, otp string) error {

	key := "otp:" + email

	return rdb.Set(ctx, key, otp, 5*time.Minute).Err()
}
func GetOTP(ctx context.Context, rdb *redis.Client, email string) (string, error) {

	key := "otp:" + email
	fmt.Println("Redis:", rdb)
	return rdb.Get(ctx, key).Result()
}
