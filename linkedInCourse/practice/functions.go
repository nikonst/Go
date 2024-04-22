package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")

	x := 2
	voidFunction()
	fmt.Println(aFunction(x))
	s, c := addAllValues(1, 3, 4)
	fmt.Println(s, " ", c)
	s, c = addAllValues(1, 3, 4, 5, 6, 7)
	fmt.Println(s, " ", c)
}

func voidFunction() {
	fmt.Println("In Void Function")
}

func aFunction(a int) int {
	return a * a
}

func addAllValues(values ...int) (int, int) {
	t := 0
	for _, v := range values {
		t += v
	}
	return t, len(values)
}
