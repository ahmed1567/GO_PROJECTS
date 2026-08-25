package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

// realProductionExample is what you'd actually deploy: signal.NotifyContext
// wires OS signals (Ctrl+C = SIGINT, or SIGTERM from a container orchestrator
// like Docker/k8s telling your process to stop) directly into a context.
// It is NOT called by demoGracefulShutdown below — this is reference code
// only, since we can't wait on a real Ctrl+C during an automated `go run .`.
func realProductionExample() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	runServer(ctx) // <-- your actual server/worker loop goes here

	fmt.Println("received shutdown signal, draining in-flight work...")
	// close listeners, wait for in-flight requests/jobs to finish, then exit
}

// runServer stands in for "the real long-running loop" — an HTTP server,
// a message consumer, a worker pool, etc. It keeps doing work until ctx
// says to stop.
func runServer(ctx context.Context) {
	tick := time.NewTicker(20 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return // graceful: stop picking up NEW work
		case <-tick.C:
			fmt.Println("handling one unit of work...")
		}
	}
}

// demoGracefulShutdown is the same runServer loop, but driven by a
// timeout-based context instead of a real OS signal, so it terminates on
// its own for verification. In production you'd swap this ctx for the
// one from signal.NotifyContext above — runServer doesn't know or care
// which kind of context it was given.
func demoGracefulShutdown() {
	fmt.Println("--- graceful shutdown (timeout stands in for Ctrl+C here) ---")

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Millisecond)
	defer cancel()

	runServer(ctx)

	fmt.Println("shutdown signal received, draining finished — exiting cleanly")
}
