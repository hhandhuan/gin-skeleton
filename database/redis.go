package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/hhandhuan/gin-skeleton/logger"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			configs.Conf.Redis.Host,
			configs.Conf.Redis.Port,
		),
		Password: configs.Conf.Redis.Password,
		DB:       configs.Conf.Redis.DB,
		PoolSize: 10,
	})
	str, err := Redis.Ping(context.Background()).Result()
	if err != nil || str != "PONG" {
		logger.I.Fatalf("Redis connect ping failed, err: %v", err)
	}
}
