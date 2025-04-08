package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/broker/", handleBrokerRequests)

	port := ":8090"
	log.Printf("Starting broker service on port %s", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
			log.Fatalf("Broker server failed to start: %v", err)
	}
}