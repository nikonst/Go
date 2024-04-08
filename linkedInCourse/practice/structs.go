package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// Upper case is exported, otherwise is private
type Dog struct {
	Breed  string // Upper case is exported, otherwise is private
	Weight int
}
