package main

import (
	"context"
	"fmt"
	"time"
)

// demoWithCancel shows the base case: you control exactly when a context
// is cancelled, by calling its cancel function yourself.
func demoWithCancel() {
	fmt.Println("--- context.WithCancel ---")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel() // signals ctx.Done() to fire
	}()

	<-ctx.Done()
	fmt.Println("cancelled, reason:", ctx.Err())
}

// demoWithTimeout cancels itself automatically after a duration — the
// direct analogue to an AbortController + setTimeout in JS.
func demoWithTimeout() {
	fmt.Println("--- context.WithTimeout ---")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() // ALWAYS defer cancel, even on the success path — frees resources tied to the context

	<-ctx.Done()
	fmt.Println("timed out, reason:", ctx.Err())
}

// demoWithDeadline is WithTimeout's sibling: same idea, but you give an
// absolute point in time instead of a duration from now.
func demoWithDeadline() {
	fmt.Println("--- context.WithDeadline ---")
	deadline := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	<-ctx.Done()
	fmt.Println("deadline passed, reason:", ctx.Err())
}

// doWork simulates real work (like an HTTP call or a DB query) that
// respects cancellation instead of ignoring it — select on ctx.Done()
// alongside whatever channel the real work would use.
func doWork(ctx context.Context, workTime time.Duration) error {
	select {
	case <-time.After(workTime):
		return nil // finished before the context gave up on us
	case <-ctx.Done():
		return ctx.Err() // caller cancelled or timed out — stop early
	}
}

// demoPropagation shows a context passed down into a goroutine doing
// "real work" — the work stops as soon as the context is cancelled,
// instead of running to completion regardless.
func demoPropagation() {
	fmt.Println("--- propagating a context into work ---")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	err := doWork(ctx, 200*time.Millisecond) // work would take 200ms, but ctx only allows 30ms
	fmt.Println("doWork result:", err)        // context.DeadlineExceeded, not nil
}
