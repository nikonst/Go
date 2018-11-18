package main

import "fmt"

func booleanOps() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func arithmeticOps() {
	fmt.Println(1 + 1)
	fmt.Println(2 * 1.5)
	fmt.Println(2 / 1.5)
	fmt.Println(7 % 3)
	x := 1
	y := 2
	x-- // --x won't work
	y++
	fmt.Println(x, y)
}
