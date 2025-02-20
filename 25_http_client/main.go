package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second}
	resp, error := client.Get("https://google.com")
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	println(string(body))
}
