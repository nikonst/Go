package main

import (
	"fmt"
)

var ( // I don't include 'unsigned' integers and 'rune' (int32)
	x int
	y int16
	z int32
	k int64
	l byte

	f1 float32
	f2 float64

	str string

	b bool
)

const (
	// PI constant
	PI = 3.14159265359
)

func main() {
	x = 9223372036854775807
	y = 32767
	z = 2147483647
	k = 9223372036854775807
	l = 255
	fmt.Println("Integers:", x, y, z, k, l)

	f1 = +3.4E+38
	f2 = +1.7E+308
	fmt.Println("Floats:", f1, f2)

	str = "Go Programming"
	fmt.Println("A String:", str)

	b = true
	fmt.Println("Boolean:", b)

	fmt.Println("PI:", PI)

	f3 := 12.5678
	fmt.Printf("The value of x is %d and f1 is %8.2f\n", l, f3) // C like
	booleanOps()
	arithmeticOps()
	printMathFunctions()
	printStringFunctions()
	varDeclarations()
}
