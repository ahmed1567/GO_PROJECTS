package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	id    int
	value int
}

type result struct {
	jobID    int
	workerID int
	output   int
}

// worker pulls jobs off "jobs" until it's closed, processes each one, and
// sends a result — tagging it with its OWN workerID so we can see which
// of the pool's goroutines actually handled each job. Many workers safely
// share one jobs channel: each value sent is delivered to exactly ONE
// receiving goroutine, never duplicated, so no two workers ever grab the
// same job.
func worker(workerID int, jobs <-chan job, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(50 * time.Millisecond) // simulate real work — without this, one fast worker races ahead and drains everything before the others get scheduled
		results <- result{jobID: j.id, workerID: workerID, output: j.value * j.value}
	}
}

// demoWorkerPool ties goroutines + channels + WaitGroup into a realistic
// pattern: a fixed pool of workers processing a queue of jobs concurrently,
// with results collected back on a second channel.
func demoWorkerPool() {
	fmt.Println("--- worker pool ---")
	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan job, numJobs)
	results := make(chan result, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- job{id: i, value: i}
	}
	close(jobs) // tells every worker's "for range jobs" loop to stop once drained

	wg.Wait()      // wait for all workers to finish before touching results
	close(results) // safe now — nothing will send to it anymore

	for r := range results { // order is NOT guaranteed — workers finish in any order
		fmt.Printf("worker %d handled job %d -> %d\n", r.workerID, r.jobID, r.output)
	}
}
