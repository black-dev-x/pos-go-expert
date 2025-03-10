package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("public"))
	mux.Handle("/", fileServer)
	http.ListenAndServe(":8080", mux)
}
