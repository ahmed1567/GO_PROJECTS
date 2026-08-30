package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// processEvent processes a webhook-style event exactly once, using
// Redis's SET NX ("set if not exists") as an atomic "claim this event ID"
// operation — the same idempotency trick real webhook/queue systems use
// to survive duplicate deliveries (retries, at-least-once delivery, etc.).
func processEvent(ctx context.Context, rdb *redis.Client, eventID string) {
	key := "idempotency:" + eventID

	// SetNX returns true only if the key didn't already exist. This
	// single call is ATOMIC, so two concurrent requests for the same
	// eventID can never both "win" the claim.
	claimed, err := rdb.SetNX(ctx, key, "processed", 24*time.Hour).Result()
	if err != nil {
		fmt.Println("redis error:", err)
		return
	}

	if !claimed {
		fmt.Println(eventID, "-> DUPLICATE, skipping (already processed)")
		return
	}

	fmt.Println(eventID, "-> processing for the first time")
}

func demoIdempotency(ctx context.Context, rdb *redis.Client) {
	fmt.Println("--- idempotency via SET NX ---")

	rdb.Del(ctx, "idempotency:evt_123", "idempotency:evt_456") // clean slate for the demo

	processEvent(ctx, rdb, "evt_123") // first delivery: processes
	processEvent(ctx, rdb, "evt_123") // duplicate delivery: skipped
	processEvent(ctx, rdb, "evt_456") // different event: processes
}
