package redis

import (
	"context"
	"errors"
	"time"
	"github.com/go-redis/redis/v8"
)

type Service interface {
	GetRedisFromKey(key string) (int, error)
}

func SetRedisNew(key string, value any) (*redis.StatusCmd, error) {
	ttl := time.Duration(1) * time.Hour

	setRedis := RedisConnect().Set(context.Background(), key, value, ttl)
	if err := setRedis.Err(); err != nil {
		return setRedis, errors.New("unable to SET data in redis")
	}

	return setRedis, nil
}

func GetRedisFromKey(key string) (int, error) {
	getRedis := RedisConnect().Get(context.Background(), key)
	if err := getRedis.Err(); err != nil {
		return 0, errors.New("unable to GET data in redis")
	}

	data, err := getRedis.Int()
	if err != nil {
		return 0, errors.New("unable to SET data in redis")
	}

	return data, nil
}