// Package main of healthCheck exposes service "localhost:8000/healthCheck"
// which just returns the servers local time as text.
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// HealthCheck API returns date time to client
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/healthcheck", HealthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
