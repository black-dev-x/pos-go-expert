package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// server is currently 5 seconds, we could test a timeout from the server,
	// a timeout from the client or a success by changing the timer
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, error := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081", nil)
	if error != nil {
		panic(error)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

}
