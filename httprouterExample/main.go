// Package main of httprouterExample exposes two services.
// - "localhost:8000/api/v1/go-version" runs "go version" and returns that over HTTP.
// - "localhost:8000/api/v1/show-file/:name" runs "cat file" and returns the provided files contents.
// Test via:
//	curl -X GET http://localhost:8000/api/v1/servicename
// Note:
// The show-file service serves files found under the github projects root directory.
// If the user requests files in a different folder as in "/show-file/foldername/file"
// an Error 404 is returned
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("/usr/bin/go", "version")
	io.WriteString(w, response)
	//return
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//fmt.Fprintf(w, getCommandOutput("pwd", ""))
	fmt.Fprintf(w, getCommandOutput("/usr/bin/cat", params.ByName("name")))
}
