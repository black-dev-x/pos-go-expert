package main

import (
	"github.com/black-dev-x/pos-go-expert/35_go_mod/math"
	"github.com/google/uuid"
)

func main() {
	println(uuid.New().String())
	m := math.Math{A: 1, B: 2}
	result := m.Add()
	println(result)
	println("Hello, World!")
}
