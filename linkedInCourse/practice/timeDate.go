package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("The time now:", n)

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}
