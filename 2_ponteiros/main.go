package main

func main() {
	a := 10
	var ponteiro *int = &a
	ponteiro2 := &a

	*ponteiro = 20

	b := &a

	println(ponteiro)
	println(ponteiro2)
	println(&a)
	println(a)
	println(*b)
}
