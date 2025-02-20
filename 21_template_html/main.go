package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name  string
	Hours int
	Price float64
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		t.Execute(w, Courses{
			{"Go", 20, 99.99},
			{"Python", 30, 199.99},
			{"Java", 40, 299.99},
		})
	})
	http.ListenAndServe(":8080", nil)
}
