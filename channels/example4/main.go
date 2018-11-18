package main

import (
	"fmt"
)

/* When using channels as function parameters, I can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program.
*/

func ping(cin chan<- string, message string) {
	cin <- message
}

func pong(cout <-chan string, cin chan<- string) {
	msg := <-cout
	msg += " Pong"
	cin <- msg
}
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go ping(c1, "Ping")
	go pong(c1, c2)

	fmt.Println(<-c2)

	go pong(c1, c2) // reverse order of the teo statements - same response
	go ping(c1, "Pang")

	fmt.Println(<-c2)

}
