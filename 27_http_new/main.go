package main

import (
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Set("Accept", "application/json")
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	println(string(body))
}
