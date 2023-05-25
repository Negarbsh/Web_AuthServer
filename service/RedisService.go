package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
func startConnection() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB

	}
}

func getValue(ctx context.Context, key string) string {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
	return val
}
