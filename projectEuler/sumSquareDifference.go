/*
Project Euler

Sum square difference
Problem 6

The sum of the squares of the first ten natural numbers is,
12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

*/
package main

func sumSquareDifference() {
	diff := squareOfTheSum(100) - someOfTheSquares(100)
	println("Sum Square Difference : ", diff)

}

func someOfTheSquares(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s = s + i*i
	}
	return s
}

func squareOfTheSum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s = s + i
	}
	return s * s
}

// The Sum square difference is 25164150
