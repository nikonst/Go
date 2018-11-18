package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2)

	messages <- "hello on channel"
	messages <- "hello again"

	fmt.Println("msg:", <-messages)
	fmt.Println("msg:", <-messages)
}
