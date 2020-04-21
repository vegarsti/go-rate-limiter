// channels.go
package main

import "fmt"

func sendMessage(messages chan<- string) {
	messages <- "ping"
}

func main() {
	messages := make(chan string)
	go sendMessage(messages)
	msg := <-messages
	fmt.Println(msg)
}
