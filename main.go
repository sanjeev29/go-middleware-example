package main

import (
	"log"
	"net/http"
)

// DefaultHandler function
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {

	// Custom ServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("/", DefaultHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", mux))
}
