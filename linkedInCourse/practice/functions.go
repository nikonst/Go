package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")

	x := 2
	voidFunction()
	fmt.Println(aFunction(x))
	fmt.Println(addAllValues(1, 3, 4))
	fmt.Println(addAllValues(1, 3, 4, 10, 11))
}

func voidFunction() {
	fmt.Println("In Void Function")
}

func aFunction(a int) int {
	return a * a
}

func addAllValues(values ...int) int {
	t := 0
	for _, v := range values {
		t += v
	}
	return t
}
