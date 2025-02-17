package main

import (
	"io"
	"net/http"
)

func main() {
	req, error := http.Get("https://www.google.com")
	treatError(error)
	defer req.Body.Close()
	res, _ := io.ReadAll(req.Body)
	println(string(res))
	anotherExample()
}

func anotherExample() {
	defer println("First line")
	println("Second line")
	println("Third line")
}
func treatError(err error) {
	if err != nil {
		panic(err)
	}
}
