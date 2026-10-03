package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Print("starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
