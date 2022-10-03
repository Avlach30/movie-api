package redis

import (
	"log"
	"movie-api/helper"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(host string, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: host,
		Password: password,
		DB: 0,
	})

	return client, nil
}

func RedisConnect() *redis.Client {
	get := helper.GetEnvValue

	REDIS_HOST := get("REDIS_HOST")
	REDIS_PASSWORD := ""

	redisConnect, err  := NewRedisClient(REDIS_HOST, REDIS_PASSWORD)
	if (err != nil) {
		log.Fatal(err.Error())
	}

	return redisConnect
}

