package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Started")
	defer log.Println("Request Ended")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Processed successfully")
		w.Write([]byte("Processed successfully"))
	case <-ctx.Done():
		log.Println("Request Canceled")
	}
}
