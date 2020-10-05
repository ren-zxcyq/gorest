// Package main of module uuidgenerator basically defines service that listens on
// localhost:8000 and returns a 10 digit random alphanumeric pattern over HTTP
// by reading from /dev/rand via "crypto/rand.Read()"
//
// To test
// curl -X GET localhost:8000
// 1e5aa61e5c9d356c3f0a
package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
)


// UUID is a custom multiplexre
type UUID struct {
}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}

func main() {
	mux := &UUID{}
	http.ListenAndServe(":8000", mux)
}
