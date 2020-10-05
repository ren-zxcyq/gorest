// Package main of module muxRouter contains a gorilla/mux example.
//
// It exposes a service at "localhost:8000/articles/{category}/{id:[0-9]+}".
// The multiplexer defined in Gorilla/mux maps incoming requests of the above
// format to a handler.
//
// Endpoints mapped can also be REGULAR EXPRESSIONS.
//
// Procedure:
// 1) Declare multiplexer:
// 		r := mux.NewRouter()
// 2) Declare endpoint Handler function:
// 		func ArticleHandler(w http.ResponseWriter, r *http.Request)
// 3) using the in-built http.Server declare our mux instance as the Handler
// 		srv := &http.Server{
// 				Handler: r, 		
//				Addr: "127.0.0.1:8000",
// 				WriteTimeout: 15 * time.Second,
//				ReadTimeout: 15 * time.Second,
//			}
//
// 4) Listen & Serve:
//		log.Fatal(srv.ListenAndServe())
//
// Test:
// 	curl -X GET localhost:8000/articles/testing-frameworks/3
// Results in
// 	Category is: testing-frameworks
// 	ID is: 3
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
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)	// mux.Vars(request) returns a map which contains our URL vars
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}
