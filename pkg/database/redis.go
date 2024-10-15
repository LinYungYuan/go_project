package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(addr string) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	return err
}
