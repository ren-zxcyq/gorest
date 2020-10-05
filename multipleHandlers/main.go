// Package main contained in multipleHandlers uses http.NewServeMux() to create a
// multiplexer, which in turn assigns handlers to URLs.
// Assigning handlers happens via "http.ServeMux.HandleFunc()"
// Services:
// - /randomFloat
// - /randomInt
// Note:
// 	the urls are CASE SENSITIVE
package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})

	http.ListenAndServe(":8000", newMux)
}
