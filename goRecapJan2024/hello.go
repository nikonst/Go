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
	fmt.Printf("Type: %T, value: %v", y, y)
}
