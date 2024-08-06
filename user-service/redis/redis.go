package redis

import (
	"context"
	"fmt"
	"user-service/config"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg config.Config) (*redis.Client, error) {
	target := fmt.Sprintf("%s:%s", cfg.Redis.RedisHost, cfg.Redis.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr: target,
	})

	if err := rdb.Ping(context.Background()); err != nil {
		return rdb, nil
	}

	return rdb, nil

}
