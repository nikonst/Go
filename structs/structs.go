package main

import (
	"fmt"
)

type Rectangle struct {
	x float32
	y float32
}

func calculateArea(r Rectangle) float32 { // A function - Not a amethod
	return r.x * r.y
}

func (r *Rectangle) area() float32 { // A Method
	return r.x * r.y
}

type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Hi! I am ", p.name)
}

type Artist struct { // Embedded
	Person
	art string
}

func (a *Artist) sayYourArt() {
	fmt.Println("My art is ", a.art)
}

func main() {
	var r1 Rectangle
	r2 := new(Rectangle) // Pointer

	r1.x = 5
	r1.y = 10
	r2.x = 2
	r2.y = 4
	fmt.Println("r1:", r1.x, " ", r1.y)
	fmt.Println("r2:", r2.x, " ", r2.y)

	r3 := Rectangle{x: 1, y: 3}
	fmt.Println("r3:", r3.x, " ", r3.y)

	r4 := Rectangle{6, 7}
	fmt.Println("r4:", r4.x, " ", r4.y)

	changeRectangleProperties(r4)
	fmt.Println("Back in main -> r4:", r4.x, " ", r4.y)

	changeRectanglePropertiesPointers(&r4)
	fmt.Println("Back in main -> r4:", r4.x, " ", r4.y)

	fmt.Println("r1 area:", calculateArea(r1)) // Not a method
	fmt.Println("r2 area method:", r2.area())

	p1 := Person{"Anderson"}
	p1.talk()

	//a1 := new(Artist{name: "Mozart", art: "Music"})
	var a1 Artist
	a1.name = "Mozart"
	a1.art = "Music"
	a1.talk()
	a1.sayYourArt()

	checkInterfaces()
}

func changeRectangleProperties(r Rectangle) {
	r.x = 100
	r.y = 100
	fmt.Println("In changeRectangleProperties x and y", r.x, r.y)
}

func changeRectanglePropertiesPointers(r *Rectangle) {
	r.x = 100
	r.y = 100
	fmt.Println("In changeRectanglePropertiesPointer x and y", r.x, r.y)
}
