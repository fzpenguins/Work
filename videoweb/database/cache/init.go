package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"videoweb/config"
)

var RedisClient *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Username:    config.RedisAddr,
		Password:    config.RedisPassword,
		DB:          config.RedisDB,
		DialTimeout: 5 * time.Second,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
