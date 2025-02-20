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
	templates := []string{"header.html", "content.html", "footer.html"}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	t.Execute(os.Stdout, Courses{
		{"Go", 20, 99.99},
		{"Python", 30, 199.99},
		{"Java", 40, 299.99},
	})
}
