package main

import (
	"fmt"
	"sync"
)

// demoRaceCondition shows what goes wrong when many goroutines mutate the
// SAME variable with no protection: the final count is unreliable, because
// "count++" is really read-then-write, and two goroutines can read the
// same value before either writes it back, silently dropping an increment.
func demoRaceCondition() {
	fmt.Println("--- race condition (no protection) ---")
	count := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // NOT safe: concurrent read-modify-write
		}()
	}

	wg.Wait()
	fmt.Println("final count (unreliable, often < 1000):", count)
	fmt.Println("run 'go run -race .' to have Go's race detector flag this line")
}

// demoMutexFix protects the same operation with sync.Mutex so only one
// goroutine can touch "count" at a time — always correct, every run.
func demoMutexFix() {
	fmt.Println("--- fixed with sync.Mutex ---")
	count := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("final count (always correct):", count)
}
