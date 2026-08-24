package main

import (
	"fmt"
	"sync"
	"time"
)

// demoFireAndForget shows the classic beginner mistake: main() doesn't wait
// for goroutines by default, so it can exit before they ever get to run.
func demoFireAndForget() {
	fmt.Println("--- fire and forget (the mistake) ---")
	go fmt.Println("this MIGHT never print — nothing is making main() wait for it")
	time.Sleep(10 * time.Millisecond) // only here so you can actually see it print; never use sleep for real synchronization
	fmt.Println("main continued without truly waiting for the goroutine")
}

// demoWaitGroup is the standard way to wait for a known number of
// goroutines to finish when you don't need any data back from them.
func demoWaitGroup() {
	fmt.Println("--- sync.WaitGroup ---")
	var wg sync.WaitGroup

	names := []string{"alice", "bob", "carol"}
	for _, name := range names {
		wg.Add(1) // "one more goroutine to wait for"
		go func(name string) {
			defer wg.Done() // MUST be called when the goroutine finishes, or Wait() blocks forever
			fmt.Println("hello from", name)
		}(name)
	}

	wg.Wait() // blocks here until every Add() has a matching Done()
	fmt.Println("all goroutines finished")
}
