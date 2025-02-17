package main

import (
	"io"
	"net/http"
)

func main() {
	req, _ := http.Get("https://www.google.com")
	res, _ := io.ReadAll(req.Body)
	println(string(res))
	req.Body.Close()
}
