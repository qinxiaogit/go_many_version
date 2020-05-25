package go_redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
"time"
)

/**
redis连接池
*/
func NewRedisClient(Addr string, PoolSize int, db int, Password string) *redis.Client {
	fmt.Println("v1 go mod")
	if Addr == "" {
		Addr = ":6379"
	}
	if PoolSize <= 0 {
		PoolSize = 10
	}
	if db < 0 {
		db = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         Addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     PoolSize,
		PoolTimeout:  30 * time.Second,
		DB:           db,
		Password:     Password,
	})
	return rdb
}
