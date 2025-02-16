package main

type MyNumber int
type Number interface {
	~int | float64
} // by using ~, it can consider types, like MyNumber, that are not directly int or float64

// func Add[T int | float64](m map[string]T) T {
func Add[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func Compare[T comparable](a T, b T) bool {
	return a == b
	// var x interface{} = 10
	// var y interface{} = "Hello, World!"
	// showType(x)
	// showType(y)
}
func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	m3 := map[string]MyNumber{"a": 1, "b": 2, "c": 3}

	println(Add(m))
	println(Add(m2))
	println(Add(m3))
	println(Compare(10, 10.0))
}
