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

	mary := Person{65, "I am Mary", "Hello everyone"}
	fmt.Println(mary)
	mary.introduceYourself()
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

type Person struct {
	Weight int
	Talk   string
	Speak  string
}

// introduceYourself is ...
func (p Person) introduceYourself() {
	fmt.Println(p.Talk, p.Speak)
}
