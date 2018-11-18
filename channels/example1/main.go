package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() { messages <- "hello on channel" }()
	msg := <-messages
	fmt.Println("msg:", msg)
}
