package main

import (
	"fmt"
	"strings"
)

func printStringFunctions() {
	var str string
	str = "Go Programming!!!"

	fmt.Printf("The string is %s\n", str)
	fmt.Printf("The string's length is %d\n", len(str))
	fmt.Println(strings.Contains("sea", str)) // Second to be contained in first
	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Index(str, "gram"))
	fmt.Println(strings.LastIndex(str, "m"))
	fmt.Println(strings.Replace(str, "gr", "rg", 1))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
}
