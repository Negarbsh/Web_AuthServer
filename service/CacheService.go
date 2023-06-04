package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client
var initialized = false

func startRedisConnection() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})
	initialized = true
}

func GetValue(ctx context.Context, key string) (string, error) {
	if !initialized {
		startRedisConnection()
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error while getting data: ", err)
		return "", err
	}
	return val, nil
}

func CacheData(ctx context.Context, key string, value string, expirationDuration time.Duration) {
	if !initialized {
		startRedisConnection()
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	if expirationDuration == -1 {
		expirationDuration = redis.KeepTTL
	}
	err := client.Set(ctx, key, value, expirationDuration).Err()
	if err != nil {
		fmt.Println("Error while caching data: ", err)
	} else {
		fmt.Println("Successfully cached data with key ", key, " and value ", value)
	}
}

func DeleteValue(ctx context.Context, key string) {
	if !initialized {
		startRedisConnection()
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	err := client.Del(ctx, key).Err()
	if err != nil {
		fmt.Println("Error while deleting data: ", err)
	} else {
		fmt.Println("Successfully deleted data with key ", key)
	}
}
