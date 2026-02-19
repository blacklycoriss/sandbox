package db

import "github.com/redis/go-redis/v9"

var Redis *redis.Client

func InitRedis(addr string) {
	Redis = redis.NewClient(&redis.Options{Addr: addr})
}
