package main

import (
	"fmt"
	"time"
)

func worker(c chan bool) {
	fmt.Println("I start working")
	time.Sleep(time.Second * 3)
	fmt.Println("I finished")
	c <- true
}
func main() {
	c := make(chan bool)
	go worker(c)
	<-c // If I leave this out, the program will exit without waiting the go function

}
