package main

import (
	"fmt"
	"time"
)

// demoSelect waits on MULTIPLE channels at once, proceeding with whichever
// one is ready first — like a switch statement, but for channel operations.
func demoSelect() {
	fmt.Println("--- select ---")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		c1 <- "from c1"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		c2 <- "from c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// demoSelectDefault shows a NON-BLOCKING channel check: "default" runs
// immediately if no channel is ready, instead of waiting.
func demoSelectDefault() {
	fmt.Println("--- select with default (non-blocking) ---")
	c := make(chan int)

	select {
	case v := <-c:
		fmt.Println("received", v)
	default:
		fmt.Println("nothing was ready — moved on instead of blocking")
	}
}

// demoSelectTimeout shows the common "give up waiting after N" pattern
// using time.After, which itself returns a channel that fires once.
func demoSelectTimeout() {
	fmt.Println("--- select with timeout ---")
	c := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		c <- "finally done"
	}()

	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timed out waiting")
	}
}
