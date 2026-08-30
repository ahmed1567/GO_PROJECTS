package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// saveProgress persists how far a long-running job has gotten, so if the
// process crashes and restarts, it can resume instead of starting over —
// the exact "Redis-based crash recovery state" pattern from the roadmap.
func saveProgress(ctx context.Context, rdb *redis.Client, jobID string, itemsDone int) {
	key := "job:" + jobID
	rdb.HSet(ctx, key, "items_done", itemsDone)
}

func loadProgress(ctx context.Context, rdb *redis.Client, jobID string) int {
	key := "job:" + jobID
	val, err := rdb.HGet(ctx, key, "items_done").Int()
	if err != nil {
		return 0 // no saved progress yet — start from scratch
	}
	return val
}

// runJob simulates processing a batch of items, saving progress after
// each one. If this function stops halfway (crash, deploy, OOM-kill), the
// NEXT run calls loadProgress and resumes instead of redoing everything.
func runJob(ctx context.Context, rdb *redis.Client, jobID string, totalItems int) {
	start := loadProgress(ctx, rdb, jobID)
	if start > 0 {
		fmt.Printf("resuming job %s from item %d/%d\n", jobID, start, totalItems)
	}

	for i := start; i < totalItems; i++ {
		saveProgress(ctx, rdb, jobID, i+1)
	}
	fmt.Printf("job %s finished: %d/%d items done\n", jobID, totalItems, totalItems)
}

func demoCrashRecovery(ctx context.Context, rdb *redis.Client) {
	fmt.Println("--- crash-recovery state via Redis hash ---")

	jobID := "batch_import_1"
	rdb.Del(ctx, "job:"+jobID) // clean slate

	// Simulate the job crashing partway through by saving progress
	// directly, without actually running items 1-6.
	saveProgress(ctx, rdb, jobID, 6)
	fmt.Println("(simulating a crash after 6/10 items)")

	// A fresh process/restart calls runJob next — it resumes from item 6
	// instead of redoing items 1-6.
	runJob(ctx, rdb, jobID, 10)
}
