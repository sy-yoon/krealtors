package utils

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func RedisInit() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "192.168.140.130:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return client
}
