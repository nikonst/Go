package main

import (
	"fmt"
	"time"
)

func func1(c chan<- string) {
	time.Sleep(1 * time.Second)
	c <- "one"
}

func func2(c chan<- string) {
	time.Sleep(3 * time.Second)
	c <- "two"
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func2(c2) // func2 before func1
	go func1(c1)

	for i := 0; i < 2; i++ { // Wait on multiple channel operations
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}
	}
}

/*
Output:

Received: one
Received: two

*/
