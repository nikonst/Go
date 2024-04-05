package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	var aString string = "Some string"
	fmt.Println("Hello ", aString)
	fmt.Printf("the var's type is %T \n", aString)

	var anInt int = 11
	fmt.Println("Hello ", anInt)
	fmt.Printf("the var's type is %T ", anInt)

	var deafultInt int
	fmt.Println("Hello ", deafultInt)
	fmt.Printf("the var's type is %T ", deafultInt)

	var anotherString = "Another string" // no type
	fmt.Println("Hello ", anotherString)
	fmt.Printf("the var's type is %T \n", anotherString)

	var anotherInt = 1 // no type
	fmt.Println("Hello ", anotherInt)
	fmt.Printf("the var's type is %T ", anotherInt)

	myString := "Oslo"
	// colon equals, no var, only inside functions
	// outside a function 'var' is needed
	fmt.Println("Hello ", myString)
	fmt.Printf("the var's type is %T ", myString)

	fmt.Println("Constant ", aConst)
	fmt.Printf("the const's type is %T ", aConst)

}
