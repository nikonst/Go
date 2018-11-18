package main

import "fmt"

func varDeclarations() {
	var x int
	x = 10

	s := "Hello there"

	var f float32 = 2.3

	fmt.Printf("%d - %s - %f", x, s, f)
}
