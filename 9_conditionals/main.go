package main

func main() {
	a := 1
	b := 2
	if a > b {
		println("a is greater than b")
	} else {
		println("b is greater than a")
	}

	switch a {
	case 1:
		println("a is 1")
	case 2:
		println("a is 2")
	default:
		println("a is neither 1 nor 2")
	}
}
