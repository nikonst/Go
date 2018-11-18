package main

import (
	"fmt"
)

//ARRAYSIZE const
const ARRAYSIZE int = 5

func main() {
	var a [ARRAYSIZE]int // Array
	var s []int          //Slice

	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		a[i] = i * i
	}
	fmt.Println(a)
	fmt.Println("Printing with Range")
	for _, v := range a {
		println(v)
	}
	fmt.Println("Array sum= ", sum(a))
	fmt.Println("Array avg= ", avg(a))

	fmt.Println(a[1:3]) // Print a slice
	fmt.Println(a[:3])  // Print a slice
	fmt.Println(a[2:])  // Print a slice
	fmt.Println(a[:])   // Print a slice

	changeValues(a[:])
	fmt.Println(a)

	//a = append(a, 133) Wrong -> a is an Array
	s = append(s, 10) // Right -> s is a Slice
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
	s = append(s, 11, 22, 33)
	fmt.Println(s)
	fmt.Println("S len = ", len(s))
	min, max := minmax(s)
	fmt.Println("S min max :", min, max)

	var strSlice = []string{"One", "Two", "Three"}
	strSlice2 := make([]string, 2)

	fmt.Println(strSlice)
	copy(strSlice2, strSlice)
	fmt.Println(strSlice2)

	aMap := make(map[string]string)
	aMap["H"] = "Hydrogen"
	aMap["He"] = "Helium"
	aMap["Li"] = "Lithium"
	aMap["Be"] = "Beryllium"
	aMap["B"] = "Boron"
	fmt.Println(aMap)
	for key, value := range aMap {
		println("key=", key, " value=", value)
	}
	delete(aMap, "Li")
	fmt.Println(aMap)
	_, found := aMap["Li"]
	fmt.Println(found)
	aMap["O"] = "Oxygen"
	fmt.Println(aMap)
}

func sum(a [ARRAYSIZE]int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func avg(a [ARRAYSIZE]int) float32 {
	return float32(sum(a) / len(a))
}

func changeValues(k []int) { // It is a call by reference
	for i := 0; i < len(k); i++ {
		k[i]++
	}
	fmt.Println("In change values function:", k)
}

func minmax(a []int) (int, int) {
	min := a[0]
	max := a[0]
	for i := 0; i < len(a); i++ {
		if min >= a[i] {
			min = a[i]
		}
		if max <= a[i] {
			max = a[i]
		}
	}
	return min, max
}
