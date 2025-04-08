package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/quotes/random", GetRandomQuoteHandler)
	http.HandleFunc("/quotes", GetAllQuotesHandler)

	port := ":8091"

	// Start the HTTP server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
