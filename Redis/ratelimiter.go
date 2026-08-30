package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// allow implements a fixed-window rate limiter: INCR a per-key counter,
// set it to expire after the window on its FIRST increment only, and
// reject once the count exceeds the limit. Simpler than a true token
// bucket, but the same production idea as the roadmap's "custom token
// bucket, per-key rate limiting" — a real per-user/per-API-key limiter.
func allow(ctx context.Context, rdb *redis.Client, key string, limit int64, window time.Duration) bool {
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println("redis error:", err)
		return false
	}

	if count == 1 {
		rdb.Expire(ctx, key, window) // only arm the TTL on the first hit in this window
	}

	return count <= limit
}

func demoRateLimiter(ctx context.Context, rdb *redis.Client) {
	fmt.Println("--- Redis-backed rate limiter ---")

	key := "ratelimit:user_42"
	rdb.Del(ctx, key) // clean slate

	const limit = 3
	for i := 1; i <= 5; i++ {
		if allow(ctx, rdb, key, limit, 10*time.Second) {
			fmt.Printf("request %d: allowed\n", i)
		} else {
			fmt.Printf("request %d: RATE LIMITED\n", i)
		}
	}
}
