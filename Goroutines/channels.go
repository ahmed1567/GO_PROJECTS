package main

import "fmt"

// demoUnbufferedChannel: send blocks until someone receives, and receive
// blocks until someone sends — they must "meet" at the same moment. This
// is why sending and receiving here happen in different goroutines.
func demoUnbufferedChannel() {
	fmt.Println("--- unbuffered channel ---")
	c := make(chan string) // capacity 0

	go func() {
		c <- "hello from goroutine" // blocks until main receives
	}()

	msg := <-c // blocks until the goroutine sends
	fmt.Println(msg)
}

// demoBufferedChannel: a buffered channel can hold N values with nobody
// receiving yet — send only blocks once the buffer is completely full.
func demoBufferedChannel() {
	fmt.Println("--- buffered channel ---")
	c := make(chan int, 3) // capacity 3

	c <- 1 // none of these three block — buffer has room
	c <- 2
	c <- 3
	// c <- 4 // would BLOCK here — buffer is full and nobody's receiving yet

	fmt.Println(len(c), "items buffered")
	fmt.Println(<-c, <-c, <-c)
}

// sendOnly demonstrates a channel DIRECTION restriction in a function
// signature: chan<- int means this function may only SEND into c —
// the compiler enforces it, receiving would be a compile error.
func sendOnly(c chan<- int, value int) {
	c <- value
}

// receiveOnly: <-chan int means this function may only RECEIVE from c.
func receiveOnly(c <-chan int) int {
	return <-c
}

func demoChannelDirections() {
	fmt.Println("--- channel directions ---")
	c := make(chan int, 1)
	sendOnly(c, 42)
	fmt.Println(receiveOnly(c))
}

// demoClosingChannel shows closing a channel and the two ways to detect
// it: "range" stops automatically once drained, and comma-ok tells you
// whether a channel is closed vs still open.
func demoClosingChannel() {
	fmt.Println("--- closing a channel ---")
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c) // no more sends allowed after this — sending to a closed channel PANICS

	for v := range c { // range stops automatically once the channel is closed AND drained
		fmt.Println("received:", v)
	}

	v, ok := <-c // reading a closed, empty channel never blocks — zero value + ok=false
	fmt.Println("after close, read again:", v, ok)
}
