package main

func main() {
	x := 1
	y := 2

	if x > y {
		println("x>y")
	} else if x < y {
		println("x>y")
	} else {
		println("x=y")
	}

	k := 5
	switch k {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	default:
		println("Some other day")
	}

	k = 15
	if k%2 == 0 {
		println("k even")
	} else {
		println("k odd")
	}

	if true && false == true {
		println("Alpha")
	} else {
		println("Beta")
	}

	z := 1
	if z == 10 { // NOT if z {}, NOT if z=10 {}
		println("True")
	} else {
		println("False")
	}

	s1 := "hello"
	s2 := "hello"
	if s1 == s2 {
		println("Same")
	} else {
		println("Different")
	}

	if v := 10; v < 0 {
		println("v negative")
	} else {
		println("v postive or zero")
	}
}
