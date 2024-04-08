package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	anInt := 23
	var p *int
	p = &anInt
	fmt.Println("Value of p:", *p)

	aFloat := 12.3
	pf := &aFloat

	fmt.Println("Value of pf:", *pf)

	*pf += 1
	fmt.Println("Value of aFloat:", aFloat)
}
