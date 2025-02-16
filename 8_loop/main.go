package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	numbers := []string{"one", "two", "three"}
	for k, v := range numbers {
		println(k+1, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Infinite loop")
	}
}
