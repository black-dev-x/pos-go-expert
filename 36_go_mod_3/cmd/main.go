package main

import (
	"github.com/black-dev-x/pos-go-expert/36_go_mod_3/math"
	"github.com/google/uuid"
)

func main() {
	println(uuid.New().String())
	m := math.Math{A: 1, B: 2}
	result := m.Add()
	println(result)5
	println("Hello, World!")
}
