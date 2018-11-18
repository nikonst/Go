package main

import "fmt"

/* Multiples of 3 and 5
Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

Ths Sum is: 234168
*/

func multiples3and5() {
	var s []int
	sum := 0
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			s = append(s, i)
			sum += i
		}
	}
	fmt.Println(s)
	fmt.Println("Sum=", sum)
}
