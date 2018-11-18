package main

import (
	"fmt"
	"math"
)

func printMathFunctions() {
	fmt.Printf("Cos %.2f\n", math.Cos(math.Pi/2))
	f := 12.34567
	fmt.Printf("Sin %.2f\n", math.Sin(f/2))
	fmt.Printf("Tan %.2f\n", math.Tan(f/3))
	fmt.Printf("Sqrt %.2f\n", math.Sqrt(f*2))
	fmt.Printf("Exp %.2f\n", math.Exp(f))
	fmt.Printf("Power %f\n", math.Pow(2, 10))
}
