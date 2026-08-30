package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // the live my-redis Docker container
	})
	defer rdb.Close()

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Println("cannot reach Redis:", err)
		return
	}
	fmt.Println("connected to Redis at localhost:6379")
	fmt.Println()

	demoIdempotency(ctx, rdb)
	fmt.Println()

	demoRateLimiter(ctx, rdb)
	fmt.Println()

	demoCrashRecovery(ctx, rdb)
}
