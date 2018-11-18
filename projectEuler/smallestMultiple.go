/*
Project Euler

Smallest multiple
Problem 5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1? to 20
*/

package main

func smallestMultiple() {
	var i int
	var d int
	for i = 20; ; /* i < 25210; */ i = i + 10 {
		for d = 20; d >= 1; d-- {
			if i%d == 0 {
				println(i, " ", d, " div OK")
			} else {
				println(d, " No")
				break
				//continue
			}

		}
		if d == 0 {
			println("Break")
			break
		}
	}
}

// The number is 232.792.560
