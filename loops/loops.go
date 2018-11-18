package main

func main() {
	var i int
	for i = 0; i < 10; i++ {
		println(i)
	}

	println("***********************")
	i = 0
	for i < 10 {
		println("i=", i)
		i++
	}

	println("***********************")
	k1 := 1
	k2 := 25
	for k1 < k2 {
		println("k1=", k1, " k2=", k2)
		k1++
		k2--
	}

	println("--- CONTINUE ---")
	i = 100
	for i > 0 {
		if i%2 == 0 {
			println(i, "even")
			i--
		} else {
			i--
			continue
		}
	}

	println("--- BREAK ---")
	i = -10
	for i < 10 {
		println("i=", i)
		if i >= 0 {
			break
		}
		i++
	}

	s := 0
	i = 0
	for i <= 1000 {
		s += i
		i++
	}
	println("s=", s)

	println("***********************")
	for l, m := 0, 1; l < 10 || m < 5; l, m = l+1, m+1 {
		println("l=", l, "m=", m)
	}
}
