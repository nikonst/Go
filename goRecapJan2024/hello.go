package main

import "fmt"

func main() {
	fmt.Println("Hello Go!!!")

	var a, b int

	a = 1
	b = 2
	fmt.Println("Values of a and b, ", a, " and", b)

	var myBool = true
	fmt.Println("Reversed value of myBook, ", !myBool)

	str := "This is a string value"
	fmt.Println("Value of someString, ", str)

	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", y, y)

	const s1 = "Hello"
	fmt.Printf("Type: %T, value: %v\n", s1, s1)

	for n := 0; n <= 50; n++ {
		if n%2 == 0 {
			fmt.Println("Even ", n)
		}
	}
}
