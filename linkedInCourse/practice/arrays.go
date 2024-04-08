package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Arrays")

	var colors = [3]string{"Red", "Green", "Blue"} // Array, size defiend

	fmt.Println(colors)
	fmt.Println("Length of colors:", len(colors))

	var cities = []string{"London", "Paris", "Rome"} //Slice, no size
	cities = append(cities, "Oslo")
	fmt.Println(cities)

	cities = append(cities[1:len(colors)])
	fmt.Println(cities)

	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)
	fmt.Println(states)

	delete(states, "OR")
	states["NY"] = "New York"
	fmt.Println(states)

	states["AZ"] = "Arizona"
	fmt.Println(states) // sorted

	for k, v := range states {
		// not sorted
		fmt.Printf("%v: %v\n", k, v) // key - value
	}

}
