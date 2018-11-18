package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printLoop(n int, s string) {
	if n > 0 {
		for i := 1; i < n; i++ {
			fmt.Println(s, " i=", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
		}
	} else {
		fmt.Println("Sorry.. no loop!")
	}
}
func main() {
	fmt.Println("Program Starts")
	go printLoop(10, "one")
	go printLoop(20, "two")
	go printLoop(-1, "three")
	go printLoop(15, "four")
	var input string
	fmt.Println("Insert a string:") // First that prints to screen
	fmt.Scanln(&input)              // Wait for the go functions to run
	fmt.Println("Program Ends")
}
