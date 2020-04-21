// ticker.go
package main

import (
	"fmt"
	"time"
)

func sendTick(rateLimiter chan<- bool) {
	rate := time.Tick(time.Second / 2)
	for range rate {
		rateLimiter <- true
	}
}

func doSomething() {
	fmt.Println(time.Now())
}

func receive(rateLimiter <-chan bool) {
	for {
		<-rateLimiter
		doSomething()
	}
}

func main() {
	rateLimiter := make(chan bool, 2)
	go sendTick(rateLimiter)
	receive(rateLimiter)
}
