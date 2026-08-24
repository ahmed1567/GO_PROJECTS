package main

import (
	"fmt"
	"net/http"
	"time"
)

// monitorLinksForever is the real-world example: it never returns. It
// combines everything above — goroutines, a channel, and a channel
// listener ("for range c") — into a continuous link monitor. Not called
// automatically from main(), since it would never let the program end;
// call it yourself (e.g. from main) when you want to run it standalone.
func monitorLinksForever() {
	links := []string{
		"https://www.google.com",
		"https://www.bing.com",
		"https://www.yahoo.com",
	}

	c := make(chan string) // a channel lets goroutines send results back to the listener

	for _, link := range links {
		go checkLink(link, c)
	}

	// the channel listener: blocks and waits for the NEXT value every
	// time, forever, until the channel is closed (which never happens here).
	for link := range c {
		go func(link string) {
			time.Sleep(5 * time.Second) // don't hammer the servers nonstop
			go checkLink(link, c)       // re-check the SAME link and feed the result back in
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down:", err)
		c <- link
		return
	}

	fmt.Println("Successfully checked", link)
	c <- link
}
