// buffering.go
package main

import "fmt"

func main() {
	messages := make(chan int, 2)
	messages <- 1
	messages <- 2
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
