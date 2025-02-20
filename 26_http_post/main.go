package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second}
	jsonVar := bytes.NewBuffer([]byte(`{"key1":"value1","key2":"value2"}`))
	resp, error := client.Post("https://google.com", "application/json", jsonVar)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
