// select.go
package main

import (
	"fmt"
	"time"
)

func doSomethingHighPriority() {
	fmt.Println(time.Now())
}

func doSomethingLowPriority() {
	fmt.Println("default action")
}

func sendTick(rateLimiter chan<- bool) {
	rate := time.Tick(time.Second / 2)
	for range rate {
		rateLimiter <- true
	}
}

// Notice the arrows on the channel types!
func receive(rateLimiter <-chan bool) {
	select {
	case <-rateLimiter:
		doSomethingHighPriority()
	default:
		doSomethingLowPriority()
	}
}

func main() {
	rateLimiter := make(chan bool, 2)
	go sendTick(rateLimiter)
	for {
		receive(rateLimiter)
		time.Sleep(time.Second / 5) // Sleep to not perform the default action too much
	}
}
