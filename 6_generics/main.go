package main

func Add[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}

	println(Add(m))
	println(Add(m2))
}
