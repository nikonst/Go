package main

import (
	"fmt"
)

type Shape interface { // Only for Methods
	perimeter() float32
	area() float32
}

type Square struct {
	a float32
}

type Circle struct {
	r float32
}

func (s Square) area() float32 {
	return s.a * s.a
}

func (s Square) perimeter() float32 {
	return 4 * s.a
}

func (c Circle) area() float32 {
	return 3.14 * c.r * c.r
}

func (c Circle) perimeter() float32 {
	return 2 * 3.14 * c.r
}

func printPerimeterAndArea(s Shape) {
	fmt.Println(s.area(), s.perimeter())
}

func checkInterfaces() {
	fmt.Println("********************")
	s1 := Square{5}
	c1 := Circle{4}
	fmt.Println("s1 area: ", s1.area())
	fmt.Println("c1 area: ", c1.area())
	printPerimeterAndArea(s1)
	printPerimeterAndArea(c1)
}
