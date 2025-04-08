package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Quote struct {
	Quote string `json:"quote"`
}

func getQuoteFromBroker(endpoint string) (Quote, error) {
	resp, err := http.Get(fmt.Sprintf("http://broker-service:8090/broker/quotes%s", endpoint))
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()

	var quote Quote
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		return Quote{}, err
	}
	return quote, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var quote Quote
	var err error

	if r.URL.Path == "/random" {
		quote, err = getQuoteFromBroker("/random")
	} else {
		quote, err = getQuoteFromBroker("")
	}

	if err != nil {
		http.Error(w, "Failed to get quote", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/templates/index.html"))
	tmpl.Execute(w, quote)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.HandleFunc("/random", handler)
	log.Println("Starting front-end on :8080")
	http.ListenAndServe(":8080", nil)
}
