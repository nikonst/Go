package main

import (
	"fmt"
)

func main() {

	somePrint()
	defer someSecondPrint() // Executed in LIFO order
	defer printLoop()       // Executed in LIFO order

	str := "Hello Go!!"
	printStr(str)

	x, y := 1, 2
	println(average(x, y))

	var min, max int
	min, max = minmax(x, y)
	println("Min=", min, " Max=", max)

	k := func(x int) int { return x * x }
	println(k(5))

	println("Add1", add(1, 2))
	println("Add2", add(1, 2, 3, 4, 5, 6))

	c1 := aClosure()
	println("cl=", c1())
	println("cl=", c1())

	c2 := bClosure(10)
	println("c2=", c2(5))
	println("c2=", c2(25))

	println("Factorial ", nFactorial(6))

	containsPanic(5)
	containsPanic(-1)

	println("Program's last statement.")
}

func containsPanic(k int) {
	defer func() {
		if r := recover(); r != nil {
			println("****** It's OK, I recoverd", r)
		}
	}()
	p := 0
	if k <= p {
		panic(fmt.Sprint("PANIC k=", k))
	}
	for p = 0; p < k; p++ {
		println("p=", p)
	}
}

func somePrint() {
	println("I print some text.")
}

func someSecondPrint() {
	println("I print a SECOND text.")
}

func printLoop() {
	for i := 0; i < 4; i++ {
		defer println("Print Loop i=", i) // LIFO Order
	}
}

func printStr(s string) {
	println(s)
}

func average(x int, y int) float32 {
	return float32((x + y) / 2)
}

func minmax(x int, y int) (int, int) {
	a := 0
	b := 0
	if x <= y {
		a = x
		b = y
	} else {
		a = y
		b = x
	}
	return a, b
}

func add(args ...int) int { //Variadic function
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}
func aClosure() func() int { // Closure
	i := 0
	return func() int {
		i++
		return i
	}
}

func bClosure(x int) func(k int) int { // Closure
	a := x
	return func(k int) int {
		a += k
		return a
	}
}

func nFactorial(n int) int { // Recursion
	if n == 0 {
		return 1
	} else {
		return n * nFactorial(n-1)
	}
}
