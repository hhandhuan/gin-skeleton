package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hhandhuan/gin-skeleton/configs"
	"log"
)

var Redis *redis.Client

func RedisInit() {
	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			configs.Conf.Redis.Host,
			configs.Conf.Redis.Port,
		),
		Password: configs.Conf.Redis.Password,
		DB:       configs.Conf.Redis.DB,
		PoolSize: 50,
	})
	str, err := Redis.Ping(context.Background()).Result()
	if err != nil || str != "PONG" {
		log.Fatalf("Redis connect ping failed, err: %s", err.Error())
	}
}
