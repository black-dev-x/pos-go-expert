package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{"header.html", "content.html", "footer.html"}
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	t.Execute(os.Stdout, Courses{
		{"Go", 20},
		{"Python", 30},
		{"Java", 30},
	})

}
