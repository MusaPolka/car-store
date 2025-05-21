package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

func NewRedisClient(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // e.g. "localhost:6379"
		Password: password, // "" if no password set
		DB:       db,       // 0 for default DB
	})
	return rdb
}
