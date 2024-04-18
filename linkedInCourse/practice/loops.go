package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	colors := []string{"Red", "Green", "Blue"}

	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println("--", color)
	}

	v := 1
	for v <= 10 {
		fmt.Println("v=", v)
		v++
	}

	s := 1
	for s <= 100 {
		fmt.Println("s=", s)
		s += s
		if s > 30 {
			break
		}
	}

}
