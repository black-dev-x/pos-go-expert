package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name  string
	Hours int
	Price float64
}

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	t.Execute(os.Stdout, Courses{
		{"Go", 20, 99.99},
		{"Python", 30, 199.99},
		{"Java", 40, 299.99},
	})
}
