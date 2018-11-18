package main

import (
	"fmt"
)

func main() {

	x := 10
	fmt.Println("In main x = ", x, "/ x address = ", &x)
	changeByValue(x)
	fmt.Println("In main after changeByValue x = ", x)

	changeByReference(&x)
	fmt.Println("After changeByReference x = ", x)

	xPointer := new(int)
	*xPointer = 20
	fmt.Println("xPointer = ", *xPointer, "xPointer address = ", xPointer)

	k1 := 1
	k2 := 2
	fmt.Println("k1 = ", k1, "k2 = ", k2)
	swap(&k1, &k2)
	fmt.Println("k1 = ", k1, "k2 = ", k2)

}

func changeByValue(x int) {
	fmt.Println("In changeByValue x = ", x)
	x++
	fmt.Println("In changeByValue x now = ", x)
}

func changeByReference(x *int) {
	*x++
}

func swap(x *int, y *int) {
	t := *x
	*x = *y
	*y = t
}
