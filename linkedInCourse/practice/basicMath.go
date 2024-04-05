package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 3, 5, 7
	intSum := i1 + i2 + i3
	fmt.Println("Int Sum = ", intSum)

	f1, f2, f3 := 23.5, 34.94, 23.1
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum = ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Rounded Float Sum = ", floatSum)

	fmt.Println(math.Pi)
	fmt.Printf("PI = %.2f", math.Pi)
}
