// Package main of module queryParameters contains an example of how to extract
// GET params from a request using "gorilla/mux".
// In order to extract them we are basically using the same procedure as in
// module muxRouter
// 1) declare a new multiplexer:
//	r := mux.NewRouter()
// 2) attach endpoint handler:
//	r.HandleFunc("/articles", QueryHandler)
// 3) Use the built-in http.Server module; declare our multiplexer as the handler
// 		srv := http.Server{
// 			Handler: r,
// 			Addr: "127.0.0.1:8000",
// 			WriteTimeout: 15 * time.Second,
// 			ReadTimeout: 15 * time.Second,
// 		}
// 4) log.Fatal(srv.ListenAndServe())
//
// Test:
// 	curl -X GET "http://localhost:8000/articles?id=4&category=sci-fi
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)

	srv := http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()	// r.URL.Query() returns a map which contains our URL params
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id: %s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category: %s\n", queryParams["category"][0])
}
