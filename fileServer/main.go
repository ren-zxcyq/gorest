// Package main of module fileServer exposes a fileserver at "localhost:8000/static/*".
// In order to do so, it uses httprouter
// 1) create httprouter	-	router := httprouter.New()
// 2) router.ServeFiles("relative_path_to_the_one_in_http_dot_Dir/static/*filepath", http.Dir("absolute_path_to_static_folder"
// 3) http.ListenAndServe(":8000", router)
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/home/zxcyq/go/src/github.com/ren-zxcyq/gorest/fileServer/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
