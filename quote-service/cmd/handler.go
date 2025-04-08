package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetAllQuotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Access the Quotes slice from data.go
	err := json.NewEncoder(w).Encode(Quotes)
	if err != nil {
		log.Printf("Error encoding all quotes: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetRandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(Quotes) == 0 {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}

	// Select a random quote
	randomQuote := Quotes[Rng.Intn(len(Quotes))]

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(randomQuote)
	if err != nil {
		log.Printf("Error encoding random quote: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}