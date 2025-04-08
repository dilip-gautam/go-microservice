// handler.go

package main

import (
        "fmt"
        "io"
        "log"
        "net/http"
)

const quoteServiceBaseURL = "http://quote-service:8091"

func proxyRequestHandler(w http.ResponseWriter, targetURL string) {
        log.Printf("Broker forwarding request to: %s", targetURL)

        resp, err := http.Get(targetURL)
        if err != nil {
                log.Printf("Broker: Error calling quote service at %s: %v", targetURL, err)
                http.Error(w, "The quote service is currently unavailable.", http.StatusServiceUnavailable)
                return
        }
        defer resp.Body.Close()

        for key, values := range resp.Header {
                for _, value := range values {
                        w.Header().Add(key, value)
                }
        }

        w.WriteHeader(resp.StatusCode)

        _, err = io.Copy(w, resp.Body)
        if err != nil {
                log.Printf("Broker: Error copying response body from %s: %v", targetURL, err)
        }
}

func handleBrokerRequests(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
                http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
                return
        }

        path := r.URL.Path
        var targetURL string

        switch path {
        case "/broker/quotes":
                targetURL = fmt.Sprintf("%s/quotes", quoteServiceBaseURL)
        case "/broker/quotes/random":
                targetURL = fmt.Sprintf("%s/quotes/random", quoteServiceBaseURL)
        default:
                http.NotFound(w, r)
                return
        }

        proxyRequestHandler(w, targetURL)
}