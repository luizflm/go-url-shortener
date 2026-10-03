package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./assets/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", home)

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
